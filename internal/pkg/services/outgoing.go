package services

import (
	"bmail/db"
	"bmail/internal/pkg/bitmessage"
	"bmail/internal/pkg/config"
	"bmail/internal/pkg/conn"
	"bmail/internal/pkg/log"

	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/boil"
)

// Outgoing holds the service state
type Outgoing struct {
	bmclient    *bitmessage.Client
	log         *log.Logger
	name        string
	description string
	quit        chan bool
	stopped     chan bool
	conn        *sqlx.DB
	config      *config.C
	ctx         context.Context
}

// NewOutgoing creates a new service
func NewOutgoing(ctx context.Context) S {
	logger := log.NewToFile("./bmail-outgoing.log")
	bmclient, err := bitmessage.New(logger)
	if err != nil {
		logger.Fatal(err)
	}
	s := &Outgoing{
		ctx:         ctx,
		bmclient:    bmclient,
		log:         logger,
		name:        "outgoing",
		description: "Handles outgoing messages from a BitMessage user",
		quit:        make(chan bool),
		stopped:     make(chan bool),
		conn:        conn.Get(),
		config:      config.Get(),
	}
	return s
}

// Name of the service
func (s *Outgoing) Name() string {
	return s.name
}

// Description of the service
func (s *Outgoing) Description() string {
	return s.description
}

// Run will start the service
func (s *Outgoing) Run() error {
	s.log.Infow("Starting service",
		"name", s.name,
	)
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
		case <-s.ctx.Done():
			return s.ctx.Err()
		}
	}
}

// Stop will stop the service
func (s *Outgoing) Stop() {
	s.quit <- true
	<-s.stopped
}

func (s *Outgoing) checkMessages() error {
	msgs, err := s.bmclient.GetAllMessages()
	if err != nil {
		return err
	}
	for _, msg := range msgs {

		fmt.Println("Processing")
		switch msg.ToAddress {
		case s.config.Addresses.ReceivingAddress:
			fmt.Println("Processing Receiving")
		}
		msg := &db.ProcessedMessage{
			MessageID: msg.Msgid,
		}
		msg.Insert(s.conn, boil.Infer())

	}
	return nil
}
