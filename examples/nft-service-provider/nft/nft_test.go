package nft

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	ethcmn "github.com/ethereum/go-ethereum/common"

	"github.com/bianjieai/bsnhub-service-demo/examples/nft-service-provider/common"
)

func TestMintNft(t *testing.T) {
	config, err := common.LoadYAMLConfig("../config/config.yaml")
	require.Nil(t, err)

	serviceInstance := MakeNFTService(config)
	serviceInstance.Logger = common.Logger

	i := new(big.Int)
	i, ok := i.SetString("0", 10)
	require.True(t, ok)

	to := ethcmn.HexToAddress("0xA0214bC81667eFe7C8A703C30d145fFd89daBD58")
	amount := big.NewInt(1)
	metaID := "-Z-2fJxzCoFJ0MOU-zA3-tiIh7dK6FjDruAxgxW6PEs"
	setPrice := i
	isForSale := true

	nftID, err := serviceInstance.MintNft(to, amount, metaID, setPrice, isForSale)
	require.Nil(t, err)

	fmt.Printf("NFT %s minted successfully: https://rinkeby.etherscan.io/token/%s?a=%s\n", nftID, serviceInstance.EthClient.NftContractAddr, nftID)
}
