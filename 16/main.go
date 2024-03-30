package main

// Quando usar Ponteiros aula F024
func main() {
	var1 := 10
	var2 := 20
	println("Valor de var1: ", var1)
	println("Valor de var2: ", var2)
	println("Endereco memoria 1: ", &var1)
	println("Endereco memoria 2: ", &var2)
	soma(&var1, &var2)
	println("Novo valor de var1: ", var1)
	println("Novo valor de var2: ", var2)
	println("Endereco memoria 1: ", &var1)
	println("Endereco memoria 2: ", &var2)

	// Ao usar ponteiros nos parametros, o valor das variaveis podem ser alterados direto na
	// memoria, ou seja, var1 e var2 terão seus valores modificados para 50 e 50. Se eu não
	// tivesse usado ponteiros, eu teria passado uma copia do valor das variaveis, e seus valores
	// permaneceriam os mesmos, mas como eu passei o endereço de memoria, eles serão alterados
	// em suas raizes.
}

func soma(a, b *int) int {
	println("-> ", &a)
	println("-> ", &b)
	*a = 50
	*b = 50
	println("-> ", &a)
	println("-> ", &b)
	return *a + *b
}
