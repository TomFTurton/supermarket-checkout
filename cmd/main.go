package main

import (
	_ "embed"
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

	err = supermarket.Checkout()
	if err != nil {
		log.Fatal(err)
	}
}
