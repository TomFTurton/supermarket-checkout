package internal

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed Items.json
var file []byte

type PricingModel struct {
	Items []struct {
		SKU          string `json:"SKU"`
		UnitPrice    int    `json:"unitPrice"`
		SpecialPrice struct {
			Amount int `json:"amount"`
			Price  int `json:"price"`
		} `json:"specialPrice"`
	} `json:"items"`
}

func NewPricingModel() (*PricingModel, error) {
	var model PricingModel
	err := json.Unmarshal(file, &model)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %w", err)
	}
	return &model, nil
}
