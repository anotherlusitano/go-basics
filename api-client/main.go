package main

import "client.api/api"

func main() {
	rate, err := api.GetRate("BTC")

	print(rate.Price, err)
}
