package outgoing

import (
	"bmail/db"
	"bmail/internal/pkg/bitmessage"
	"bmail/internal/pkg/config"
	"bmail/internal/pkg/conn"
	"bmail/internal/pkg/log"
	"bmail/internal/pkg/service"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/boil"
)

// Service holds the service state
type Service struct {
	bmclient    *bitmessage.Client
	log         *log.Logger
	name        string
	description string
	quit        chan bool
	stopped     chan bool
	conn        *sqlx.DB
	config      *config.C
}

// New creates a new service
func New() service.S {
	logger := log.NewToFile("./bmail-outgoing.log")
	bmclient, err := bitmessage.New(logger)
	if err != nil {
		logger.Fatal(err)
	}
	s := &Service{
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

// Name returns the name of the service
func (s *Service) Name() string {
	return s.name
}

// Description returns the description of the service
func (s *Service) Description() string {
	return s.description
}

// Start will start the service
func (s *Service) Start() {
	s.log.Infow("Starting service",
		"name", s.name,
	)
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-ticker.C:
			case <-s.quit:
				s.stopped <- true
				return
			}
		}
	}()
}

// Stop will stop the service
func (s *Service) Stop() {
	s.quit <- true
	<-s.stopped
}

func (s *Service) checkMessages() error {
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
