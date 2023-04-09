package tests

import (
	"encoding/json"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/stretchr/testify/assert"

	"github.com/HartleyIntegrity/hartley-chain/chaincode"
)

func TestCreateTenant(t *testing.T) {
	cc := new(chaincode.HartleyChaincode)
	stub := shim.NewMockStub("TestCreateTenant", cc)

	// Prepare the args
	args := []string{
		"createTenant",
		"1",
		"John Doe",
		"john.doe@example.com",
		"1234567890",
	}
	argsBytes := make([][]byte, len(args))
	for i, arg := range args {
		argsBytes[i] = []byte(arg)
	}

	// Invoke the CreateTenant function
	response := stub.MockInvoke("1", argsBytes)
	assert.Equal(t, int32(shim.OK), response.GetStatus(), "CreateTenant failed")

	// Check the saved tenant
	expectedTenant := &chaincode.Tenant{
		ObjectType: "tenant",
		ID:         "1",
		Name:       "John Doe",
		Email:      "john.doe@example.com",
		Phone:      "1234567890",
	}
	tenantBytes, _ := json.Marshal(expectedTenant)
	savedTenant, _ := stub.GetState("tenant_1")
	assert.Equal(t, tenantBytes, savedTenant, "Saved tenant does not match expected tenant")
}
