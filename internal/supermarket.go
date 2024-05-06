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
	// for item, count := range items {
	// 	itemSubTotal := 0

	// 	// s.model.Items :=
	// }
	return 0, errors.New("not implemented")
}
