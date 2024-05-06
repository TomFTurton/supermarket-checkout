package internal

import "errors"

type Supermarket struct {
	model *PricingModel
}

func NewSupermarket(model *PricingModel) *Supermarket {
	return &Supermarket{model: model}
}

func (s *Supermarket) Checkout([]string) (int, error) {
	return 0, errors.New("not implemented")
}
