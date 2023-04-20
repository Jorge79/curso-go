package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"-"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if res != nil {
		println(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	jsonPuro := []byte(`{"n": 2, "s": 200}`)
	var contaX Conta

	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	println(contaX.Saldo)
	println(contaX.Numero)
	fmt.Printf(string(jsonPuro))
}
