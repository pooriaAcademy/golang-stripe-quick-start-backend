package main



import(
	"github.com/gorilla/mux"
)




func main(){
	r := mux.NewRouter()
	r.HandleFunc("/checkout", CheckoutCreator)
}














