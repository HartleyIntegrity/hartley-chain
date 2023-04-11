package handles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"hartley-chain/backend/models"
)

var (
	tenancies     = make(map[string]models.Tenancy)
	tenanciesLock sync.RWMutex
)

func GetAllTenancies(w http.ResponseWriter, r *http.Request) {
	tenanciesLock.RLock()
	defer tenanciesLock.RUnlock()

	var tenancyList []models.Tenancy
	for _, t := range tenancies {
		tenancyList = append(tenancyList, t)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tenancyList)
}

func GetTenancy(w http.ResponseWriter, r *http.Request) {
	tenanciesLock.RLock()
	defer tenanciesLock.RUnlock()

	vars := mux.Vars(r)
	id := vars["id"]

	if t, ok := tenancies[id]; ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(t)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func CreateTenancy(w http.ResponseWriter, r *http.Request) {
	tenanciesLock.Lock()
	defer tenanciesLock.Unlock()

	var t models.Tenancy
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	t.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	tenancies[t.ID] = t

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

func UpdateTenancy(w http.ResponseWriter, r *http.Request) {
	tenanciesLock.Lock()
	defer tenanciesLock.Unlock()

	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := tenancies[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var t models.Tenancy
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	t.ID = id
	tenancies[id] = t

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(t)
}

func DeleteTenancy(w http.ResponseWriter, r *http.Request) {
	tenanciesLock.Lock()
	defer tenanciesLock.Unlock()

	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := tenancies[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	delete(tenancies, id)
	w.WriteHeader(http.StatusOK)
}
