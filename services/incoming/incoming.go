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
	"github.com/volatiletech/sqlboiler/queries/qm"
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
