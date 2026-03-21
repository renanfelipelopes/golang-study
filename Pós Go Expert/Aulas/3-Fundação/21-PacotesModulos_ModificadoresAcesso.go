package main

/* --> Pacotes e módulos - Modificadores de acesso

- Modificadores de acesso
Para saber se uma função, variavel, struct, propriedade de uma struct, metodo,
está exportado para fora do pacotes, basta olharmos se o nome daquilo que queremos
está com a primeira letra Maiúscula.

Se deixarmos o nome da nossa funcao Soma da pasta matematica em minusculo,
esse arquivo main não irá encontrar a função. Ela ficará reservada para uso apenas
dentro do pacote matematica. Fora dele, ninguém mais consegue acessar.

*/

import (
	"curso-go/matematica"
	"fmt"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Honda"}

	fmt.Println(carro.Andar())
	fmt.Println("Resultado:", s)
}
