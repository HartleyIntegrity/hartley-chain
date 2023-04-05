package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type PropertyContract struct {
	contractapi.Contract
}

type Property struct {
	ID             string  `json:"id"`
	Address        string  `json:"address"`
	Price          float64 `json:"price"`
	Tenant         string  `json:"tenant"`
	Owner          string  `json:"owner"`
	RentDuration   int     `json:"rent_duration"`
	Status         string  `json:"status"`
	StartDate      string  `json:"start_date"`
	EndDate        string  `json:"end_date"`
}

func (pc *PropertyContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	// Default properties for initial setup
	properties := []Property{...}

	for _, property := range properties {
		propertyBytes, _ := json.Marshal(property)
		err := ctx.GetStub().PutState(property.ID, propertyBytes)

		if err != nil {
			return fmt.Errorf("Failed to put property: %v", err)
		}
	}

	return nil
}

// Additional functions for adding, updating, and retrieving properties will be implemented
