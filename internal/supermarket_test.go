package internal

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheckout(t *testing.T) {
	// create test pricing model
	model, err := NewPricingModel(file)
	require.NoError(t, err)
	sm := NewSupermarket(model)

	tt := []struct {
		name string
		// test input
		input []string
		// expected output
		total int
		// expected errors
		expectedError error
	}{
		{
			name:  "happy path",
			input: []string{"A", "A", "A", "B"},
			total: 135,
		},
	}
	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			total, err := sm.Checkout(test.input)
			if test.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.total, total)
			}
		})
	}
}
