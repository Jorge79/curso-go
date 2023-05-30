package main

import (
	"fmt"

	"github.com/Jorge79/estudos-go/packaging/1/math"
)

func main() {
	fmt.Println("Hello, World!")

	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
}
