package main

import "github.com/nats-io/nats-server/server"

type Controller struct {
	Event chan string
}

func NewServer() *Controller {
	s := server.New()
	s.Start()
	return &Controller{}
}

func Listen() {
	for {
		select {
		case ev := <-c.Event:
			switch ev {
			case "startingUp":
			case "newMessage":
			case "newBroadcast":
			default:
				log.Error("Unrecognised event:", ev)
			}
		}
	}
}

func HandleRegistration(from, subject, body string) error {
	_, err := RegisterUser.Exec(map[string]interface{}{
		"username":      subject,
		"bitmessage_id": from,
	})
	return err
}
func HandleDeregistration(from, subject, body string) error {
	_, err := UnregisterUser.Exec(map[string]interface{}{
		"bitmessage_id": from,
	})
	return err
}
func HandleReceiving(from, subject, body string) error {
	return nil
}
func HandleBugReport(from, subject, body string) error {
	return nil
}
