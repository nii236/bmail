package conn

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var conn *sqlx.DB

// New creates a new connection to the DB
func New() {
	var err error
	if conn != nil {
		log.Fatal("db conn already established")
	}
	conn, err = sqlx.Connect("sqlite3", "bmail.db")
	if err != nil {
		log.Fatal(err)
	}
}

// Get returns the singleton pointer to the DB
func Get() *sqlx.DB {
	if conn == nil {
		log.Fatal("db conn not established")
	}
	return conn
}
