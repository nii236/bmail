package main

import (
	"bmail/internal/pkg/config"
	"bmail/internal/pkg/conn"
	"bmail/internal/pkg/log"
	"bmail/internal/pkg/services"

	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/alecthomas/kingpin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/oklog/run"
)

var fqdn *string
var configPath *string
var app *kingpin.Application

func init() {

	app = kingpin.New("bmail", "An email gateway for Bitmessage")
	fqdn = app.Flag("fqdn", "Set FQDN.").Default("mail.nii236-work.local").String()
	configPath = app.Flag("config", "Set config path.").Default("./config.toml").String()
	log.NewToFile("./bmail.log")

}

func main() {
	fmt.Println("Bmail is running...")
	app.Parse(os.Args[1:])

	conn.New()
	config.New(*configPath)

	var g run.Group
	ctx, cancel := context.WithCancel(context.Background())

	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-ctx.Done():
			return ctx.Err()
		}
	}, func(error) {
		cancel()
	})

	g.Add(func() error {
		c := services.NewClean(ctx, &services.CleanOpts{
			Log:         log.NewToFile("./bmail-clean.log"),
			Name:        "clean",
			Description: "Removes messages that aren't bound for any registered BitMessage user",
			Quit:        make(chan bool),
			Stopped:     make(chan bool),
			Conn:        conn.Get(),
			Config:      config.Get(),
		})
		return c.Run()
	}, func(err error) {
		cancel()
	})
	g.Add(func() error {
		a := services.NewAgent(ctx)
		return a.Run()
	}, func(err error) {
		cancel()
	})
	g.Add(func() error {
		i := services.NewIncoming(ctx)
		return i.Run()
	}, func(err error) {
		cancel()
	})
	g.Add(func() error {
		o := services.NewOutgoing(ctx)
		return o.Run()
	}, func(err error) {
		cancel()
	})

	fmt.Println(g.Run())

}
