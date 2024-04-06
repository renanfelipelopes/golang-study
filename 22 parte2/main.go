package main

/* --> Pacotes e módulos aula 31

- Modificadores de acesso
Para saber se uma função, variavel, struct, propriedade de uma struct, metodo,
está exportado para fora do pacotes, basta olharmos se o nome daquilo que queremos
está com a primeira letra Maiúscula.

Se deixarmos o nome da nossa funcao Soma da pasta matematica em minusculo,
esse arquivo main não irá encontrar a função. Ela ficará reservada para uso apenas
dentro do pacote matematica. Fora dele, ninguém mais consegue acessar.

*/

import (
	"fmt"
)

func main() {
	s := matematica.soma(10, 20)

	fmt.Println("Resultado: ", s)
}
