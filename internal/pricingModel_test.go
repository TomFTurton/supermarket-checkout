package internal

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testItems.json
var file []byte

func TestNewPaymentModel(t *testing.T) {
	model, err := NewPricingModel(file)

	assert.NoError(t, err)
	assert.NotEmpty(t, model)
	assert.Equal(t, model.Items[0].SKU, "A")
	assert.Equal(t, model.Items[0].UnitPrice, 50)
	assert.Equal(t, model.Items[0].SpecialPrice.Price, 130)
	assert.Equal(t, model.Items[0].SpecialPrice.Amount, 3)

	assert.Equal(t, model.Items[1].SKU, "B")
	assert.Equal(t, model.Items[1].UnitPrice, 5)
	assert.Equal(t, model.Items[1].SpecialPrice.Price, 0)
	assert.Equal(t, model.Items[1].SpecialPrice.Amount, 0)
}
