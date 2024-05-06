package internal

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

type ItemModel struct {
	SKU          string `json:"SKU"`
	UnitPrice    int    `json:"unitPrice"`
	SpecialPrice struct {
		Amount int `json:"amount"`
		Price  int `json:"price"`
	} `json:"specialPrice"`
}

func NewPricingModel(file []byte) (map[string]ItemModel, error) {
	pricingModel := make(map[string]ItemModel)
	err := json.Unmarshal(file, &pricingModel)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %w", err)
	}
	return pricingModel, nil
}
