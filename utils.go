package main

import (
	"encoding/json"
	"github.com/stripe/stripe-go/v72"
	"io/ioutil"
	"net/http"
)

func getEvent(w http.ResponseWriter, req * http.Request) (eventRes * stripe.Event, err error){
	const MaxBodyBytes = int64(65536)
	req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
	payload, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	event := stripe.Event{}
	err = json.Unmarshal(payload, &event)
	if err != nil {
		return nil, err
	}
	return &event, nil
}