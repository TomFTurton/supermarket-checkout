package internal

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testItems.json
var file []byte

//go:embed invalidItems.json
var invalidFile []byte

func TestNewPaymentModel(t *testing.T) {
	model, err := NewPricingModel(file)

	assert.NoError(t, err)
	assert.NotEmpty(t, model)
	assert.Equal(t, model["A"].SKU, "A")
	assert.Equal(t, model["A"].UnitPrice, 50)
	assert.Equal(t, model["A"].SpecialPrice.Price, 130)
	assert.Equal(t, model["A"].SpecialPrice.Amount, 3)

	assert.Equal(t, model["B"].SKU, "B")
	assert.Equal(t, model["B"].UnitPrice, 5)
	assert.Equal(t, model["B"].SpecialPrice.Price, 0)
	assert.Equal(t, model["B"].SpecialPrice.Amount, 0)
}

func TestInvalidModelFile(t *testing.T) {
	_, err := NewPricingModel(invalidFile)

	assert.Error(t, err)
}
