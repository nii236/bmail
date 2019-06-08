package incoming

import (
	"bmail/db"
	"bmail/internal/pkg/config"
	"bmail/internal/pkg/conn"
	"bmail/internal/pkg/log"
	"bmail/internal/pkg/maildir"
	"bmail/internal/pkg/service"
	"time"

	"github.com/jmoiron/sqlx"
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
		log:         log.NewToFile("./bmail-incoming.log"),
		name:        "incoming",
		description: "Handles incoming messages bound for a BitMessage user",
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
				s.checkNew()
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

func (s *Service) checkNew() error {
	s.log.Infow("Checking for new incoming mail...")
	users, err := db.Users().All(s.conn)
	if err != nil {
		return err
	}
	for _, u := range users {
		dir, err := maildir.Open(s.config.Storage.MailFolder, u.Username)
		if err != nil {
			return err
		}
		s.log.Infow("Opened dir", "username", u.Username, "dir", dir)

		newIDs, err := maildir.ReadNew(dir)
		if err != nil {
			return err
		}
		s.log.Infow("Found new mail", "ids", newIDs)
	}
	return nil
}
