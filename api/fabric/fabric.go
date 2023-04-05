package fabric

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

var sdk *fabsdk.FabricSDK

// Initialize initializes the SDK
func Initialize() error {
	var err error
	configPath := "config/config.yaml" // Path to your config.yaml
	sdk, err = fabsdk.New(config.FromFile(configPath))
	if err != nil {
		return fmt.Errorf("failed to create new SDK: %v", err)
	}

	return nil
}

// InvokeChaincode invokes a chaincode function
func InvokeChaincode(function string, args []string) (string, error) {
	clientContext := sdk.ChannelContext("mychannel", fabsdk.WithUser("User1"), fabsdk.WithOrg("Org1"))
	client, err := channel.New(clientContext)
	if err != nil {
		return "", fmt.Errorf("failed to create new channel client: %v", err)
	}

	response, err := client.Execute(channel.Request{
		ChaincodeID: "tenancy",
		Fcn:         function,
		Args:        convertArgs(args),
	})
	if err != nil {
		return "", fmt.Errorf("failed to invoke chaincode: %v", err)
	}

	return string(response.Payload), nil
}

// QueryChaincode queries the chaincode state
func QueryChaincode(function string, args []string) ([]byte, error) {
	clientContext := sdk.ChannelContext("mychannel", fabsdk.WithUser("User1"), fabsdk.WithOrg("Org1"))
	client, err := channel.New(clientContext)
	if err != nil {
		return nil, fmt.Errorf("failed to create new channel client: %v", err)
	}

	response, err := client.Query(channel.Request{
		ChaincodeID: "tenancy",
		Fcn:         function,
		Args:        convertArgs(args),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query chaincode: %v", err)
	}

	return response.Payload, nil
}

func convertArgs(args []string) [][]byte {
	var bytesArgs [][]byte
	for _, arg := range args {
		bytesArgs = append(bytesArgs, []byte(arg))
	}
	return bytesArgs
}
