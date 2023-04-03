package main

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	valor1 := 10
	valor2 := 20

	println(soma(&valor1, &valor2))
	println(valor1, valor2)
}
