package services

import (
	"bmail/internal/pkg/config"
	"bmail/internal/pkg/log"

	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

// Clean holds the service state
type Clean struct {
	log         *log.Logger
	name        string
	description string
	quit        chan bool
	conn        *sqlx.DB
	config      *config.C
	stopped     chan bool
	ctx         context.Context
}

// CleanOpts initialises the struct values
type CleanOpts struct {
	Log         *log.Logger
	Name        string
	Description string
	Quit        chan bool
	Stopped     chan bool
	Conn        *sqlx.DB
	Config      *config.C
}

// NewClean creates a new service
func NewClean(ctx context.Context, opts *CleanOpts) S {
	s := &Clean{
		ctx:         ctx,
		log:         opts.Log,
		name:        opts.Name,
		description: opts.Description,
		quit:        opts.Quit,
		stopped:     opts.Stopped,
		conn:        opts.Conn,
		config:      opts.Config,
	}
	return s
}

// Name of the service
func (s *Clean) Name() string {
	return s.name
}

// Description of the service
func (s *Clean) Description() string {
	return s.description
}

// Run the service
func (s *Clean) Run() error {
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
func (s *Clean) Stop() {
	s.quit <- true
	<-s.stopped
}
