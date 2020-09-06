package market

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetMarket(t *testing.T) {
	base := "link"
	quote := "usdt"

	// test huobi
	MarketType = "huobi"
	market := GetMarket()
	price, err := market.GetPrice(base, quote)
	fmt.Printf("huobi %s-%s price: %v\n", base, quote, price)
	require.NotEqual(t, float64(0), price)
	require.Equal(t, "", err)

	// test bian
	MarketType = "bian"
	market = GetMarket()
	price, err = market.GetPrice(base, quote)
	fmt.Printf("bian %s-%s price: %v\n", base, quote, price)
	require.NotEqual(t, float64(0), price)
	require.Equal(t, "", err)

	// test empty
	MarketType = ""
	market = GetMarket()
	price, err = market.GetPrice(base, quote)
	require.Equal(t, "", err)
}
