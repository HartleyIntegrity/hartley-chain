package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/HartleyIntegrity/hartley-chain/api/fabric"
	"github.com/gorilla/mux"
)

// TenancyRequest represents a request for a tenancy
type TenancyRequest struct {
	ID             string  `json:"id"`
	Landlord       string  `json:"landlord"`
	Tenant         string  `json:"tenant"`
	PropertyID     string  `json:"propertyId"`
	RentAmount     float64 `json:"rentAmount"`
	StartDate      string  `json:"startDate"`
	EndDate        string  `json:"endDate"`
	ContractSigned bool    `json:"contractSigned"`
}

func CreateTenancy(w http.ResponseWriter, r *http.Request) {
	var tenancyReq TenancyRequest
	err := json.NewDecoder(r.Body).Decode(&tenancyReq)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	args := []string{
		tenancyReq.ID,
		tenancyReq.Landlord,
		tenancyReq.Tenant,
		tenancyReq.PropertyID,
		fmt.Sprintf("%f", tenancyReq.RentAmount),
		tenancyReq.StartDate,
		tenancyReq.EndDate,
		strconv.FormatBool(tenancyReq.ContractSigned),
	}

	_, err = fabric.InvokeChaincode("createTenancy", args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetTenancy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	tenancyAsBytes, err := fabric.QueryChaincode("getTenancy", []string{id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(tenancyAsBytes)
}

func UpdateTenancy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var tenancyReq TenancyRequest
	err := json.NewDecoder(r.Body).Decode(&tenancyReq)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	args := []string{
		id,
		tenancyReq.Landlord,
		tenancyReq.Tenant,
		tenancyReq.PropertyID,
		fmt.Sprintf("%f", tenancyReq.RentAmount),
		tenancyReq.StartDate,
		tenancyReq.EndDate,
		strconv.FormatBool(tenancyReq.ContractSigned),
	}

	_, err = fabric.InvokeChaincode("updateTenancy", args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteTenancy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := fabric.InvokeChaincode("deleteTenancy", []string{id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
