package main

import (
	"bmail/db"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/kolo/xmlrpc"
	"github.com/volatiletech/sqlboiler/boil"
)

type Controller struct {
	Event chan string
}

func Listen() {
	wg := &sync.WaitGroup{}
	for {
		wg.Add(1)
		fmt.Println("CHECKING")
		CheckMessages()
		wg.Done()
		wg.Wait()
		time.Sleep(1 * time.Second)
	}
}

func CheckMessages() {
	client, err := xmlrpc.NewClient("http://devdev:devdev@localhost:8442", nil)
	if err != nil {
		log.Error(err)
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
		log.Error(err)
		return
	}
	result := &Result{}
	err = json.NewDecoder(strings.NewReader(resultStr)).Decode(result)
	if err != nil {
		log.Error(err)
		return
	}

	for _, msg := range result.InboxMessages {
		config := GetConfig()
		fmt.Println("Processing")
		switch msg.ToAddress {
		case config.Addresses.SendingAddress:
			fmt.Println("Processing Sending")
		case config.Addresses.ReceivingAddress:
			fmt.Println("Processing Receiving")
		case config.Addresses.DeregistrationAddress:
			fmt.Println("Processing Deregistration")
		case config.Addresses.RegistrationAddress:
			fmt.Println("Processing Registration")
		}
		msg := &db.ProcessedMessage{
			MessageID: msg.Msgid,
		}
		msg.Insert(conn, boil.Infer())

	}
}
