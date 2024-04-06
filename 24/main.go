package main

/* --> For aula 33
No Go, laços de repetições como foreach, while, do while, não existem,
apenas o for e map
Segue exemplos:
*/

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "três"}
	for k, _ := range numeros {
		println(k)
	}

	// podemos usar o for parecido como um while
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// temos o loop infinito que tbm é comum em Go
	// porque é usado? imagine que vc queira consumir mensagens de uma fila,
	// vc pode criar um metodo operador que fique sempre escutando e executando o que vier da fila
	for {
		print("Hello, World!")
	}

}
