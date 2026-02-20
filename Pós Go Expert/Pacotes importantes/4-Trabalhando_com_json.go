package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int
	Saldo  int
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	jsonPuro := []byte(`{"Numero":2,"Saldo":200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	println(contaX.Saldo)
}

/*
Aula: Trabalhando com JSON em Go

==============================
README DIDÁTICO
==============================

JSON é o formato padrão para troca de dados entre sistemas (APIs, filas, arquivos, etc).

Em Go, trabalhamos com JSON principalmente usando:
- json.Marshal      -> Struct → JSON
- json.Unmarshal    -> JSON → Struct
- json.Encoder      -> Serializa direto para saída (arquivo, stdout, response HTTP)
- json.Decoder      -> Lê JSON direto de streams (request, arquivo grande, etc)

--------------------------------

🔹 Marshal (Struct → JSON)

Converte uma struct para JSON em formato []byte.

Exemplo:
conta := Conta{Numero: 1, Saldo: 100}
jsonBytes, _ := json.Marshal(conta)

Sempre retorna:
[]byte

Por isso normalmente fazemos:
string(jsonBytes)

--------------------------------

🔹 Encoder (Struct → Destino)

Diferente do Marshal:
➡ Não retorna JSON
➡ Escreve direto em um destino (Writer)

Exemplo:
json.NewEncoder(os.Stdout).Encode(conta)

Muito usado em:
- HTTP Response
- Arquivos
- Streams

--------------------------------

🔹 Unmarshal (JSON → Struct)

Converte JSON para struct.

Exemplo:
var conta Conta
json.Unmarshal(jsonBytes, &conta)

IMPORTANTE:
➡ Precisa passar ponteiro (&)
➡ Senão não consegue alterar o valor

--------------------------------

🔹 Tags JSON

Permitem mapear campos da struct para nomes diferentes no JSON.

Exemplo:
type Conta struct {
	Numero int `json:"numero_conta"`
	Saldo  int `json:"saldo_total"`
}

--------------------------------

🔹 Zero Value Problem

Se o JSON tiver nome diferente do campo:
➡ Go não consegue mapear
➡ Campo recebe valor zero

Exemplo:
int → 0
string → ""
bool → false

==============================
VISÃO DE ARQUITETURA (NÍVEL SENIOR)
==============================

1️⃣ JSON como Boundary de Sistema

JSON normalmente é usado nas bordas:
- API Gateway
- REST APIs
- Mensageria
- Integrações externas
- Persistência em logs/eventos

Boa prática:
➡ Converter JSON → Struct → Domain Model
➡ Nunca trabalhar com map[string]interface{} no core

--------------------------------

2️⃣ Encoder vs Marshal em Produção

Marshal:
✔ Fácil
✔ Bom para uso interno
❌ Duplica memória (gera buffer)

Encoder:
✔ Stream direto
✔ Menos memória
✔ Melhor para arquivos grandes
✔ Melhor para APIs

Em APIs de alto throughput:
➡ Prefira Encoder

--------------------------------

3️⃣ Performance Interna

encoding/json usa reflection.

Impactos:
- Mais CPU
- Mais alocação
- Pode ser gargalo em sistemas de alta escala

Alternativas em sistemas ultra críticos:
- jsoniter
- easyjson
- sonic (muito usado em alta performance)

Mas:
➡ encoding/json é padrão seguro e estável

--------------------------------

4️⃣ Segurança e Validação

Nunca confiar direto no JSON externo.

Boa prática:
DTO → Validate → Domain

Exemplo fluxo:
Request JSON
↓
DTO Struct
↓
Validação
↓
Domain Struct
↓
Regra de negócio

--------------------------------

5️⃣ Problemas Reais de Produção

Erros comuns:
❌ Esquecer ponteiro no Unmarshal
❌ Campos não exportados (lowercase não serializa)
❌ Tags erradas
❌ Null inesperado
❌ Tipos diferentes (string vs int)

--------------------------------

6️⃣ Exportação de Campos

JSON só funciona com campos exportados:

✅ Numero int
❌ numero int

--------------------------------

7️⃣ Versionamento de APIs

JSON permite evolução:
- Campos opcionais
- Campos novos
- Backward compatibility

Muito usado em:
- Bancos
- Fintechs
- Marketplaces
- Big Tech APIs

--------------------------------

Resumo Mental:

Marshal → Quero JSON para mim
Encoder → Quero mandar JSON para alguém
Unmarshal → Quero transformar JSON em objeto Go

Regra de ouro:
Entrada externa → Validar sempre

==============================
Este exemplo cobre:
Serialização → Desserialização → Streams → Tags → Performance → Arquitetura de APIs
==============================
*/
