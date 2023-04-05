package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

// TenancyContract is the main structure of our smart contract
type TenancyContract struct {
}

// Tenancy represents a property rental agreement
type Tenancy struct {
	ID             string    `json:"id"`
	Landlord       string    `json:"landlord"`
	Tenant         string    `json:"tenant"`
	PropertyID     string    `json:"property_id"`
	RentAmount     float64   `json:"rent_amount"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	ContractSigned bool      `json:"contract_signed"`
}

func main() {
	err := shim.Start(new(TenancyContract))
	if err != nil {
		fmt.Printf("Error starting Tenancy contract: %s", err)
	}
}

// Init initializes the smart contract
func (t *TenancyContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

// Invoke is called for every transaction
func (t *TenancyContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	// Route to the appropriate function
	if function == "createTenancy" {
		return t.createTenancy(stub, args)
	} else if function == "getTenancy" {
		return t.getTenancy(stub, args)
	} else if function == "updateTenancy" {
		return t.updateTenancy(stub, args)
	} else if function == "deleteTenancy" {
		return t.deleteTenancy(stub, args)
	}

	return shim.Error("Invalid function name.")
}

// createTenancy creates a new tenancy agreement
func (t *TenancyContract) createTenancy(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8.")
	}

	rentAmount, err := strconv.ParseFloat(args[4], 64)
	if err != nil {
		return shim.Error("Rent amount must be a valid float.")
	}

	startDate, err := time.Parse(time.RFC3339, args[5])
	if err != nil {
		return shim.Error("Start date must be in RFC3339 format.")
	}

	endDate, err := time.Parse(time.RFC3339, args[6])
	if err != nil {
		return shim.Error("End date must be in RFC3339 format.")
	}

	contractSigned, err := strconv.ParseBool(args[7])
	if err != nil {
		return shim.Error("Contract signed must be true or false.")
	}

	tenancy := Tenancy{
		ID:             args[0],
		Landlord:       args[1],
		Tenant:         args[2],
		PropertyID:     args[3],
		RentAmount:     rentAmount,
		StartDate:      startDate,
		EndDate:        endDate,
		ContractSigned: contractSigned,
	}

	tenancyAsBytes, _ := json.Marshal(tenancy)
	err = stub.PutState(tenancy.ID, tenancyAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to create tenancy: %s", args[0]))
	}

	return shim.Success(nil)
}

// getTenancy retrieves a tenancy agreement by ID
func (t *TenancyContract) getTenancy(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}

	tenancyAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get tenancy.")
	}

	return shim.Success(tenancyAsBytes)
}

// updateTenancy updates an existing tenancy agreement
func (t *TenancyContract) updateTenancy(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8.")
	}

	tenancyAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get tenancy.")
	}

	tenancy := Tenancy{}
	_ = json.Unmarshal(tenancyAsBytes, &tenancy)

	// Update fields
	tenancy.Landlord = args[1]
	tenancy.Tenant = args[2]
	tenancy.PropertyID = args[3]
	tenancy.RentAmount, _ = strconv.ParseFloat(args[4], 64)
	tenancy.StartDate, _ = time.Parse(time.RFC3339, args[5])
	tenancy.EndDate, _ = time.Parse(time.RFC3339, args[6])
	tenancy.ContractSigned, _ = strconv.ParseBool(args[7])

	tenancyAsBytes, _ = json.Marshal(tenancy)
	err = stub.PutState(tenancy.ID, tenancyAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to update tenancy: %s",
			args[0]))
	}

	return shim.Success(nil)
}

// deleteTenancy deletes a tenancy agreement by ID
func (t *TenancyContract) deleteTenancy(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}

	err := stub.DelState(args[0])
	if err != nil {
		return shim.Error("Failed to delete tenancy.")
	}

	return shim.Success(nil)
}
