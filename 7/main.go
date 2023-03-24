package main

import "fmt"

func main() {
	salarios := map[string]int{"Jorge": 10, "Maria": 50, "José": 100}
	fmt.Println(salarios["Jorge"])
	delete(salarios, "Jorge")
	salarios["João"] = 500

	for nome, salario := range salarios {
		fmt.Printf("O índice %s tem o valor de: %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário é: %d\n", salario)
	}
}
