package main

type Conta struct {
	saldo int
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println("Saldo foi adicionado com sucesso, o total é de:", c.saldo)
	return c.saldo
}

func main() {
	conta := Conta{saldo: 100}

	conta.simular(200)
	println("Saldo final:", conta.saldo)
}
