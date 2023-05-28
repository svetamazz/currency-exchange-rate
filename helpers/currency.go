package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	. "github.com/svetamazz/currency-exchange-rate/types"
) 
	

func GetExchangeRate() (*ExchangeRate, error) {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=uah"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]map[string]float64
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	rate := result["bitcoin"]["uah"]
	exchangeRate := &ExchangeRate{
		BTCtoUAH: rate,
	}

	return exchangeRate, nil
}