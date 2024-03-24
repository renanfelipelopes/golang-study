package main

import (
	"fmt"
)

// Funcoes variaticas aula F017
func main() {
	var soma = sum(1, 24, 45, 5, 65, 67, 221, 45, 12, 32, 433)
	fmt.Println(soma)
}

// a ideia agora é somar uma infinidade de numeros que nao sabemos a quantidade
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

// um exemplo de onde essa funcao é utilizada, é quando vc precisa passar parametros de configuracao
// para um banco de dados, mas esses parametros podem ser apenas um como varios parametros.
// nesse caso, podemos usar os 3 pontinhos e o tipo da variavel e percorre toda
// a variavel para pegar os valores das configuracoes
