package matematica

func Soma[T int | float32](a, b T) T {
	return a + b
}
