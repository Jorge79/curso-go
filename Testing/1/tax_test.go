package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 6.0

	result := calculateTax(amount)

	if result != expected {
		t.Errorf("calculateTax(%f) = %f; expected %f", amount, result, expected)
	}
}
