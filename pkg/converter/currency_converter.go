package converter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	DKK = "DKK"
	SEK = "SEK"
	EUR = "EUR"
	USD = "USD"
)

var (
	allowedCurrencies = map[string]bool{
		DKK: true,
		SEK: true,
		EUR: true,
		USD: false, // USD is not allowed
	}
	ErrInvalidCurrency = fmt.Errorf("invalid currency")
)

type Currency struct {
	Currency string
}

const API_BASE_URL = "https://api.currencyapi.com/v3/latest?apikey="
const API_KEY = "cur_live_JfqyXEJ9hyM9dWZBOG2Zb1tzVfyEwjvhOkahRuEO"

type CurrencyResponse struct {
	Meta struct {
		LastUpdatedAt string `json:"last_updated_at"`
	} `json:"meta"`
	Data map[string]struct {
		Code  string  `json:"code"`
		Value float64 `json:"value"`
	} `json:"data"`
}

var client = &http.Client{Timeout: 10 * time.Second}

func jsonToStruct(url string, targetStruct any) error {
	response, err := client.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"HTTP request failed with status code %d",
			response.StatusCode,
		)
	}

	if err := json.NewDecoder(response.Body).Decode(targetStruct); err != nil {
		return err
	}

	return nil
}

func getCurrencyResponse(url string) (*CurrencyResponse, error) {
	var currencyResponse CurrencyResponse
	if err := jsonToStruct(url, &currencyResponse); err != nil {
		return nil, fmt.Errorf("failed to get currency response: %v", err)
	}
	return &currencyResponse, nil
}

func (c *Currency) Convert(amount float64, targetCurrency *Currency) (float64, error) {
	queryParams := "&base_currency=" + c.Currency +
		"&currencies=" + targetCurrency.Currency
	requestUrl := API_BASE_URL + API_KEY + queryParams

	currencyResponse, err := getCurrencyResponse(requestUrl)

	if err != nil {
		return 0, fmt.Errorf("failed to convert currency: %v", err)
	}

	result := amount * currencyResponse.Data[targetCurrency.Currency].Value
	return result, nil
}

func NewCurrency(currency string) (*Currency, error) {
	if !allowedCurrencies[currency] {
		return nil, ErrInvalidCurrency
	}

	return &Currency{
		Currency: currency,
	}, nil
}
