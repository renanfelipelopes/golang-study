package main

import "github.com/devfullcycle/goexpert/Packaging/3/math"

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
}

/*
	go mod edit -replace github.com/devfullcycle/goexpert/Packaging/3/math=../math
	go mod tidy
*/
