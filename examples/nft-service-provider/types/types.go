package types

import (
	"math/big"

	ethcmn "github.com/ethereum/go-ethereum/common"
)

const (
	ServiceName = "nft"
)

type Input struct {
	ABIEncoded   string         `json:"abi_encoded"`
	To           ethcmn.Address `json:"to"`
	AmountToMint *big.Int       `json:"amount_to_mint"`
	MetaID       string         `json:"meta_id"`
	SetPrice     *big.Int       `json:"set_price"`
	IsForSale    bool           `json:"is_for_sale"`
}

type RawInput struct {
	To           ethcmn.Address `json:"to"`
	AmountToMint *big.Int       `json:"amount_to_mint"`
	MetaID       string         `json:"meta_id"`
	SetPrice     *big.Int       `json:"set_price"`
	IsForSale    bool           `json:"is_for_sale"`
}

type Output struct {
	NftID string `json:"nft_id"`
}

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
