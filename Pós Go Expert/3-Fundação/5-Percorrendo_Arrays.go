package main

import "fmt"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Renan"
	e float64 = 1.2
	f ID      = 1
)

//trabalhando com array
func main() {
	// exemplo 1
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	for i, v := range meuArray { // o range é para percorrer todos os itens do array
		fmt.Printf("Indice: %d; valor: %d \n", i, v)
	}
	fmt.Println()
	fmt.Println()

	// exemplo 2
	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}

	for _, shark := range sharks {
		fmt.Println(shark)
	}
	fmt.Println()
	fmt.Println()

	// exemplo 3
	integers := make([]int, 10)
	fmt.Println(integers)

	for i := range integers {
		fmt.Println("-> ", i, " ")
		integers[i] = i
	}

	fmt.Println(integers)
	fmt.Println()
	fmt.Println()

	// exemplo 4
	sammy := "Sammy"

	for _, letter := range sammy {
		fmt.Printf("%c\n", letter)
	}
}

// func main() {
// 	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}

// 	for i, shark := range sharks {
// 		fmt.Println(i, shark)
// 	}
// }

// Neste caso, estamos imprimindo cada item da lista. Embora tenhamos usado as variáveis
// i e shark, poderíamos ter chamado as variáveis por qualquer outro nome de variável válido e
// ainda assim, obteríamos o mesmo resultado:
// Output:
// 0 hammerhead
// 1 great white
// 2 dogfish
// 3 frilled
// 4 bullhead
// 5 requiem

// Ao usar range em uma fatia, ele irá sempre retornar dois valores. O primeiro valor será
// o índice em que a iteração atual do loop está e o segundo será o valor naquele índice.
// Neste caso, para a primeira iteração, o índice era 0, e o valor era hammerhead.

// Às vezes, queremos apenas o valor dentro dos elementos da fatia, não do índice.
// Se alterarmos o código anterior para imprimir apenas o valor, no entanto, vamos
// receber um erro de tempo de compilação:

// func main() {
// 	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}

// 	for i, shark := range sharks {
// 		fmt.Println(shark)
// 	}
// }
// Output:
// src/range-error.go:8:6: i declared and not used

// Como o i foi declarado no loop for, mas nunca foi usado, o compilador responderá com
// o erro i declared and not used. Este é o mesmo erro que você receberá no Go sempre que
// for declarar uma variável e não a utilizar.

// Por isso, o Go tem o identificador em branco, que é um sublinhado (_). Em um
// loop for, é possível utilizar o identificador em branco para ignorar qualquer
// valor retornado da palavra-chave range. Neste caso, queremos ignorar o índice,
// que é o primeiro argumento retornado.

// func main() {
// 	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}

// 	for _, shark := range sharks {
// 		fmt.Println(shark)
// 	}
// }

// Output:
// hammerhead
// great white
// dogfish
// frilled
// bullhead
// requiem

// Esse resultado mostra que o loop for iterou por toda a fatia de strings e imprimiu
// cada item da fatia sem o índice.

// Também é possível usar o range para adicionar itens a uma lista:

// func main() {
// 	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}

// 	for range sharks {
// 		sharks = append(sharks, "shark")
// 	}

// 	fmt.Printf("%q\n", sharks)
// }

// Output:
// ['hammerhead', 'great white', 'dogfish', 'frilled', 'bullhead', 'requiem', 'shark', 'shark', 'shark', 'shark', 'shark', 'shark']

// Aqui, adicionamos uma string com espaço reservado de "shark"para cada item do comprimento
// da fatia sharks.

// Note que não precisamos usar o identificador em branco _ para ignorar nenhum dos
// valores retornados do operador range. O Go nos permite omitir toda a parte da instrução
// range se não precisarmos usar qualquer um do valores retornados.
