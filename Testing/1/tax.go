package tax

import "time"

func calculateTax(amount float64) float64 {
	if amount <= 0 {
		return 0
	}

	if amount >= 1000 && amount < 20000 {
		return 10.0
	}

	if amount >= 20000 {
		return 20.0
	}

	return 5.0
}

func calculateTax2(amount float64) float64 {
	time.Sleep(1 * time.Second)
	if amount == 0 {
		return 0
	}

	if amount >= 1000 {
		return 10
	} else {
		return 5
	}
}
