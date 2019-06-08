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

var drop *bool
var fqdn *string
var log *zap.SugaredLogger

var app *kingpin.Application

func init() {

	app = kingpin.New("bmail", "An email gateway for Bitmessage")
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

	conn.MustExec("DROP TABLE IF EXISTS users;")
	conn.MustExec("DROP TABLE IF EXISTS processed_messages;")

}
