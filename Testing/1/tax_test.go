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

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculateTax(500)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculateTax2(500)
	}
}
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := calculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f; expected 0", result)
		}

		if amount > 20000 && result != 20 {
			t.Errorf("Received %f; expected 20", result)
		}
	})
}
