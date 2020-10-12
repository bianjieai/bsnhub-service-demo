package market

import (
	"math/rand"

	"github.com/bianjieai/bsnhub-service-demo/examples/market/market/binance"
	"github.com/bianjieai/bsnhub-service-demo/examples/market/market/huobi"
)

type Market interface {
	GetPrice(base string, quote string) (price float64, error string)
}

type RandomMarket struct {
}

func (market RandomMarket) GetPrice(base string, quote string) (price float64, error string) {
	return rand.Float64(), ""
}

var MarketType string

func GetMarket() Market {
	switch MarketType {
	case "huobi":
		return huobi.HuobiMarket{}
	case "binance":
		return binance.BinanceMarket{}
	default:
		return RandomMarket{}
	}
}
