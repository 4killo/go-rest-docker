package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/4killo/go-rest-docker/data"
	"github.com/4killo/go-rest-docker/service"
	"github.com/gorilla/mux"

	_ "net/http/pprof"

	_ "github.com/go-sql-driver/mysql"
)

func insertCard(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var paymentRequest data.Request
	b, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(b, &paymentRequest)
	//fmt.Println("request:", paymentRequest)
	st, err := service.Insert(&paymentRequest)
	if err != nil {
		res.WriteHeader(st)
	}
	res.WriteHeader(st)

	// fmt.Println("request:", paymentRequest)
	// j, _ := json.Marshal(paymentRequest)
	// fmt.Println(paymentRequest)
	// res.Write(j)

}

func getCards(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var paymentRequest data.Request
	b, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(b, &paymentRequest)

	//we can use this log form or below encoding
	// j, _ := json.Marshal(GetCards(&paymentRequest))
	// fmt.Println(paymentRequest)
	// res.Write(j)

	json.NewEncoder(res).Encode(service.GetCards(&paymentRequest))

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/payment/cards/add", insertCard).Methods("POST")
	r.HandleFunc("/payment/cards", getCards).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
