package main

import (
	"fmt"

	"younes.dev/go/crypto/api"
)

func main() {
	rate, err := api.GetRate("BTC")
	if err != nil {
		fmt.Printf("The rate for %v is %v \n", rate.Currency, rate.Price)
	}
}