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
	m := map[string]int{"Hugo": 1000, "Mariana": 2000, "Luara": 5000}
	m2 := map[string]float64{"Hugo": 100.20, "Mariana": 200.3, "Luara": 300.0}
	m3 := map[string]MyNumber{"Hugo": 1000, "Mariana": 2000, "Luara": 3000}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10))

}
