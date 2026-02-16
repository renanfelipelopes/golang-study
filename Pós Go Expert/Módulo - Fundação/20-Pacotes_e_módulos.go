package main

/* --> Pacotes e módulos aula 29 e 30

Quando vamos iniciar o desenvolvimento de qualquer aplicação em Go, é recomendado
que criemos o Go Mod, assim, evitaremos do Go ficar olhar para a pasta padrão
de projetos (src), e passará a observar o modulo do projeto.
É interessante usar o link do github do projeto como nome do modulo,
assim, é fácil instalar o projeto.
para criar o modulo, rodamos o comando:
- go mod init (seguido do nome do seu modulo, podendo ser o diretorio do github)
exemplo: go mod init github.com/renanfelipelopes/golang-study

Para atribuir um nome usando o github, pode rodar o comando:
- go mod init github.com/renanfelipelopes/golang-study
Se repararmos agora, teremos um arquivo chamado go.mod no final dos arquivos do projeto.

Feito isso, o Go passará a buscar os arquivos dentro da nossa própria pasta e não mais na
usr/local/go/src

Agora podemos pegar o valor do 'diretorio' da nossa pasta no arquivo go.mod e colocar no import.
Pode ser que nosso import fique vermelho, mas vai funcionar. Podemos rodar o comando
- go mod tidy
O nosso import ficará assim:

*/

import (
	"curso-go/matematica"
	"fmt"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Println("Resultado:", s)
}
