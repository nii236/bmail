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
	user_id    			INTEGER UNIQUE NOT NULL PRIMARY KEY,
	username			VARCHAR(40) UNIQUE NOT NULL,
	bitmessage_id      	VARCHAR(40) UNIQUE NOT NULL,
	archived 			BOOL NOT NULL DEFAULT false,
	created_at			TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS messages (
	encoding_type 	INTEGER,
	to_address    	VARCHAR(200),
	read         	INTEGER,
	msgid        	VARCHAR(200),
	message      	VARCHAR(200),
	from_address  	VARCHAR(200),
	received_time 	VARCHAR(200),
	subject      	VARCHAR(200),
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

	app.Command("listen", "Start the gateway")
	app.Command("trigger", "Trigger the gateway to check for new messages").Default()
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // flushes buffer, if any
	log = logger.Sugar()

}

// RegisterUser is a prepared query for the database
var RegisterUser *sqlx.NamedStmt

// UnregisterUser is a prepared query for the database
var UnregisterUser *sqlx.NamedStmt

func Trigger() {
	os.Create("/home/nii236/newmessage.lock")
}

func main() {

	log.Info("Bmail starting")
	conn, err := sqlx.Connect("sqlite3", "bitportal.db")
	if err != nil {
		log.Fatal(err)
	}

	if *dev {
		conn.MustExec("DROP TABLE IF EXISTS users;")
		conn.MustExec("DROP TABLE IF EXISTS messages;")
	}
	conn.MustExec(schema)
	RegisterUser, err = conn.PrepareNamed("INSERT INTO users (username, bitmessage_id) VALUES (:username, :bitmessage_id)")
	if err != nil {
		log.Fatal(err)
	}
	UnregisterUser, err = conn.PrepareNamed("UPDATE users SET archived=true WHERE bitmessage_id = :bitmessage_id")
	if err != nil {
		log.Fatal(err)
	}

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case "listen":
		Listen()
	case "notify":
		Trigger()
	default:
		Trigger()
	}

}
