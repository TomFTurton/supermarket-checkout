package internal

import "errors"

type Supermarket struct {
	model map[string]ItemModel
}

func NewSupermarket(model map[string]ItemModel) *Supermarket {
	return &Supermarket{model: model}
}

func (s *Supermarket) Checkout(input []string) (int, error) {
	// create count of the items
	items := make(map[string]int)
	for _, item := range input {
		items[item] = items[item] + 1
	}

	// calculate total
	shopTotal := 0
	for item, count := range items {
		pricing, found := s.model[item]
		if !found {
			return 0, errors.New("error: item not found")
		}
		itemSubTotal := 0
		if pricing.SpecialPrice.Amount > 0 {
			itemSubTotal += (count / pricing.SpecialPrice.Amount) * pricing.SpecialPrice.Price
			itemSubTotal += (count % pricing.SpecialPrice.Amount) * pricing.UnitPrice
		} else {
			itemSubTotal = count * pricing.UnitPrice
		}
		shopTotal += itemSubTotal
	}
	return shopTotal, nil
}
