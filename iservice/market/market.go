package market

import (
	"math/rand"

	"github.com/bianjieai/bsnhub-service-demo/iservice/market/bian"
	"github.com/bianjieai/bsnhub-service-demo/iservice/market/huobi"
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
		return bian.BianMarket{}
	default:
		return RandomMarket{}
	}
}
