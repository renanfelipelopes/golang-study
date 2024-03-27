package main

// Ponteiros aula F023
func main() {

	// Em algum lugar da Memoria, existe um Endereco que tem um Valor
	// Toda vez que eu digo que a := 10, o Go abre uma 'caixinha' na memoria
	// guarda o valor de a, e da um endereco a essa 'caixinha'. Toda vez que
	// precisar buscar esse valor de a, é só ir no endereco de caixinha.
	// Toda vez que eu quiser saber o endereçamento de memoria dessa caixinha de a,
	// eu posso dar um printl(&a), ou seja, passar um 'e comercial' antes do
	// nome da variavel, o resultado sera algo parecido com 0xc000055f38
	a := 10
	println(&a)

	// Vamos dizer que eu queira alterar o valor que está dentro da variavel a
	// direto na memoria, acessando diretamento o endereco da variavel a, abrindo
	// a caixinha e colocar mais coisas ali dentro.
	// Para isso, eu posso fazer o seguinte:
	var ponteiro *int = &a

	// toda vez que usar um asteriscos, eu estou apotando para o endereco da memoria: 0xc000055f38
	// se eu der um println no ponteiro, ele ira retornar o endereço da memoria: 0xc000055f38
	println(ponteiro)

	// variavel -> ponteiro que tem um endereco na memoria -> valor
}
