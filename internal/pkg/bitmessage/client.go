package bitmessage

import (
	"bmail/internal/pkg/log"
	"encoding/json"

	"strings"

	"github.com/kolo/xmlrpc"
)

// Client is the XMLRPC client for the BitMessage API
type Client struct {
	*xmlrpc.Client
	log *log.Logger
}

// New returns the XMLRPC client for the BitMessage API
func New(l *log.Logger) (*Client, error) {
	c, err := xmlrpc.NewClient("http://devdev:devdev@localhost:8442", nil)
	if err != nil {
		return nil, err
	}
	return &Client{c, l}, nil

}

// InboxMessagesResp is the JSON response for fetching all messages
type InboxMessagesResp struct {
	InboxMessages []*InboxMessages `json:"inboxMessages"`
}

// InboxMessages is the JSON subkey response for fetching all messages
type InboxMessages struct {
	EncodingType int    `json:"encodingType"`
	ToAddress    string `json:"toAddress"`
	Read         int    `json:"read"`
	Msgid        string `json:"msgid"`
	Message      string `json:"message"`
	FromAddress  string `json:"fromAddress"`
	ReceivedTime string `json:"receivedTime"`
	Subject      string `json:"subject"`
}

// GetAllMessages returns all messages in BitMessage
func (c *Client) GetAllMessages() ([]*InboxMessages, error) {
	var resultStr string
	err := c.Call("getAllInboxMessageIDs", nil, &resultStr)
	if err != nil {
		return nil, err
	}
	result := &InboxMessagesResp{}
	err = json.NewDecoder(strings.NewReader(resultStr)).Decode(result)
	if err != nil {
		return nil, err
	}
	return result.InboxMessages, nil
}
