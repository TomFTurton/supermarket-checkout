package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPaymentModel(t *testing.T) {
	model, err := NewPricingModel()

	assert.NoError(t, err)
	assert.NotEmpty(t, model)
	assert.Equal(t, model.Items[0].SKU, "A")
	assert.Equal(t, model.Items[0].UnitPrice, 50)
	assert.Equal(t, model.Items[0].SpecialPrice.Price, 130)
	assert.Equal(t, model.Items[0].SpecialPrice.Amount, 3)
}
