package main

import (
	"fmt"
	"sync"

	"client.api/api"
)

func main() {
	currencies := []string{"BTC", "ETH", "BCH"}

	var wg sync.WaitGroup

	for _, currency := range currencies {
		wg.Go(func() {
			getCurrencyData(currency)
		})
	}

	wg.Wait()
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err != nil {
	}

	fmt.Printf("The rate for %v is %.2f\n", rate.Currency, rate.Price)
}
