package main

import (
	"fmt"

	"github.com/devfullcycle/goexpert/Packaging/1/math"
)

func main() {
	m := math.NewMath(1, 2)
	m.C = 3
	fmt.Println(m)
	fmt.Println(m.C)
}
