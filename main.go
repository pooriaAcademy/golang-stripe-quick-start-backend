package main

import (
	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go/v72"
	"net/http"
)


func init() {
	stripe.Key = "sk_test_51IoMpOKNyEKTIrMReXkfEIsEZhTRIspduEAD7KzNhNgDn2A8MpWeQsxXxuNmcD5jEumM6eUnQAIUKIjOBCDmGmt600nMiLC9JI"
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/checkout", CORSCheck(CheckoutCreator))
	r.HandleFunc("/event", CORSCheck(HandleEvent))

	http.ListenAndServe(":8080", r)
}

