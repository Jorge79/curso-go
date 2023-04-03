package main

import (
	"fmt"
)

func main() {
	var minhaVar interface{} = "Jorge Francisco"
	fmt.Println(minhaVar.(string))

	res, ok := minhaVar.(string)
	fmt.Println("O valor de res é", res, "e o valor de ok é", ok)

	res2 := minhaVar.(int)
	fmt.Println("O valor de res2 é", res2)
}
