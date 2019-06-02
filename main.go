package main

import (
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
	id	    			INTEGER UNIQUE NOT NULL PRIMARY KEY,
	username			VARCHAR(40) UNIQUE NOT NULL,
	bitmessage_id      	VARCHAR(40) UNIQUE NOT NULL,
	archived 			BOOL NOT NULL DEFAULT false,
	created_at			TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS processed_messages (
	id    			INTEGER UNIQUE NOT NULL PRIMARY KEY,
	message_id      VARCHAR(200) NOT NULL,
	processed		BOOL NOT NULL DEFAULT FALSE
);
`

var dev *bool
var log *zap.SugaredLogger

var app *kingpin.Application

func init() {

	NewConfig()

	app = kingpin.New("bmail", "An email gateway for Bitmessage")
	dev = app.Flag("dev", "Enable dev mode.").Bool()

	app.Command("listen", "Start the gateway").Default()
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // flushes buffer, if any
	log = logger.Sugar()

}

func main() {
	app.Parse(os.Args[1:])
	log.Info("Bmail starting")
	conn, err := sqlx.Connect("sqlite3", "bmail.db")
	if err != nil {
		log.Fatal(err)
	}
	if *dev {
		conn.MustExec("DROP TABLE IF EXISTS users;")
		conn.MustExec("DROP TABLE IF EXISTS processed_messages;")
	}

	conn.MustExec(schema)

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case "listen":
		fallthrough
	default:
		Listen()
	}

}
