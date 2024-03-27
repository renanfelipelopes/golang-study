package main

// Ponteiros aula F023
func main() {

	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	println(a)

	b := &a
	println(b)  // o resultado será o endereco de memoria de a: 0xc000055f38
	println(*b) // o resultado será: 20. Isso é chamado de 'dereferencing'.
	// Pode ser encontrado aqui: https://go.dev/tour/moretypes/1

	*b = 30
	println(a) // o resultado será: 30.
}
