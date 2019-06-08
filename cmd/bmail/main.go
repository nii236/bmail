package main

import (
	"bmail/internal/pkg/config"
	"bmail/internal/pkg/conn"
	"bmail/internal/pkg/log"
	"bmail/internal/pkg/service"
	"bmail/services/agent"
	"bmail/services/clean"
	"bmail/services/incoming"
	"bmail/services/outgoing"
	"os"
	"os/signal"
	"sync"

	"github.com/alecthomas/kingpin"
	_ "github.com/mattn/go-sqlite3"
)

var fqdn *string

var app *kingpin.Application

func init() {

	app = kingpin.New("bmail", "An email gateway for Bitmessage")
	fqdn = app.Flag("fqdn", "Set FQDN.").Default("mail.nii236-work.local").String()
	log.New(false)

}

func main() {
	app.Parse(os.Args[1:])

	conn.New()
	config.New()

	wg := sync.WaitGroup{}
	wg.Add(1)

	c := clean.New()
	c.Start()
	i := incoming.New()
	i.Start()
	o := outgoing.New()
	o.Start()
	a := agent.New()
	a.Start()

	signalClose(c, i, o, a)

	wg.Wait()

}

func signalClose(services ...service.S) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		select {
		case <-c:

			log.Info("Received interrupt, stopping all services...")
			for _, s := range services {
				s.Stop()
			}
			log.Info("Stopped.")
			os.Exit(0)
		}
	}()
}
