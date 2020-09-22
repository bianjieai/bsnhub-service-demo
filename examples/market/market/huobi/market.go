package huobi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HuobiMarket struct{}

type Result struct {
	Status string `json:"status"`
	CH     string `json:"ch"`
	TS     int64  `json:"ts"`
	Tick   Tick   `json:"tick"`
}

type Tick struct {
	TS   int64  `json:"ts"`
	Data []Data `json:"data"`
}

type Data struct {
	Amount    float64 `json:"amount"`
	TradeID   int64   `json:"trade-id"`
	TS        int64   `json:"ts"`
	Price     float64 `json:"price"`
	Direction string  `json:"direction"`
}

func (market HuobiMarket) GetPrice(base string, quote string) (price float64, error string) {
	url := fmt.Sprintf("https://api.huobi.pro/market/trade?symbol=%s",
		base+quote)
	resp, err := http.Get(url)
	if err != nil {
		return price, fmt.Sprintf("can not get price from api.huobi.pro, base: %s, quote: %s, err:%s",
			base, quote, err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return price, fmt.Sprintf("can not get price from api.huobi.pro, bad reqponse status: %d",
			resp.StatusCode)
	}

	var result Result
	err = json.Unmarshal(body, &result)
	if err != nil {
		return price, fmt.Sprintf("can not unmarshal price result: %s",
			err.Error())
	}

	if result.Status != "ok" {
		return price, fmt.Sprintf("failed to get price from api.huobi.pro fa, response status must be ok: %s",
			result.Status)
	}

	if len(result.Tick.Data) == 0 {
		return price, ""
	}

	return result.Tick.Data[0].Price, ""
}
