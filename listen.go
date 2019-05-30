package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kolo/xmlrpc"
)

type Controller struct {
	Event chan string
}

func Listen() {
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			log.Info("Run ticker")
			_, err := os.Stat("/home/nii236/newmessage.lock")
			if err != nil {
				continue
			}
			err = os.Remove("/home/nii236/newmessage.lock")
			if err != nil {
				log.Error(err)
				continue
			}
			CheckMessages()
		}
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
	err = client.Call("getAllInboxMessages", nil, &resultStr)
	if err != nil {
		log.Error(err)
		return
	}
	result := &Result{}
	json.NewDecoder(strings.NewReader(resultStr)).Decode(result)
	// fmt.Println("RPC RESPONSE:", resultStr)
	// fmt.Println("RPC RESPONSE:", result)

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
			_, err = UnregisterUser.Exec(map[string]interface{}{
				"bitmessage_id": msg.FromAddress,
			})
		case config.Addresses.RegistrationAddress:
			fmt.Println("Processing Registration")
			_, err = RegisterUser.Exec(map[string]interface{}{
				"username":      msg.Subject,
				"bitmessage_id": msg.FromAddress,
			})
			if err != nil {
				log.Error(err)
				return
			}
		}
	}
}
