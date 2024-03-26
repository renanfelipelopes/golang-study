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
	Nome     string
	Idade    int
	Ativo    bool
	Endereco // esse o struct criado acima
}

func main() {
	renan := Cliente{
		Nome:  "Renan",
		Idade: 28,
		Ativo: true,
	}
	renan.Ativo = false // posso alterar o valor
	renan.Cidade = "Sao Paulo"
	renan.Endereco.Cidade = "Sao Paulo" // posso fazer das duas formas, isso é uma composicao

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", renan.Nome, renan.Idade, renan.Ativo)
}
