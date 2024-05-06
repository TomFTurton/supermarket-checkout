package internal

import (
	_ "embed"
	"errors"
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
		name          string
		input         []string
		total         int
		expectedError error
	}{
		{
			name:  "happy path",
			input: []string{"A", "B", "A", "A", "B", "A"},
			total: 190,
		},
		{
			name:          "item not found in pricing model",
			input:         []string{"INVALID"},
			total:         0,
			expectedError: errors.New("error: item not found"),
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
