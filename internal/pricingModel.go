package internal

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

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

func NewPricingModel(file []byte) (*PricingModel, error) {
	var model PricingModel
	err := json.Unmarshal(file, &model)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %w", err)
	}
	return &model, nil
}
