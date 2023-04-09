package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/stretchr/testify/assert"

	"github.com/HartleyIntegrity/hartley-chain/chaincode"
)

func TestCreateProperty(t *testing.T) {
	cc := new(chaincode.HartleyChaincode)
	stub := shim.NewMockStub("TestCreateProperty", cc)

	// Prepare the args
	args := []string{
		"createProperty",
		"1",
		"123 Main St",
		"London",
		"United Kingdom",
		"100000",
	}
	argsBytes := make([][]byte, len(args))
	for i, arg := range args {
		argsBytes[i] = []byte(arg)
	}

	// Invoke the CreateProperty function
	response := stub.MockInvoke("1", argsBytes)
	assert.Equal(t, int32(shim.OK), response.GetStatus(), "CreateProperty failed")

	// Check the saved property
	expectedProperty := &chaincode.Property{
		ObjectType: "property",
		ID:         "1",
		Address:    "123 Main St",
		City:       "London",
		Country:    "United Kingdom",
		Price:      100000,
	}
	propertyBytes, _ := json.Marshal(expectedProperty)
	savedProperty, _ := stub.GetState("property_1")
	assert.Equal(t, propertyBytes, savedProperty, "Saved property does not match expected property")
}
