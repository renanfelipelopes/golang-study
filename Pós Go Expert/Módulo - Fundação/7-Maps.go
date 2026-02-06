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

// trabalhando com maps aula F015
func main() {
	salarios := map[string]int{"Renan": 1000, "Joao": 500, "Maria": 250}
	delete(salarios, "Renan")

	// salar := make(map[string]int)
	// salar1 := map[string]int{}
	// salar1["RenanF"] = 1500
	// fmt.Println(salar)
	// fmt.Println(salar1)

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	// se eu quiser exibir apenas o salario e o nome não, eu posso usa o 'blank identifier' que é o underline _:
	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}
