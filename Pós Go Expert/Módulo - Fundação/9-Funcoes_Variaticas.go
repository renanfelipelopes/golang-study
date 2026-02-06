/*
	Aula: Funções variáticas


	Para entender melhor o que é funções variáticas, podemos usar o seguinte exemplo:
	Se eu precisar somar uma infinidade de números mas eu não sei a quantidade de números
	que eu vou somar. O que fazer nesse caso?

	Utilizar os 3 pontos.
	🔹 O que são os três pontos (...) em parâmetros de função?
	Em Go, esses três pontos indicam que a função é variádica.

	📌 Nome técnico
	👉 Função variádica (variadic function)

	🧠 Conceito
	Uma função variádica é uma função que pode receber zero, um ou vários valores
	de um mesmo tipo, sem que você saiba antecipadamente quantos virão.

	func sum(numeros ...int) int {
	}

	Isso significa:
	- numeros pode receber quantos int forem passados
	- Dentro da função, numeros vira um slice ([]int)
	- Você pode iterar normalmente com for

	🔍 O que o Go faz por baixo dos panos?
		Esse código:
			sum(1, 2, 3, 4)


		É tratado internamente como:
			sum([]int{1, 2, 3, 4}...)

		Ou seja:
			numeros é um slice
			O ... empacota os argumentos em um slice

	🔁 Passando um slice existente
	Se você já tem um slice:
		nums := []int{1, 2, 3, 4}

	❌ Isso NÃO funciona:
		sum(nums)

	✅ Isso funciona:
		sum(nums...)
		👉 Aqui o ... desempacota o slice em argumentos individuais.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(12, 154, 1, 26, 3, 659, 78, 9, 984, 14, 56, 12))
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

/*
⚠️ Regras importantes (cai em prova e entrevista)
	Só pode existir UM parâmetro variádico
	Ele tem que ser o último parâmetro
Exemplo válido:
	func log(prefix string, values ...int) {}

Exemplo inválido:
	func log(values ...int, prefix string) {} // ❌ erro

*/
