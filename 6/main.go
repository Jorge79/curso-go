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
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(s[:3]), cap(s[4:]), s[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(s[4:]), cap(s[4:]), s[4:])

	s = append(s, 110)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])
}
