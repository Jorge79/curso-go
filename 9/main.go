package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum50(50, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)

	fmt.Println(sum50(1, 1))
}

func sum50(a, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
