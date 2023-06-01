package main

import (
	"fmt"

	"github.com/Jorge79/estudos-go/packaging/1/math"
)

func main() {
	m := math.NewMath(1, 2)
	// m2 := math.Math{}
	// fmt.Println(m2.Add())
	fmt.Println(m.Add())
	println(math.X)
}
