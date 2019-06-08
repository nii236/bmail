package outgoing

import (
	"bmail/db"
	"bmail/internal/pkg/config"
	"bmail/internal/pkg/conn"
	"bmail/internal/pkg/log"
	"bmail/internal/pkg/service"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kolo/xmlrpc"
	"github.com/volatiletech/sqlboiler/boil"
)

// Service holds the service state
type Service struct {
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
	s := &Service{
		log:         log.NewToFile("./bmail-outgoing.log"),
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

func (s *Service) checkMessages() {
	client, err := xmlrpc.NewClient("http://devdev:devdev@localhost:8442", nil)
	if err != nil {
		s.log.Error(err)
		return
	}
	defer client.Close()
	type Result struct {
		InboxMessages []struct {
			EncodingType int    `json:"encodingType"`
			ToAddress    string `json:"toAddress"`
			Read         int    `json:"read"`
			Msgid        string `json:"msgid"`
			Message      string `json:"message"`
			FromAddress  string `json:"fromAddress"`
			ReceivedTime string `json:"receivedTime"`
			Subject      string `json:"subject"`
		} `json:"inboxMessages"`
	}

	var resultStr string
	err = client.Call("getAllInboxMessageIDs", nil, &resultStr)
	if err != nil {
		s.log.Error(err)
		return
	}
	result := &Result{}
	err = json.NewDecoder(strings.NewReader(resultStr)).Decode(result)
	if err != nil {
		s.log.Error(err)
		return
	}

	for _, msg := range result.InboxMessages {

		fmt.Println("Processing")
		switch msg.ToAddress {
		case s.config.Addresses.SendingAddress:
			fmt.Println("Processing Sending")
		case s.config.Addresses.ReceivingAddress:
			fmt.Println("Processing Receiving")
		case s.config.Addresses.DeregistrationAddress:
			fmt.Println("Processing Deregistration")
		case s.config.Addresses.RegistrationAddress:
			fmt.Println("Processing Registration")
		}
		msg := &db.ProcessedMessage{
			MessageID: msg.Msgid,
		}
		msg.Insert(s.conn, boil.Infer())

	}
}
