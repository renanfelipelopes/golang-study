package main

import (
	"errors"
	"fmt"
)

// Aula: Funcoes
func main() {
	var soma = sum(2, 2)
	fmt.Println(soma)

	var valorSoma, valorBool = sum2(2, 2)
	fmt.Println(valorSoma, valorBool)

	// se eu nao quiser o valor de valorSoma, posso usar o 'Blank identifier'
	var _, valBool = sum2(1, 1)
	fmt.Println(valBool)
}

// no Go, os parametros vem antes do seu tipo, exemplo: func pessoa(nome string, salario int)
// e o tipo do retorno da função vem após os parenteses. Se a função for retornar um numero inteiro
// entao ficaria assim, exemplo: func pessoa(nome string, salario int) int {logica da funcao + return}
// se a função não for retornar nada, então não deve conter o tipo após os parenteses,
// exemplo: func pessoa(nome string, salario int) {logica da funcao sem o return}
// outra curiosidade, é que se os parametros forem do mesmo tipo, podemos declarar do seguindo
// modo: func sum(a, b int) int {}

func sum(a int, b int) int {
	return a + b
}

// as funções em Go podem retornar mais do que 1 valor, por exemplo, podemos dizer que nossa
// função deve retornar um valor inteiro e um boolean, ficaria assim:
func sum2(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}
	return a + b, false
}

// no Go, o uso de mais de um retorno é muito comum, principalmente em relacao a erros
// porque o Go não tem exception, como try catch, entao normalmente, se faz algo do tipo:
func sum3(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50.")
	}
	return a + b, nil
}

// o nil basicamente é um valor vazio, igual a null. Portanto, a funcao main() que chama sum3,
// deve invocar a funcao do seguinte modo:
func main2() {
	valor, err := sum3(50, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

// em go, geralmente o ultimo item de retorno será um error para verificar se deu tudo certo e tratar
