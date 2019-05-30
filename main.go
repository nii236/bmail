package main

import (
	"flag"

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
`

var dev bool
var log *zap.SugaredLogger

func init() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // flushes buffer, if any
	log = logger.Sugar()

	flag.BoolVar(&dev, "dev", false, "Drop DB on start")
}

// RegisterUser is a prepared query for the database
var RegisterUser *sqlx.NamedStmt

// UnregisterUser is a prepared query for the database
var UnregisterUser *sqlx.NamedStmt

func main() {
	flag.Parse()
	log.Info("Bmail starting")
	conn, err := sqlx.Connect("sqlite3", "bitportal.db")
	if err != nil {
		log.Fatal(err)
	}

	if dev {
		conn.MustExec("DROP TABLE IF EXISTS users;")
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

}
