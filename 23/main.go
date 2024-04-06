package main

/* --> Instalando pacotes aula 32

Para instalarmos pacotes no Go, precisamos rodar um comando
- go get enderecoGitHub.com/algumacoisa

Por exemplo, se eu quiser usar uma biblioteca de UUID, basta eu pesquisar no google,
e vou achar o pacote no github:
- go get https://github.com/google/uuid

O go vai criar uma arquivo go.mod e um go.sum

Posso rodar o comando
- go mod tidy

que serve para verificar as versoes, otimizar, e verificar
compatibilidade e se estou usando os pacotes no meu projeto,
senao, ele ira resolver.

Posso inserir o pacote manualmente colocando seu endereco de github no import do projeto.
Só é necessário rodar o - go mod tidy para ele verificar e instalar

*/

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println(uuid.New())
}
