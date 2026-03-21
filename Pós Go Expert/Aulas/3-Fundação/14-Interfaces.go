/*
	Aula: Interfaces

	Go possui interfaces, mas elas funcionam de forma DIFERENTE
	das interfaces em linguagens OO clássicas (Java, C#).

	Interface em Go:
	- Define um CONJUNTO DE COMPORTAMENTOS (métodos).
	- Não define estado (não possui atributos/campos).
	- Não existe palavra-chave "implements".
	- A implementação é IMPLÍCITA.

	Implementação implícita:
	- Qualquer struct que implemente TODOS os métodos da interface
	  automaticamente satisfaz (implementa) essa interface.
	- Isso reduz acoplamento e aumenta flexibilidade.

	Interfaces em Go são:
	- Pequenas
	- Focadas em comportamento
	- Muito usadas para abstração, testes e desacoplamento

	Frase clássica do Go:
	"Accept interfaces, return structs"
*/

package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

// Interface define comportamento, não dados
type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

// Implementação implícita da interface Pessoa
func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado.\n", c.Nome)
}

func main() {
	renan := Cliente{
		Nome:  "Renan",
		Idade: 28,
		Ativo: true,
	}

	renan.Desativar()
}

/*

	🔍 O conceito MAIS importante dessa aula
	🔹 Implementação implícita

	Em Go, você não escreve algo como:
	JAVA -> class Cliente implements Pessoa
	GO   -> func (c Cliente) Desativar() {}

	👉 Se compila, implementa.

	Isso torna o código:
	- Mais desacoplado
	- Mais fácil de evoluir
	- Melhor para testes

	⚠️ Atenção: Receiver por valor vs ponteiro
	Seu método está assim:
		func (c Cliente) Desativar() {
			c.Ativo = false
		}

	⚠️ Isso NÃO altera o estado real do cliente, porque c é uma cópia.
	Forma correta (muito importante):
		func (c *Cliente) Desativar() {
			c.Ativo = false
			fmt.Printf("O cliente %s foi desativado.\n", c.Nome)
		}
	Agora sim:
	- O estado original é alterado
	- É o padrão do mercado
*/

/*

	🛠 Interfaces no mundo real (Go de verdade)
	Você vai ver isso o tempo todo:
	🔹 io.Reader
		type Reader interface {
			Read(p []byte) (n int, err error)
		}

	- Arquivo
	- Buffer
	- HTTP Body
	- Todos implementam Reader

	🧠 Interface pequena é regra
	Em Go, interfaces costumam ter 1 ou 2 métodos.
	👉 Interfaces grandes são consideradas code smell.
*/
