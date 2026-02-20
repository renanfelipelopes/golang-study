/*
	Aula 3: Defer

	O Defer atrasa alguma coisa...

	Pode ser que no final de um processo, alguém esqueça de fechar o recurso,
	como por exemplo o "req.Body.Close()", ou até mesmo, nem queremos correr
	o risco de esquecer.

	Para isso, pode escrever "defer req.Body.Close()"

	O Defer é um statement que faz atrasar a execução da função que vem em seguida do defer.
	Se o colocarmos antes de abrir um arquivo, ou antes de fazer um http response,
	ele vai executar o close após ter concluído a execução de todas as linhas.
	Se você quer que algo execute por último, pode usar o defer antes.

*/

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://www.google.com")
	defer req.Body.Close()
	if err != nil {
		panic(err)
	}
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))
}

// Outro exemplo:
func main2() {
	defer fmt.Print("Primeira linha")
	fmt.Print("Segunda linha")
	fmt.Print("Terceira linha")
}
