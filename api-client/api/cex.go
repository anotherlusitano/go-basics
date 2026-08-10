package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"client.api/data"
)

const apiUrl string = "https://cex.io/api/ticker/%s/EUR"

func GetRate(currency string) (*data.Rate, error) {
	upCurrency := strings.ToUpper(currency)

	url := fmt.Sprintf(apiUrl, upCurrency)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status code received: %v", res.StatusCode)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response CEXResponse

	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}

	rate := data.Rate{Currency: currency, Price: response.Bid}

	// Get the JSON data
	// son := string(bodyBytes)
	// fmt.Println(son)

	return &rate, nil
}
