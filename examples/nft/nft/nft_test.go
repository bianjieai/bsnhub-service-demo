package nft

import (
	"fmt"
	"github.com/bianjieai/bsnhub-service-demo/common"
	"github.com/bianjieai/bsnhub-service-demo/examples/nft/nft/ethereum"
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestMintNft(t *testing.T) {
	viper, err := common.LoadYAMLConfig("../config/config.yaml")
	require.Nil(t, err)

	ethClient := ethereum.MakeEthChain(ethereum.NewConfig(viper))

	to := ethcmn.HexToAddress("0xA0214bC81667eFe7C8A703C30d145fFd89daBD58")
	amount := big.NewInt(1)
	metaId := "-Z-2fJxzCoFJ0MOU-zA3-tiIh7dK6FjDruAxgxW6PEs"
	setPrice := big.NewInt(10000000000000000)
	isForSale := true
	nftId, err := ethClient.MintNft(to, amount, metaId, setPrice, isForSale)

	require.Nil(t, err)
	fmt.Printf("Nft %s minted successfully: https://rinkeby.etherscan.io/token/%s?a=%s\n", nftId, ethClient.NftCoreAddr, nftId)
}
