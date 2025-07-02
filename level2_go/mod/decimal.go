package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}
	fmt.Println(price)
}
