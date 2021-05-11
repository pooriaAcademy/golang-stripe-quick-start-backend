package main

import (
	"github.com/stripe/stripe-go/v72/customer"
	"log"
	"net/http"
)


func HandleEvent(w http.ResponseWriter, req * http.Request) error {
	event, err := getEvent(w, req)

	if err != nil{
		return err
	}

	log.Println(event.Type)

	if event.Type == "customer.subscription.created" {
		c, err := customer.Get(event.Data.Object["FinalEmail"].(string), nil)
		if err != nil {
			return err
		}
		email := c.Metadata["FinalEmail"]
		log.Println("Subscription created by", email)
	}


	return nil
}
