/*
	Aula: Composição de Structs

	Go NÃO possui herança.
	Em vez disso, utiliza COMPOSIÇÃO como principal forma de reutilização de código
	e modelagem de domínio.

	Composição de Structs:
	- Uma struct pode conter outra struct.
	- Isso cria uma relação "tem um" (has-a), não "é um" (is-a).
	- É mais explícito, previsível e simples que herança.

	Embedding (Struct Embedding):
	- Quando uma struct é declarada sem nome de campo (apenas o tipo),
	  dizemos que ela foi EMBUTIDA (embedded).
	- Os campos da struct embutida são PROMOVIDOS para a struct externa.

	Promoção de campos:
	- Permite acessar os campos internos diretamente:
	  cliente.Cidade em vez de cliente.Endereco.Cidade
	- Não é herança.
	- Não existe polimorfismo automático.

	Essa abordagem é muito usada em Go:
	- Para composição de domínio
	- Para reaproveitar comportamento
	- Para criar APIs simples e legíveis
*/

package main

import "fmt"

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
	Endereco // Struct embutida (composição)
}

func main() {
	renan := Cliente{
		Nome:  "Renan",
		Idade: 23,
		Ativo: true,
	}

	// Alterando campo da própria struct
	renan.Ativo = false

	// Alterando campo da struct embutida (campo promovido)
	renan.Cidade = "Sao Paulo"

	// Forma explícita (sem promoção)
	renan.Endereco.Cidade = "Sao Paulo"

	fmt.Printf(
		"Nome: %s, Idade: %d, Ativo: %t\n",
		renan.Nome,
		renan.Idade,
		renan.Ativo,
	)
}

/*
	🔍 O ponto-chave da aula
	🔹 Endereco sem nome de campo
		type Cliente struct {
			Endereco
		}

	Isso é chamado de:
	👉 Struct Embedding

	E causa:
	👉 Promoção de campos

	🧠 Por que posso acessar das duas formas?
		renan.Cidade
		renan.Endereco.Cidade

	Porque o Go:
	- Procura primeiro na struct Cliente
	- Se não achar, procura nas structs embutidas
	- Se existir sem conflito, ele promove o campo
*/
