// Aula 2: Realizando chamada HTTP

package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))
	req.Body.Close()
}

/*
Aula de Requisições HTTP em Go

Este código demonstra como realizar uma chamada HTTP simples utilizando a biblioteca padrão do Go.

Fluxo do programa:

1. Requisição HTTP GET
   - Usa http.Get para fazer uma requisição para "https://www.google.com".
   - Retorna a resposta do servidor (Response) e um possível erro.
   - Sempre tratar erro em operações de rede.

2. Leitura do corpo da resposta
   - Usa io.ReadAll para ler todo o conteúdo retornado pelo servidor.
   - O conteúdo vem como []byte (bytes).
   - Depois é convertido para string para exibição.

3. Exibição do conteúdo
   - Usa println para imprimir o HTML retornado pelo site.

4. Fechamento do Body da resposta
   - req.Body.Close() libera recursos da conexão HTTP.
   - Boa prática obrigatória para evitar vazamento de conexões.

Conceitos importantes:

- HTTP é baseado em requisição e resposta.
- O Body da resposta é um stream (fluxo de dados).
- Sempre fechar o Body após leitura.
- io.ReadAll carrega tudo na memória → bom para respostas pequenas.
- Para respostas grandes, usar leitura em stream (sem carregar tudo).

Boas práticas que poderiam ser aplicadas futuramente:

- Usar defer req.Body.Close() logo após verificar erro.
- Validar status code (ex: 200 OK).
- Criar client HTTP customizado com timeout.

Este exemplo cobre:
Requisição HTTP -> Leitura da resposta -> Conversão para string -> Fechamento da conexão
*/
