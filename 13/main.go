package main

import "fmt"

// Composicao de Structs aula F021

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
	Endereco
}

// Go não possui POO, mas possui structs.
// Uma classe em POO tem métodos, em Go as Structs também tem métodos,
// Para fazer isso, basta utilizarmos parenteses antes do nome da funcao e,
// "atachar" a Struct dentro desse parenteses, onde no exemplo c representa a
// Struct Cliente, e por meio de c, conseguimos acessar as propriedades da Struct Cliente
func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado.", c.Nome)
}

func main() {
	renan := Cliente{
		Nome:  "Renan",
		Idade: 28,
		Ativo: true,
	}
	renan.Desativar()
}
