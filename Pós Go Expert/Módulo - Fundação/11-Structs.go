/*
	Aula: Structs

	Go não é uma linguagem orientada a objetos clássica (como Java ou C#).
	Em Go, não existem classes, herança ou construtores da forma tradicional.

	O principal recurso para modelar dados e comportamentos é a STRUCT.

	Struct:
	- É um tipo composto que agrupa campos relacionados.
	- Representa "dados com significado".
	- Pode ter métodos associados.
	- NÃO é uma classe.

	Composição:
	- Go não suporta herança.
	- Em vez disso, usa COMPOSIÇÃO (struct dentro de struct).
	- A composição é preferida porque gera código mais simples, explícito e previsível.

	Encapsulamento:
	- Controlado pelo uso de letras maiúsculas e minúsculas.
	- Campos com letra MAIÚSCULA são exportados (públicos).
	- Campos com letra minúscula são não exportados (privados ao pacote).

	Go incentiva:
	- Simplicidade
	- Clareza
	- Menos abstrações artificiais
*/

package main

import "fmt"

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

	fmt.Printf(
		"Nome: %s, Idade: %d, Ativo: %t\n",
		renan.Nome,
		renan.Idade,
		renan.Ativo,
	)
}

/*

	🔍 O que está acontecendo aqui?
		type Cliente struct {}

			Cria um novo tipo
			Agrupa dados relacionados
			Muito usado para:
				Entidades de domínio
				DTOs
				Models de banco
				Payloads de API


	🔧 Métodos em Structs
	Structs podem ter métodos:
		func (c Cliente) Ativar() Cliente {
			c.Ativo = true
			return c
		}

		renan = renan.Ativar()

		OBS: Isso NÃO transforma a struct numa classe.
			 O método apenas recebe a struct como receiver.

	🔁 Pointer Receiver (muito importante)
	Se quiser alterar o valor original:
		func (c *Cliente) Desativar() {
			c.Ativo = false
		}

		renan.Desativar()
*/

/*
	🧩 Composição (substituto da herança)
		type Endereco struct {
			Cidade string
			Estado string
		}

		type Cliente struct {
			Nome     string
			Idade    int
			Ativo    bool
			Endereco Endereco
		}
	Uso:
		cliente := Cliente{
			Nome: "Renan",
			Endereco: Endereco{
				Cidade: "Santo André",
				Estado: "SP",
			},
		}

		fmt.Println(cliente.Endereco.Cidade)

	👉 Isso é composição explícita.


	🔥 Composição com promoção de campos
		type Endereco struct {
			Cidade string
			Estado string
		}

		type Cliente struct {
			Nome string
			Endereco
		}

	Agora você pode fazer:
		fmt.Println(cliente.Cidade)


	Isso lembra herança, mas não é.
	É composição com promoção de campos.
*/
