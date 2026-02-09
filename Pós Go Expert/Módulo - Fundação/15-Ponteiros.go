/*
	Aula: Ponteiros

	Ponteiro é uma variável que ARMAZENA UM ENDEREÇO DE MEMÓRIA,
	e não um valor direto.

	Em Go, toda variável:
	- Possui um valor
	- Está armazenada em algum endereço da memória

	Quando fazemos:
		a := 10

	O Go:
	1. Reserva um espaço na memória
	2. Guarda o valor 10
	3. Associa esse espaço ao nome 'a'

	O operador '&' (e comercial):
	- Retorna o ENDEREÇO DE MEMÓRIA da variável

	O operador '*' (asterisco):
	- Usado para:
	  a) Declarar um ponteiro
	  b) Acessar (desreferenciar) o valor apontado
*/

package main

func main() {

	// Em algum lugar da memória, existe um endereço que contém um valor.
	// Quando fazemos a := 10, o Go cria uma "caixinha" na memória,
	// guarda o valor 10 dentro dela e associa um endereço a essa caixinha.
	a := 10

	// &a retorna o endereço de memória onde o valor 10 está guardado
	println(&a)

	// Criando um ponteiro que aponta para o endereço de 'a'
	var ponteiro *int = &a

	// O ponteiro guarda o ENDEREÇO de 'a'
	println(ponteiro)

	// Para acessar o VALOR que está no endereço, usamos '*'
	println(*ponteiro)

	// Alterando o valor diretamente na memória
	*ponteiro = 20

	// Agora 'a' também mudou
	println(a)
}

----------------------------------------------------------------------------------------

🧠 Visualizando a memória (isso fixa de vez)
Imagine a memória assim:
┌───────────────┐
│ Endereço      │ Valor
├───────────────┤
│ 0xc000055f38  │ 10   ← variável 'a'
└───────────────┘

Quando você faz:
a := 10

Depois:
ponteiro := &a

Fica assim:
ponteiro
   │
   ▼
┌───────────────┐
│ 0xc000055f38  │ 10
└───────────────┘

E quando faz:
*ponteiro = 20

┌───────────────┐
│ 0xc000055f38  │ 20
└───────────────┘

👉 Você mudou o valor direto na memória, não uma cópia.


| Operador | Significado                  |
| -------- | ---------------------------- |
| `&a`     | endereço de `a`              |
| `*int`   | tipo ponteiro para int       |
| `*p`     | valor armazenado no endereço |

----------------------------------------------------------------------------------------

📌 Stack vs Heap (ponto que confunde muita gente)

🧱 Stack
Memória automática
- Mais rápida
- Variáveis locais simples
- Vida curta (escopo da função)

🧠 Heap
Memória dinâmica
Usada quando:
- Retorna ponteiro
- Compartilha dados
- Escapa do escopo da função

⚠️ IMPORTANTE:
👉 Em Go, você não escolhe stack ou heap manualmente
👉 O escape analysis do compilador decide

func criaNumero() *int {
	n := 10
	return &n // n "escapa", vai para heap
}


🧪 Exemplo prático: sem ponteiro (cópia)
func alteraValor(x int) {
	x = 20
}

func main() {
	a := 10
	alteraValor(a)
	println(a) // 10
}

👉 x é uma cópia

----------------------------------------------------------------------------------------

🧪 Exemplo prático: com ponteiro (memória)
func alteraValor(x *int) {
	*x = 20
}

func main() {
	a := 10
	alteraValor(&a)
	println(a) // 20
}

👉 Aqui você alterou direto na memória

----------------------------------------------------------------------------------------

🔥 Ponteiros + Structs (uso real)
type Cliente struct {
	Nome  string
	Ativo bool
}

func desativar(c *Cliente) {
	c.Ativo = false
}

func main() {
	cliente := Cliente{Nome: "Renan", Ativo: true}
	desativar(&cliente)
	println(cliente.Ativo) // false
}

👉 Isso conecta ponteiros + methods + interfaces (aula passada).

----------------------------------------------------------------------------------------

🧠 Regra mental definitiva
	Valor → cópia
	Ponteiro → memória compartilhada

Se:
- Precisa mudar estado → ponteiro
- Precisa evitar cópia → ponteiro
- Struct grande → ponteiro

----------------------------------------------------------------------------------------

🧠 Resumo mental final
- Ponteiro guarda endereço, não valor
- & pega o endereço
- * acessa o valor
- Modificar via ponteiro altera a memória original
- Stack vs Heap é decisão do compilador

Ponteiros são essenciais em:
- Métodos
- Structs
- Interfaces
- Performance