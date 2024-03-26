package main

import "fmt"

// Interface aula F022

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

// Uma interface em Go é similiar a uma interface em POO
// No exemplo seguinte, implementamos uma interface, e significa
// que qualquer Struct que tiver o metodo Desativar, estará implementando
// uma interface. Uma interface não precisa ser implementada igual ao Java.
type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

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
