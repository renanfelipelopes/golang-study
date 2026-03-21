package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}
		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
		}
		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.Uf))
	}
}

/*
Aula: Busca CEP via HTTP + JSON + CLI + Escrita em Arquivo

==============================
README DIDÁTICO
==============================

Este programa faz:

1️⃣ Recebe CEPs via linha de comando
2️⃣ Faz requisição HTTP para API ViaCEP
3️⃣ Converte JSON → Struct Go
4️⃣ Extrai dados relevantes
5️⃣ Salva resultado em arquivo

Fluxo geral:

CLI Args → HTTP Request → JSON → Struct → Arquivo TXT

--------------------------------

🔹 Entrada via Linha de Comando (os.Args)

os.Args contém:
[0] Nome do executável
[1...] Argumentos passados

Exemplo execução:
go run main.go 01001000 22290040

os.Args será:
[main.go, 01001000, 22290040]

--------------------------------

🔹 Entendendo:
for _, cep := range os.Args[1:]

range → percorre lista
_ → ignora índice
cep → valor do item

os.Args[1:] → pega só os CEPs (ignora nome do programa)

--------------------------------

🔹 Struct + JSON Tags

type ViaCEP struct {
	Cep string `json:"cep"`
}

Tags dizem:
➡ Como mapear JSON → Struct
➡ JSON usa "cep"
➡ Go usa Cep

Sem tag:
➡ Go tentaria mapear "Cep"
➡ JSON tem "cep"
➡ Resultado: zero value

--------------------------------

🔹 HTTP GET

req, err := http.Get(url)

Faz requisição HTTP GET.
Retorna:
Response + Error

--------------------------------

🔹 io.ReadAll()

Lê todo Body da resposta HTTP.

Retorna []byte.

--------------------------------

🔹 json.Unmarshal()

Converte:
JSON → Struct Go

Precisa ponteiro:
&data

Porque:
➡ Unmarshal precisa alterar valor da variável.

--------------------------------

🔹 fmt.Sprintf()

Cria string formatada sem imprimir.

Exemplo:
texto := fmt.Sprintf("CEP: %s", cep)

--------------------------------

🔹 fmt.Fprintf()

Escreve formatado em qualquer Writer.

Exemplo:
fmt.Fprintf(os.Stderr, "Erro: %v", err)

--------------------------------

🔹 os.Stderr

Saída padrão de erro.

Boa prática:
Logs → Stderr
Output → Stdout

--------------------------------

🔹 Escrita em Arquivo

file.WriteString()

Escreve string no arquivo.

--------------------------------

🔹 Defer

Garante fechamento do recurso.

defer req.Body.Close()
defer file.Close()

==============================
VISÃO DE ARQUITETURA (NÍVEL SENIOR)
==============================

1️⃣ Boundary Pattern

API externa = Boundary do sistema.

Boa prática:
JSON → DTO → Domain Model

Nunca usar struct externa direto na regra de negócio.

--------------------------------

2️⃣ Streaming vs Buffer

Aqui usamos:
io.ReadAll()

Problema:
➡ Carrega tudo em memória

Em produção:
Preferir:
json.NewDecoder(req.Body).Decode(&data)

--------------------------------

3️⃣ Observabilidade

fmt.Fprintf(os.Stderr)

Simples, mas em produção usar:
Structured Logs
Correlation ID
Tracing

--------------------------------

4️⃣ Problema Real: Defer dentro de Loop

Aqui temos:

for {
    defer req.Body.Close()
}

Isso acumula defer até função terminar.

Em produção:
Fechar manualmente ou usar função separada.

--------------------------------

5️⃣ Idempotência e Retry

APIs externas falham.

Produção precisa:
Retry
Timeout
Circuit Breaker

--------------------------------

6️⃣ Segurança

Nunca confiar em dados externos.

Validar:
CEP formato
Campos obrigatórios
Campos inesperados

--------------------------------

7️⃣ Concorrência (Evolução Natural)

Aqui é sequencial.

Escalável seria:
Goroutines + Worker Pool

--------------------------------

8️⃣ Escrita de Arquivo em Produção

Problemas comuns:
Concorrência
Lock de arquivo
Sobrescrita

Soluções:
Append Mode
Fila de escrita
Storage externo

--------------------------------

9️⃣ Separação por Camadas

CLI Layer
HTTP Client Layer
DTO Layer
Service Layer
Persistence Layer

--------------------------------

10️⃣ Evolução Natural desse Código

Junior → Código atual
Pleno → Timeout + Retry + Logs
Senior → Pool + Observabilidade + DTO + Domain
Staff → Circuit Breaker + Metrics + Cache

==============================
EXPLICAÇÃO DETALHADA DO CÓDIGO
==============================

🔹 Loop dos argumentos

for _, cep := range os.Args[1:]

Processa múltiplos CEPs.

--------------------------------

🔹 Construção URL

"http://viacep.com.br/ws/" + cep + "/json/"

Alternativa melhor:
fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

--------------------------------

🔹 Unmarshal

var data ViaCEP
json.Unmarshal(res, &data)

Converte JSON → Struct.

--------------------------------

🔹 Escrita formatada no arquivo

fmt.Sprintf(
  "CEP: %s, Localidade: %s, UF: %s",
  data.Cep,
  data.Localidade,
  data.Uf,
)

--------------------------------

🔹 Fprintf para erros

fmt.Fprintf(os.Stderr, "Erro: %v", err)

Permite separar log de erro da saída normal.

==============================
RESUMO MENTAL
==============================

os.Args → Entrada CLI
HTTP → Busca dados
ReadAll → Bytes
Unmarshal → Struct
Sprintf → Monta string
Fprintf → Log em destino específico
Defer → Fecha recurso com segurança

==============================
Esse exemplo é MUITO próximo do mundo real.
CLI + HTTP + JSON + Arquivo + Tratamento de erro
==============================
*/
