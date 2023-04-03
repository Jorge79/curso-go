package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{} = 10
	var y interface{} = "Teste"

	showType(x)
	showType(y)
}

func showType(x interface{}) {
	fmt.Println("O tipo da variável é", reflect.TypeOf(x), "e o valor é", x)
}
