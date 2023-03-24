package main

import "fmt"

const a = "Hello World"

type ID int

var z ID = 3

var (
	b bool    = true
	c int     = 5
	d string  = "jorge"
	e float32 = 1.32
	f float64 = 1.64
	g ID      = 2
)

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 5
	meuArray[2] = 10

	fmt.Println(meuArray[len(meuArray)-1])

	for i, v := range meuArray {
		fmt.Printf("O índice %d tem o valor de: %d\n", i, v)
	}
}
