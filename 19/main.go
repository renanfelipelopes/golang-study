package main

import "fmt"

// Interfaces vazias aula F026
func main() {
	var x interface{} = 10
	var y interface{} = "Hello, World!"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é '%T' e o valor é %v \n", t, t)
}

// O conceito de interface vazia eh similar ao conceito de Generics
// A interface vazia suporta qualquer coisa, tipo, ela implementa todo mundo.
// Uma vantagem é que vc consegue colocar qualquer tipo e trata esse tipos
// E uma desvantagem, é ter vulnerabilidades e facilidades a criar bugs.
