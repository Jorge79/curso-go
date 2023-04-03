package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     string
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	jorge := Cliente{
		Nome:  "Jorge",
		Idade: 30,
		Ativo: true,
	}
	jorge.Ativo = false

	jorge.Cidade = "São Paulo"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", jorge.Nome, jorge.Idade, jorge.Ativo)
}
