package main

import "fmt"

func main() {
	//Memória -> Endereço -> Valor

	var a uint8 = 10
	var ponteiro *uint8 = &a
	*ponteiro = 20

	fmt.Println(a, &a, ponteiro, *ponteiro, &ponteiro)
}
