package service

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	ethcmn "github.com/ethereum/go-ethereum/common"

	"github.com/bianjieai/bsnhub-service-demo/common"
)

func TestMintNft(t *testing.T) {
	config, err := common.LoadYAMLConfig("../config/config.yaml")
	require.Nil(t, err)

	serviceInstance := MakeService(config)
	serviceInstance.Logger = common.Logger

	to := ethcmn.HexToAddress("0xA0214bC81667eFe7C8A703C30d145fFd89daBD58")
	amount := big.NewInt(1)
	metaID := "-Z-2fJxzCoFJ0MOU-zA3-tiIh7dK6FjDruAxgxW6PEs"
	setPrice := big.NewInt(10000000000000000)
	isForSale := true

	nftID, err := serviceInstance.MintNft(to, amount, metaID, setPrice, isForSale)
	require.Nil(t, err)

	fmt.Printf("NFT %s minted successfully: https://rinkeby.etherscan.io/token/%s?a=%s\n", nftID, serviceInstance.EthClient.NftContractAddr, nftID)
}
