package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Property struct {
	// Property struct as defined in the smart contract
	ID           string  `json:"id"`
	Address      string  `json:"address"`
	Price        float64 `json:"price"`
	Tenant       string  `json:"tenant"`
	Owner        string  `json:"owner"`
	RentDuration int     `json:"rent_duration"`
	Status       string  `json:"status"`
	StartDate    string  `json:"start_date"`
	EndDate      string  `json:"end_date"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/properties", GetProperties).Methods("GET")
	router.HandleFunc("/properties", AddProperty).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetProperties(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Use the Fabric Gateway SDK to interact with the network
	// and get properties from the ledger
}

func AddProperty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Use the Fabric Gateway SDK to interact with the network
	// and add a property to the ledger
}
