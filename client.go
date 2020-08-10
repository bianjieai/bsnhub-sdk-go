package sdk

import (
	sdk "github.com/bianjieai/irita-sdk-go"
	"github.com/bianjieai/irita-sdk-go/types"

	"github.com/bianjieai/bsnhub-sdk-go/irismod/oracle"
)

func NewIRITAClient(cfg types.ClientConfig) sdk.IRITAClient {
	client := sdk.NewIRITAClient(cfg)
	return registerBSNClient(client)
}

func registerBSNClient(client sdk.IRITAClient) sdk.IRITAClient {
	oracleClient := oracle.NewClient(client.BaseClient, client.AppCodec())
	client.RegisterModule(oracleClient)
	return client
}
