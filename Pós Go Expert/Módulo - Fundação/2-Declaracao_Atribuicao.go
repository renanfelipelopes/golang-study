package main

const a = "Hello, World!"

/*
podemos declarar uma variavel fortemente tipada
nesse caso abaixo, estamos declarando varias variaveis
e atribuindo um tipo para cada uma a nivel global de escopo
*/
var (
	b bool
	c int
	d string
	e float64
)

func main() {
	/*
		mas podemos declarar uma variavel usando o operador := um "short hand"
		que significa que estamos criando uma variavel, declarando e atribuindo uma
		tipagem pelo valor, porém, só podemos usar o := na primeira vez quando se cria a variavel
		se fizermos:
			b := true
			b = 25
		ocorrerá um erro, pois declaramos a variavel b como sendo do tipo bool e tentamos atribuir
		um valor posterior do tipo int
	*/
	b := true

	println(b)
}
