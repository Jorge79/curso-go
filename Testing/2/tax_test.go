package tax

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	tax, err := calculateTax(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = calculateTax(0)
	assert.Error(t, err, "amount must be greater than zero")
	assert.Equal(t, 0.0, tax)
	assert.Contains(t, err.Error(), " greater than zero")
}
