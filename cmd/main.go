package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"os"
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

	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var items []string
	for scanner.Scan() {
		// add line to item list
		items = append(items, scanner.Text())
	}

	total, err := supermarket.Checkout(items)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the total of your shop is: %d \n", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()
}
