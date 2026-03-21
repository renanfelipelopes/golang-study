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

func main() {
	fmt.Printf("O tipo de E é %T", e) //imprime o tipo %T
	fmt.Printf("O tipo de E é %v", e) //imprime o valor %v
}
