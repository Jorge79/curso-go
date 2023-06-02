package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := calculateTax(amount)

	if result != expected {
		t.Errorf("calculateTax(%f) = %f; expected %f", amount, result, expected)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{0, 0},
		{500, 5},
		{1000, 10},
		{2000, 10},
	}

	for _, item := range table {
		result := calculateTax(item.amount)
		if result != item.expected {
			t.Errorf("calculateTax(%f) = %f; expected %f", item.amount, result, item.expected)
		}
	}
}
