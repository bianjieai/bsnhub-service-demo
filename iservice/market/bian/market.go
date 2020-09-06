package bian

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type BianMarket struct{}

type Result struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func (market BianMarket) GetPrice(base string, quote string) (price float64, error string) {
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s",
		strings.ToUpper(base+quote))
	resp, err := http.Get(url)
	if err != nil {
		return price, fmt.Sprintf("can not get price from api.binance.com, base: %s, quote: %s, err:%s",
			base, quote, err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return price, fmt.Sprintf("can not get price from api.binance.com, bad reqponse status: %d",
			resp.StatusCode)
	}

	var result Result
	err = json.Unmarshal(body, &result)
	if err != nil {
		return price, fmt.Sprintf("can not unmarshal price result: %s",
			err.Error())
	}

	price, err = strconv.ParseFloat(result.Price, 10)
	if err != nil {
		return price, fmt.Sprintf("can not parse price: %s",
			result.Price)
	}

	return price, ""
}
