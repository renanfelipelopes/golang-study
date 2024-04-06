package main

/*
// Generics aula F028

Sem generics, teriamos que criar uma classe para a mesma operacao apenas porque
o tipo muda de int para float nesse exemplo, mas a funcao faz a mesma coisa que eh somar.

func SomaInt(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

Posso criar uma funcao com tipo generico que aceite receber tanto int como float:

func SomaGenerics[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

E posso também, criar uma funcao com o tipo constraints, que define um tipo generico
numa interface que é passada como parametro:

type Number interface {
	int | float64
}

func SomaGenerics[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

Posso criar um type para tentar forçar a funcao SomaGenerica aceitar um tipo inteiro na invocacao do metodo,
mas para funcionar, preciso adicionar um sinal de til ~ antes do tipo para o Go abrir
uma exceção e permitir que um int assuma o valor do int do generico:

type MyNumber int

type Number interface {
	~int | ~float64
}

func SomaGenerics[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	m := map[string]int{"Renan": 1000, "João": 2000, "Maria": 3000}
	println(SomaGenerics(m))
}
*/

type MyNumber int

type Number interface {
	~int | ~float64
}

func SomaGenerics[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	m := map[string]int{"Renan": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Renan": 1000.1, "João": 2000.1, "Maria": 3000.1}
	// println(SomaInt(m))
	// println(SomaFloat(m2))
	println(SomaGenerics(m))
	println(SomaGenerics(m2))
}
