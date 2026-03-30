package main

type MyNumber int

type Number interface {
	~int | ~float64 // ~ Permite o Go considerar utilizar outras contraints que são dos tipos especificados
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[G Number](a G, b G) bool {
	if a > b {
		return true
	}

	return false
}

func main() {
	m := map[string]int{"Leonardo": 9000, "Lais": 4500, "Neville": 15000}
	m2 := map[string]float64{"Leonardo": 100.20, "Lais": 2000.3, "Neville": 300.0}
	m3 := map[string]MyNumber{"Leonardo": 9000, "Lais": 4500, "Neville": 15000}
	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10.0))
}
