package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {

	mapInt := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	mapFloat := map[string]float64{"Wesley": 100.20, "João": 200.30, "Maria": 300.40}
	mapMyNumber := map[string]MyNumber{"Wesley": 5000, "João": 6000, "Maria": 7000}

	println(Soma(mapInt))
	println(Soma(mapFloat))
	println(Soma(mapMyNumber))

	println(Compara(10, 10))
	println(Compara(10.3, 10.5))
}

// go run 3-foundation/17-generics/*.go
