package main

import (
	_ "embed"
	"fmt"
	"log"
	"supermarket-checkout/internal"
)

//go:embed Items.json
var file []byte

func main() {
	model, err := internal.NewPricingModel(file)
	if err != nil {
		log.Fatal(err)
	}
	supermarket := internal.NewSupermarket(model)

	total, err := supermarket.Checkout([]string{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the total of your shop is: %v", total)
}
