package main

import "fmt"

// Structs aula F019
// Go não é uma linguagem orientada a objetos, a forma mais clara de escrever em Go é
// usando Structs. Uma Struct não é uma classe

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	renan := Cliente{
		Nome:  "Renan",
		Idade: 28,
		Ativo: true,
	}
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", renan.Nome, renan.Idade, renan.Ativo)
}
