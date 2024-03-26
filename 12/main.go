package main

import "fmt"

// Composicao de Structs aula F020

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

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
