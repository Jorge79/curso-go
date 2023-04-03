package main

type MyNumber int

type Number interface {
	~int | ~float32
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](x, y T) bool {
	return x == y
}

func main() {
	m := map[string]int{"Jorge": 10, "Francisco": 20}
	m2 := map[string]float32{"Jorge": 10.5, "Francisco": 20.5}
	println(Soma(m))
	println(Soma(m2))

	m3 := map[string]MyNumber{"Jorge": 20, "Francisco": 30}
	println(Soma(m3))
	println(Compara(19, 20))
}
