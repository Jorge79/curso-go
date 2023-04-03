package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	jorge := Cliente{
		Nome:  "Jorge",
		Idade: 30,
		Ativo: true,
	}
	jorge.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", jorge.Nome, jorge.Idade, jorge.Ativo)
}
