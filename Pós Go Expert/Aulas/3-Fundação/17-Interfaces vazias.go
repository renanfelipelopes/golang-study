/*
	Aula: Interfaces Vazias (Empty Interface)

	Interface vazia:
	- É uma interface sem métodos.
	- Representada por: interface{}

	Significado:
	- Qualquer tipo implementa interface{} automaticamente,
	  porque toda struct já "implementa" zero métodos.

	A partir do Go 1.18:
	- Foi criado o alias 'any'
	- any == interface{}

	Interface vazia NÃO é a mesma coisa que Generics.
	- Interface vazia aceita qualquer tipo, mas perde segurança de tipo.
	- Generics mantém segurança de tipo em tempo de compilação.

	Uso comum:
	- JSON dinâmico
	- Logs
	- APIs genéricas
	- Estruturas dinâmicas
*/

package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Hello, World!"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é '%T' e o valor é %v \n", t, t)
}

/*
🔍 O que é interface vazia na prática?
var x interface{}
👉 Significa:
"x pode receber QUALQUER tipo"

Exemplo:
x = 10
x = "Renan"
x = true
x = Cliente{}

🧠 Como isso funciona internamente (conceito avançado)
Uma interface vazia guarda:
┌──────────────┐
│ Tipo real    │ → int
│ Valor real   │ → 10
└──────────────┘

Ela guarda:
Tipo concreto
Valor concreto

Por isso o Go consegue fazer:
fmt.Printf("%T", x)


⚠️ Interface vazia ≠ Generics
❌ Interface vazia não é igual generics
✅ Ambos resolvem flexibilidade, mas diferente

🔹 Interface vazia
func print(v interface{}) {}

Problema:
- Perde type safety
- Precisa type assertion
- Bugs só aparecem em runtime

🔹 Generics
func print[T any](v T) {}

Vantagem:
- Type safety
- Melhor performance
- Erros em compile time

🧪 Problema clássico com interface vazia
func soma(a, b interface{}) interface{} {
	return a.(int) + b.(int)
}
	Se vier string → panic runtime.

🧪 Forma segura → type switch
func show(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("É int")
	case string:
		fmt.Println("É string")
	default:
		fmt.Println("Desconhecido")
	}
}

⚠️ Desvantagens reais
- Perde autocomplete
- Perde validação compile time
- Pode gerar panic
- Mais difícil de manter

👉 Por isso hoje:
- Prefere generics quando possível
- Interface vazia só quando necessário

🧠 Regra mental moderna Go
Use:
✔ Generics → quando sabe o formato
✔ Interface → quando quer comportamento
✔ Interface vazia → quando é realmente dinâmico

🧠 Resumo mental final
- interface{} aceita qualquer tipo
- Guarda tipo + valor internamente
- Não é generics
- Pode gerar bugs runtime
- any é alias moderno
- Use com cuidado
*/
