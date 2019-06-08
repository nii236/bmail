package clean

import (
	"bmail/internal/pkg/config"
	"bmail/internal/pkg/conn"
	"bmail/internal/pkg/log"
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
	conn        *sqlx.DB
	config      *config.C
	stopped     chan bool
}

// New creates a new service
func New() service.S {
	s := &Service{
		log:         log.NewToFile("./bmail-clean.log"),
		name:        "clean",
		description: "Removes messages that aren't bound for any registered BitMessage user",
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
