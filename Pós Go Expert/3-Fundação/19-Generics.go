/*
	Aula: Generics (Go 1.18+)

	Generics permitem escrever código reutilizável com segurança de tipo
	em tempo de compilação.

	Antes dos Generics:
	- Precisávamos duplicar funções para cada tipo.
	- Ou usar interface{} e perder segurança de tipo.

	Com Generics:
	- Código reutilizável
	- Type safety
	- Melhor performance que interface{}
*/

/*

Sem generics, teriamos que criar uma classe para a mesma operacao apenas porque
o tipo muda de int para float nesse exemplo, mas a funcao faz a mesma coisa que eh somar.

func SomaInt(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

Posso criar uma funcao com tipo generico que aceite receber tanto int como float:

func SomaGenerics[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

E posso também, criar uma funcao com o tipo constraints, que define um tipo generico
numa interface que é passada como parametro:

type Number interface {
	int | float64
}

func SomaGenerics[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

Posso criar um type para tentar forçar a funcao SomaGenerica aceitar um tipo inteiro na invocacao do metodo,
mas para funcionar, preciso adicionar um sinal de til ~ antes do tipo para o Go abrir
uma exceção e permitir que um int assuma o valor do int do generico:

type MyNumber int

type Number interface {
	~int | ~float64
}

func SomaGenerics[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	m := map[string]int{"Renan": 1000, "João": 2000, "Maria": 3000}
	println(SomaGenerics(m))
}
*/

package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func SomaGenerics[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	m := map[string]int{"Renan": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Renan": 1000.1, "João": 2000.1, "Maria": 3000.1}
	// println(SomaInt(m))
	// println(SomaFloat(m2))
	println(SomaGenerics(m))
	println(SomaGenerics(m2))
}

// ========================================================================================================== //

/*

🧠 Problema antes dos Generics

func SomaInt(m map[string]int) int
func SomaFloat(m map[string]float64) float64

👉 Mesmo algoritmo
👉 Só muda tipo

Isso gera:
- Código duplicado
- Difícil manutenção
- Mais chance de bug

-----------------------------------------------------------------

🔥 Solução: Generics

func SomaGenerics[T int | float64](m map[string]T) T

🔍 O que significa [T int | float64]
👉 T é um Type Parameter
👉 Pode ser:
int
float64

🧠 Como o compilador trata isso
Ele gera versões especializadas:
SomaGenerics[int]
SomaGenerics[float64]

👉 Sem reflection
👉 Sem interface boxing
👉 Performance próxima de código manual

-----------------------------------------------------------------

🔥 Melhor prática: Constraints com Interface

Exemplo:

type Number interface {
	int | float64
}

Uso:
func SomaGenerics[T Number](m map[string]T) T

-----------------------------------------------------------------

🚀 Parte AVANÇADA — o ~ (Underlying Type Constraint)
Essa parte é MUITO importante.

❌ Sem ~
type Number interface {
	int | float64
}

Isso aceita:
✔ int
✔ float64

Mas NÃO aceita:
type MyNumber int

Porque:
👉 MyNumber ≠ int
👉 Só tem int como underlying type

✅ Com ~
type Number interface {
	~int | ~float64
}

Agora aceita:
✔ int
✔ float64
✔ type MeuInt int
✔ type MeuFloat float64

🧠 O que ~ significa?
👉 "Qualquer tipo cujo tipo base seja esse"

🧠 Visual mental

Sem ~

Aceita:
 int
 float64


Com ~

Aceita:
 int
 MyInt
 CustomInt
 float64
 MyFloat

 -----------------------------------------------------------------

⚠️ Quando NÃO usar Generics

Quando quer comportamento → Interface normal

Exemplo:

type Reader interface {
	Read([]byte) (int, error)
}

🧠 Regra mental definitiva Go moderno

👉 Generics → para dados
👉 Interface → para comportamento

-----------------------------------------------------------------

🧠 Resumo final

- Generics evitam duplicação
- Constraints definem limites do tipo
- ~ permite tipos derivados
- Melhor que interface{} na maioria dos casos
- Padrão moderno Go
*/
