package services

import (
	"bmail/db"
	"bmail/internal/pkg/bmaildir"
	"bmail/internal/pkg/config"
	"bmail/internal/pkg/conn"
	"bmail/internal/pkg/log"

	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// Incoming holds the service state
type Incoming struct {
	log         *log.Logger
	name        string
	description string
	quit        chan bool
	stopped     chan bool
	conn        *sqlx.DB
	config      *config.C
	ctx         context.Context
}

// NewIncoming creates a new service
func NewIncoming(ctx context.Context) S {
	s := &Incoming{
		ctx:         ctx,
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

// Name of the service
func (s *Incoming) Name() string {
	return s.name
}

// Description of the service
func (s *Incoming) Description() string {
	return s.description
}

// Run the service
func (s *Incoming) Run() error {
	s.log.Infow("Starting service",
		"name", s.name,
	)

	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			err := s.checkNew()
			if err != nil {
				return err
			}
		case <-s.ctx.Done():
			return s.ctx.Err()
		}
	}
}

// Stop will stop the service
func (s *Incoming) Stop() {
	s.quit <- true
	<-s.stopped
}

func (s *Incoming) checkNew() error {
	s.log.Infow("Checking for new incoming mail...")
	dir, err := bmaildir.Open(s.config.Storage.MailFolder)
	if err != nil {
		return err
	}
	s.log.Infow("Opened maildir", "dir", dir)

	newIDs, err := bmaildir.ReadNew(dir)
	if err != nil {
		return err
	}
	s.log.Infow("Found new mail", "ids", newIDs)
	for _, id := range newIDs {
		msg, err := dir.Message(id)
		if err != nil {
			return err
		}
		to := msg.Header.Get("To")
		exists, err := db.Users(qm.Where("username = ?", to)).Exists(s.conn)
		if err != nil {
			return err
		}
		if !exists {
			s.log.Debug("recipient for email does not exist, continuing...")
			continue
		}

		// Forward message to BitMessage user here
	}

	return nil
}
