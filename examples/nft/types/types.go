package types

import (
	ethcmn "github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Input struct {
	To           ethcmn.Address `json:"to"`
	AmountToMint *big.Int       `json:"amount_to_mint"`
	MetaId       string         `json:"meta_id"`
	SetPrice     *big.Int       `json:"set_price"`
	IsForSale    bool           `json:"is_for_sale"`
}

type Output struct {
	NftId string `json:"nft_id"`
}

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
