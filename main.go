package main

import (
	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go/v72"
	"net/http"
)

func CORSCheck(handler func (w http.ResponseWriter, req *http.Request)) func (w http.ResponseWriter, req *http.Request){
	res := func (w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Max-Age", "3600")
		if req.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		handler(w, req)
	}
	return res
}

func init() {
	stripe.Key = "sk_test_51IoMpOKNyEKTIrMReXkfEIsEZhTRIspduEAD7KzNhNgDn2A8MpWeQsxXxuNmcD5jEumM6eUnQAIUKIjOBCDmGmt600nMiLC9JI"
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/checkout", CORSCheck(CheckoutCreator))
	r.HandleFunc("/event", CORSCheck(HandleEvent))

	http.ListenAndServe(":8080", r)
}

