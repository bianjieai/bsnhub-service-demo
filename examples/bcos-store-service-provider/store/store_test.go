package store_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	bcosclient "github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"

	"github.com/bianjieai/bsnhub-service-demo/examples/bcos-store-service-provider/store/bcos"
)

type StoreTestSuite struct {
	suite.Suite
	client *bcosclient.Client
}

func TestStoreTestSuite(t *testing.T) {
	suite.Run(t, new(StoreTestSuite))
}

func (suite *StoreTestSuite) SetupTest() {
	configs, err := conf.ParseConfigFile("config.toml")
	suite.NoError(err)

	client, err := bcosclient.Dial(&configs[0])
	suite.NoError(err)

	suite.client = client
}

func (suite *StoreTestSuite) TestDeployStoreContract() {
	// deploy store contract
	address, tx, _, err := bcos.DeployStore(suite.client.GetTransactOpts(), suite.client)
	suite.NoError(err)

	fmt.Printf("store contract deployed, contract address: %s, tx hash: %s", address.Hex(), tx.Hash().Hex())
}
