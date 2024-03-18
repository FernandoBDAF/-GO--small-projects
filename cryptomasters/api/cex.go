package api

import (
	"cryptomasters/datatypes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"encoding/json"
)

const apiURL = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf("3 characters expected, recivied %v", len(currency))
	}
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiURL, upCurrency))
	if err != nil {
		return nil, err
	}
	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("API request failed with status code %d", res.StatusCode)
	}
	rate := datatypes.Rate{Currency: upCurrency, Price: response.Bid}
	return &rate, nil
}