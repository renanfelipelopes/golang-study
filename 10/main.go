package main

import (
	"fmt"
	"net/http"
)

// Closures aula F018
// closures sao funcoes anonimas, podendo ter um funcao dentro de outra funcao
func main() {
	total := func() int {
		return sum(1, 24, 45, 5, 65, 67, 221, 45, 12, 32, 433) * 2
	}()

	fmt.Println(total)

	// o mais comum é usar a funcao anonima para rodar um web server por exemplo
	// Definindo uma função anônima para tratar todas as requisições HTTP
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bem-vindo ao meu servidor web em Go!")
	})

	// Iniciando o servidor na porta 8080
	fmt.Println("Servidor web iniciado. Acesse http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
	// Neste exemplo, a função anônima é definida dentro de http.HandleFunc(),
	// que define como lidar com as requisições HTTP. Essa função simplesmente
	// escreve "Bem-vindo ao meu servidor web em Go!" no corpo da resposta HTTP.

	// Ao executar este programa, um servidor web será iniciado na porta 8080 do
	// localhost. Quando você acessar http://localhost:8080 em seu navegador, verá
	// a mensagem "Bem-vindo ao meu servidor web em Go!".
}

// a ideia agora é somar uma infinidade de numeros que nao sabemos a quantidade
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
