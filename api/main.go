package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HartleyIntegrity/hartley-chain/api/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/tenancy", handlers.CreateTenancy).Methods("POST")
	router.HandleFunc("/tenancy/{id}", handlers.GetTenancy).Methods("GET")
	router.HandleFunc("/tenancy/{id}", handlers.UpdateTenancy).Methods("PUT")
	router.HandleFunc("/tenancy/{id}", handlers.DeleteTenancy).Methods("DELETE")

	fmt.Println("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
