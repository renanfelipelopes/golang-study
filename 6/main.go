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

//trabalhando com slice
func main() {
	// os colchetes em significa que estou trabalhando com slices, que é um array por debaixo dos panos
	// s := []int{2, 4, 6, 8, 10}
	// len = tamanho, cap = capacidade
	// s[:0] significa que a partir do inicio do slice para toda a direita, mostre zero itens
	s := []int{10, 20, 30, 50, 60, 70, 80, 90, 100}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	// s = append(s, 12)
	// fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
