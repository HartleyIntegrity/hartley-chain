// smart-contracts/property-contract/property-contract.go

package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Property struct {
	ID          string    `json:"id"`
	Address     string    `json:"address"`
	Owner       string    `json:"owner"`
	Tenant      string    `json:"tenant"`
	Rent        float64   `json:"rent"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	IsOccupied  bool      `json:"is_occupied"`
	LastUpdated time.Time `json:"last_updated"`
}

type PropertyContract struct {
	contractapi.Contract
}

func (pc *PropertyContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	properties := []Property{
		{
			ID:          "property1",
			Address:     "123 Hartley Street",
			Owner:       "Alice",
			Tenant:      "",
			Rent:        1000.00,
			StartDate:   time.Time{},
			EndDate:     time.Time{},
			IsOccupied:  false,
			LastUpdated: time.Now(),
		},
	}

	for _, property := range properties {
		propertyJSON, err := json.Marshal(property)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(property.ID, propertyJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state: %v", err)
		}
	}

	return nil
}

// smart-contracts/property-contract/property-contract.go (continued)

func (pc *PropertyContract) CreateProperty(ctx contractapi.TransactionContextInterface, id string, address string, owner string) error {
	exists, err := pc.PropertyExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the property %s already exists", id)
	}

	property := Property{
		ID:          id,
		Address:     address,
		Owner:       owner,
		Tenant:      "",
		Rent:        0,
		StartDate:   time.Time{},
		EndDate:     time.Time{},
		IsOccupied:  false,
		LastUpdated: time.Now(),
	}
	propertyJSON, err := json.Marshal(property)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, propertyJSON)
}

func (pc *PropertyContract) ReadProperty(ctx contractapi.TransactionContextInterface, id string) (*Property, error) {
	propertyJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if propertyJSON == nil {
		return nil, fmt.Errorf("the property %s does not exist", id)
	}

	var property Property
	err = json.Unmarshal(propertyJSON, &property)
	if err != nil {
		return nil, err
	}

	return &property, nil
}

func (pc *PropertyContract) UpdateProperty(ctx contractapi.TransactionContextInterface, id string, address string, owner string) error {
	property, err := pc.ReadProperty(ctx, id)
	if err != nil {
		return err
	}

	property.Address = address
	property.Owner = owner
	property.LastUpdated = time.Now()

	propertyJSON, err := json.Marshal(property)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, propertyJSON)
}

func (pc *PropertyContract) DeleteProperty(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := pc.PropertyExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the property %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

func (pc *PropertyContract) PropertyExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	propertyJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return propertyJSON != nil, nil
}

// smart-contracts/property-contract/property-contract.go (continued)

func (pc *PropertyContract) SignTenancyAgreement(ctx contractapi.TransactionContextInterface, id string, tenant string, rent float64, startDate string, endDate string) error {
	property, err := pc.ReadProperty(ctx, id)
	if err != nil {
		return err
	}
	if property.IsOccupied {
		return fmt.Errorf("the property %s is already occupied", id)
	}

	start, err := time.Parse(time.RFC3339, startDate)
	if err != nil {
		return fmt.Errorf("failed to parse start date: %v", err)
	}

	end, err := time.Parse(time.RFC3339, endDate)
	if err != nil {
		return fmt.Errorf("failed to parse end date: %v", err)
	}

	property.Tenant = tenant
	property.Rent = rent
	property.StartDate = start
	property.EndDate = end
	property.IsOccupied = true
	property.LastUpdated = time.Now()

	propertyJSON, err := json.Marshal(property)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, propertyJSON)
}

func (pc *PropertyContract) TerminateTenancyAgreement(ctx contractapi.TransactionContextInterface, id string) error {
	property, err := pc.ReadProperty(ctx, id)
	if err != nil {
		return err
	}
	if !property.IsOccupied {
		return fmt.Errorf("the property %s is not occupied", id)
	}

	property.Tenant = ""
	property.Rent = 0
	property.StartDate = time.Time{}
	property.EndDate = time.Time{}
	property.IsOccupied = false
	property.LastUpdated = time.Now()

	propertyJSON, err := json.Marshal(property)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, propertyJSON)
}

// smart-contracts/property-contract/property-contract.go (continued)

func main() {
	propertyChaincode, err := contractapi.NewChaincode(&PropertyContract{})
	if err != nil {
		fmt.Printf("Error creating property chaincode: %s", err.Error())
		return
	}

	if err := propertyChaincode.Start(); err != nil {
		fmt.Printf("Error starting property chaincode: %s", err.Error())
	}
}
