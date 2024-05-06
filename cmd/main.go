package main

import (
	"log"
	"supermarket-checkout/internal"
)

func main() {
	model := internal.NewPricingModel()
	supermarket := internal.NewSupermarket(model)

	err := supermarket.Checkout()
	if err != nil {
		log.Fatal(err)
	}
}
