// api/main.go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (app *App) Initialize() {
	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

func (app *App) Run(addr string) {
	fmt.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func (app *App) initializeRoutes() {
	app.Router.HandleFunc("/properties", app.getProperties).Methods("GET")
	app.Router.HandleFunc("/properties/{id}", app.getProperty).Methods("GET")
}

// api/main.go (continued)

func (app *App) getProperties(w http.ResponseWriter, r *http.Request) {
	fabricClient, err := NewFabricClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := fabricClient.channelClient.Query("GetAllProperties", [][]byte{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	properties := make([]Property, 0)
	err = json.Unmarshal(response.Payload, &properties)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(properties)
}

func (app *App) getProperty(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fabricClient, err := NewFabricClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := fabricClient.channelClient.Query("ReadProperty", [][]byte{[]byte(id)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var property Property
	err = json.Unmarshal(response.Payload, &property)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(property)
}

func main() {
	app := App{}
	app.Initialize()
	app.Run(":8000")
}
