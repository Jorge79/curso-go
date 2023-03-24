package main

import (
	"fmt"
	"net/http"
)

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
	ID := 10
	println(a, b, c, d, e, f, ID, z)
	fmt.Printf("O tipo de ID é %T", g)
}
