package main

/* --> Condicionais aula 34
No Go, não existe um if encadeado, apenas if e else, nada de elif, else if....


*/

func main() {
	a := 1
	b := 2
	c := 3

	if a > b {
		println(a)
	} else {
		println(b)
	}

	if a > b || c < b && b == a {
		println("Amostra de operadores logicos")
	}

	// temos tbm o switch case
	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default:
		println("default")
	}
}
