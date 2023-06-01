package math

type math struct {
	a int
	b int
}

var X string = "Hello World"

func NewMath(a, b int) math {
	return math{a: a, b: b}
}

func (m *math) Add() int {
	return m.a + m.b
}
