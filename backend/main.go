package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"hartley-chain/backend/handles"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/tenancies", handles.GetAllTenancies).Methods("GET")
	r.HandleFunc("/tenancies/{id}", handles.GetTenancy).Methods("GET")
	r.HandleFunc("/tenancies", handles.CreateTenancy).Methods("POST")
	r.HandleFunc("/tenancies/{id}", handles.UpdateTenancy).Methods("PUT")
	r.HandleFunc("/tenancies/{id}", handles.DeleteTenancy).Methods("DELETE")

	fmt.Println("Server started on port 8000")

	// Add CORS middleware
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	handler := handlers.CORS(headers, methods, origins)(r)

	// Start the server
	http.ListenAndServe(":8000", handler)
}
