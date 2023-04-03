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

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome  string
	Ativo bool
}

func (e Empresa) Desativar() {
	e.Ativo = false
	fmt.Printf("A empresa %s foi desativada\n", e.Nome)
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O status do cliente %s é: %t\n", c.Nome, c.Ativo)
}

func Desativacao(p Pessoa) {
	p.Desativar()
}

func main() {
	jorge := Cliente{
		Nome:  "Jorge",
		Idade: 30,
		Ativo: true,
	}
	minhaEmpresa := Empresa{
		Nome: "Google",
	}

	Desativacao(minhaEmpresa)

	Desativacao(jorge)

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", jorge.Nome, jorge.Idade, jorge.Ativo)
	fmt.Printf("Nome: %s, Ativo: %t", minhaEmpresa.Nome, minhaEmpresa.Ativo)
}
