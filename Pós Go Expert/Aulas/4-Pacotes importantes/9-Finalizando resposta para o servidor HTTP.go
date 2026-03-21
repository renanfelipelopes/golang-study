package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cep, error := BuscaCep(cepParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cep)
}

func BuscaCep(cep string) (*ViaCEP, error) {
	resp, error := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}
	var c ViaCEP
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error
	}
	return &c, nil
}

/*
Aula: HTTP Server + Handler + Integração com API Externa + JSON Response

==============================
README DIDÁTICO
==============================

Este programa cria:

1️⃣ Um servidor HTTP local
2️⃣ Um endpoint GET
3️⃣ Recebe CEP via query string
4️⃣ Consulta API ViaCEP
5️⃣ Retorna JSON para quem chamou

--------------------------------

Fluxo completo:

Cliente HTTP
    ↓
Servidor Go (Handler)
    ↓
Função BuscaCep()
    ↓
API ViaCEP
    ↓
Struct Go
    ↓
JSON Response para Cliente

--------------------------------

🔹 Subindo servidor HTTP

http.HandleFunc("/", BuscaCepHandler)
http.ListenAndServe(":8080", nil)

HandleFunc:
Registra função handler para rota "/"

ListenAndServe:
Inicia servidor HTTP na porta 8080. Aqui o Go sobe algo que chamamos de multiplexer,
que é um componente onde você "atacha" rotas nele.

--------------------------------

🔹 O que é Handler?

Função que recebe request HTTP e retorna response HTTP.

func BuscaCepHandler(w http.ResponseWriter, r *http.Request)

w → Resposta que vamos enviar
r → Requisição recebida

--------------------------------

🔹 Validando Path

if r.URL.Path != "/"

Evita acessar rota inválida.

--------------------------------

🔹 Query Params

cepParam := r.URL.Query().Get("cep")

Exemplo chamada:
http://localhost:8080/?cep=01001000

--------------------------------

🔹 Status HTTP

400 → Bad Request
404 → Not Found
500 → Internal Error
200 → OK

--------------------------------

🔹 Response Header

w.Header().Set("Content-Type", "application/json")

Define tipo do retorno.

--------------------------------

🔹 json.NewEncoder(w).Encode()

Serializa struct direto no Response.

Melhor que Marshal + Write manual.

--------------------------------

🔹 Função BuscaCep()

Responsável por:
➡ Chamar API externa
➡ Converter JSON → Struct
➡ Retornar ponteiro da struct

--------------------------------

🔹 Por que retornar ponteiro?

return &c

Evita cópia da struct
Melhor performance

--------------------------------

🔹 defer resp.Body.Close()

Garante liberar conexão HTTP.

==============================
ARQUITETURA POR TRÁS (NÍVEL SENIOR)
==============================

🔹 Camadas Implícitas no Código

Controller Layer:
BuscaCepHandler

Service Layer:
BuscaCep()

External Integration Layer:
http.Get(ViaCEP)

--------------------------------

🔹 Problema Arquitetural Atual

Handler chama Service diretamente.
Service chama API diretamente.

Produção ideal:
Controller → UseCase → Service → Gateway → External API

--------------------------------

🔹 Problema Real #1 — Timeout

http.Get NÃO tem timeout.

Produção:
http.Client{ Timeout: 3 * time.Second }

--------------------------------

🔹 Problema Real #2 — Falta de Context

Sem context:
Request pode ficar pendurada.

Produção:
ctx := r.Context()

--------------------------------

🔹 Problema Real #3 — Falta Observabilidade

Produção precisa:
Logs estruturados
Tracing
Metrics

--------------------------------

🔹 Problema Real #4 — DTO vs Domain

ViaCEP deveria ser DTO.
Converter para Domain Model.

--------------------------------

🔹 Problema Real #5 — Segurança

Validar CEP:
Regex
Length
Sanitização

--------------------------------

🔹 Problema Real #6 — Escalabilidade

Hoje:
1 request → 1 chamada externa

Produção:
Cache Redis
Rate Limit
Circuit Breaker

--------------------------------

🔹 Problema Real #7 — ioutil.ReadAll()

Deprecated.
Hoje usar:
io.ReadAll()

--------------------------------

🔹 Problema Real #8 — Status Code Perdido

Não validamos:
resp.StatusCode

API pode retornar erro.

--------------------------------

🔹 Problema Real #9 — Error Handling Fraco

Hoje:
return nil, error

Produção:
Wrap error
Categorizar erro

--------------------------------

🔹 Problema Real #10 — Conexão HTTP

Produção:
Reusar http.Client
Connection pooling

==============================
EXPLICAÇÃO LINHA A LINHA (PARTES IMPORTANTES)
==============================

🔹 Registro de rota

http.HandleFunc("/", BuscaCepHandler)

Toda requisição "/" cai nesse handler.

--------------------------------

🔹 Servidor HTTP

http.ListenAndServe(":8080", nil)

nil = usa DefaultServeMux.

--------------------------------

🔹 Query Param

r.URL.Query().Get("cep")

Extrai query string.

--------------------------------

🔹 Encoder direto no Response

json.NewEncoder(w).Encode(cep)

Evita buffer intermediário.

--------------------------------

🔹 Chamada externa

http.Get("http://viacep.com.br/ws/" + cep + "/json/")

--------------------------------

🔹 Unmarshal JSON

json.Unmarshal(body, &c)

Precisa ponteiro.

--------------------------------

🔹 Retorno ponteiro

return &c, nil

==============================
COMO TESTAR
==============================

Rodar:
go run main.go

Abrir navegador:
http://localhost:8080/?cep=01001000

==============================
EVOLUÇÃO NATURAL DESSE CÓDIGO
==============================

Junior:
Server simples (esse)

Pleno:
Timeout + Context + Logs + Status Code check

Senior:
DTO + Domain + Retry + Client Pool + Metrics

Staff:
Circuit Breaker + Cache + Observability + Tracing + Rate Limit

==============================
RESUMO MENTAL
==============================

HandleFunc → Endpoint
Handler → Controller
BuscaCep → Service
http.Get → Gateway externo
Unmarshal → JSON → Struct
Encoder → Struct → JSON Response

==============================
Esse código já é BASE de microserviço real.
==============================
*/
