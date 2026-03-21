package main

import "github.com/devfullcycle/goexpert/Packaging/3/math"

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
}

/*
	Na pasta raiz do projeto (na pasta 4 no caso) rodar o comando:
	go work init ./math ./sistema

	Porém, se adicionarmos bibliotecas externas, o go mod tidy não irá baixar,
	pois o go work trava para apenas bibliotecas no nosso workspace.

	Nessa caso, podemos publicar nossa biblioteca no github e após isso
	rodar o go mod tidy, ou rodar o comando:
		- go mod tidy -e

	Esse comando ignora as bibliotecas que ele não consegue baixar (as nossas
	do workspace no caso) e baixar as externas (como por exemplo a do github.com/google/uuid)

*/
