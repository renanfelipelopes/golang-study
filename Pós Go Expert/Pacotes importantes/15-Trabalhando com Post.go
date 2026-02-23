package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "joaquim"}`))
	resp, err := c.Post("http://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}

/*
Aula: Buffer em Programação + Buffer no Go + Buffer em Requisição HTTP

==============================
1️⃣ O QUE É UM BUFFER?
==============================

Buffer é uma área temporária de memória usada para armazenar dados
enquanto eles estão sendo transferidos de um lugar para outro.

Ele serve para:

✔ Evitar acesso direto e constante a disco/rede
✔ Melhorar performance
✔ Controlar fluxo de dados (streaming)
✔ Trabalhar com dados em pedaços (chunks)

--------------------------------
DESENHO MENTAL:

[ Origem ] → [ BUFFER ] → [ Destino ]

Ex:
Rede → Buffer → Sua aplicação
*/

// ============================================================ //

/*
2️⃣ O QUE É BUFFER NO GO?

No Go, buffer normalmente é:
[]byte  (slice de bytes)

Porque:
- Tudo que vem de rede é byte
- Tudo que vai pra rede é byte
- Arquivos são bytes
- JSON é transmitido como bytes

==============================

Por que slice de bytes?
- Porque slice é:
- Dinâmico
- Contíguo em memória
- Eficiente
- Base do IO no Go

Como funciona no core?
- Um slice é basicamente:
struct {
    pointer → memória
    length
    capacity
}

Slice ([]byte)
   |
   v
[ 123 34 110 97 109 101 ... ]

Cada número = 1 byte na memória.

============================================================

3️⃣ bytes.NewBuffer()
No seu código:
jsonVar := bytes.NewBuffer([]byte(`{"name": "joaquim"}`))

Isso cria um buffer que implementa:
io.Reader
io.Writer

Ou seja:
Ele pode ser lido como stream.

Por que usar Buffer aqui?
Porque:
http.Post() precisa de um io.Reader
Assinatura interna simplificada:
func (c *Client) Post(url, contentType string, body io.Reader)

Ele não quer string.
Ele quer algo que possa ser lido em stream.

JSON String
   ↓
[]byte
   ↓
Buffer (io.Reader)
   ↓
http.Client lê aos poucos
   ↓
Envia pela rede

============================================================

4️⃣ O QUE ACONTECE NA REQUISIÇÃO HTTP?

Quando você faz:
c.Post(...)

Internamente:
- O client abre conexão TCP
- Lê o buffer aos poucos
- Envia dados em pacotes
- Fecha envio

Por que não envia tudo de uma vez?
Porque:
- Rede funciona em STREAM.
- Não é bloco único.
- São pedaços.

Buffer memória
   ↓
[ chunk 1 ] → rede
[ chunk 2 ] → rede
[ chunk 3 ] → rede

Isso economiza memória e permite enviar arquivos grandes.

============================================================

5️⃣ O QUE É io.CopyBuffer?
io.CopyBuffer(os.Stdout, resp.Body, nil)

Copia dados de:
resp.Body (io.Reader)
para
os.Stdout (io.Writer)

O que acontece por baixo?
Ele faz algo assim:
- cria []byte temporário
- lê um pedaço
- escreve
- repete até EOF

resp.Body (rede)
	↓
buffer []byte
	↓
os.Stdout

============================================================

6️⃣ POR QUE BUFFER É IMPORTANTE EM HTTP?

Porque HTTP é STREAM.

Se você carregasse tudo na memória:

❌ Pode estourar RAM
❌ Pode travar aplicação
❌ Não escala

Buffer permite:

✔ Streaming
✔ Controle de fluxo
✔ Baixo consumo de memória
✔ Alta performance

============================================================

7️⃣ VISÃO SENIOR (ARQUITETURA REAL)

Buffer é base de:
- Upload de arquivos grandes
- Download de vídeos
- APIs de streaming
- Kafka consumers
- Processamento de dados

Em produção você vê:
io.Copy()
json.NewDecoder()
bufio.Reader
bufio.Writer
bytes.Buffer

Tudo gira em torno de buffer + stream.

============================================================

8️⃣ RESUMO FINAL

Buffer = área temporária de memória
Go usa []byte como base
Rede trabalha com stream
http.Client consome io.Reader
bytes.Buffer transforma []byte em stream

============================================================
*/
