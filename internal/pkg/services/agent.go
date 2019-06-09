package services

import (
	"bmail/internal/pkg/config"
	"bmail/internal/pkg/conn"
	"bmail/internal/pkg/log"

	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

// Agent holds the service state
type Agent struct {
	log         *log.Logger
	name        string
	description string
	quit        chan bool
	conn        *sqlx.DB
	config      *config.C
	stopped     chan bool
	ctx         context.Context
}

// NewAgent creates a new service
func NewAgent(ctx context.Context) S {
	s := &Agent{
		ctx:         ctx,
		log:         log.NewToFile("./bmail-agent.log"),
		name:        "agent",
		description: "Admin RPC between BitMessage and Bmail",
		quit:        make(chan bool),
		stopped:     make(chan bool),
		conn:        conn.Get(),
		config:      config.Get(),
	}
	return s
}

// Name of the service
func (s *Agent) Name() string {
	return s.name
}

// Description of the service
func (s *Agent) Description() string {
	return s.description
}

// Run the service
func (s *Agent) Run() error {
	s.log.Infow("Starting service",
		"name", s.name,
	)
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			err := s.checkMessages()
			if err != nil {
				return err
			}
		case <-s.ctx.Done():
			return s.ctx.Err()

		}
	}
}

// Stop the service
func (s *Agent) Stop() {
	s.quit <- true
	<-s.stopped
}

func (s *Agent) checkMessages() error {
	return nil
}
