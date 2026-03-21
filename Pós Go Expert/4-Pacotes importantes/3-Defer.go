/*
Aula de Defer em Go

==============================
README DIDÁTICO
==============================

O "defer" é usado para adiar a execução de uma função até o final da função atual.

Ou seja:
➡ O código com defer NÃO executa na hora.
➡ Ele executa quando a função terminar (antes do return final).

Exemplo comum:
- Fechar arquivos
- Fechar conexões HTTP
- Liberar recursos
- Unlock de mutex
- Finalizar transações

Exemplo clássico:
req, err := http.Get(...)
defer req.Body.Close()

Mesmo que aconteça erro depois, o Close será executado.

--------------------------------

Ordem de execução do defer:

Os defers funcionam como uma pilha (Stack - LIFO):

defer A
defer B
defer C

Execução será:
C
B
A

--------------------------------

Exemplo visual:

func main() {
	defer fmt.Print("Primeira linha")
	fmt.Print("Segunda linha")
	fmt.Print("Terceira linha")
}

Saída:
Segunda linha
Terceira linha
Primeira linha

==============================
VISÃO DE ARQUITETURA (NÍVEL SENIOR)
==============================

1️⃣ Gerenciamento de Recursos (Resource Safety)

Em sistemas reais (APIs, workers, mensageria, microservices):
- Conexões HTTP
- Conexões com banco
- Arquivos
- Locks
- Streams

Se esquecer de fechar:
❌ Memory leak
❌ Socket leak
❌ Connection pool esgotado
❌ Travamento de threads
❌ Queda de performance progressiva

O defer garante fechamento mesmo com:
- panic
- return antecipado
- erro intermediário

--------------------------------

2️⃣ Quando o defer é registrado?

O defer é registrado no momento em que é encontrado,
mas executado apenas no final da função.

IMPORTANTE:
Os argumentos são avaliados na hora do defer, não na execução.

--------------------------------

3️⃣ Custo de Performance

Defer tem custo pequeno, mas existe.

Hot paths (código ultra crítico):
Às vezes times evitam defer dentro de loops gigantes.

Mas:
➡ Na maioria dos casos, vale MUITO mais a segurança.

--------------------------------

4️⃣ Boas práticas em Produção

✅ Declarar defer logo após criar o recurso

Exemplo correto:
req, err := http.Get(...)
if err != nil {
	panic(err)
}
defer req.Body.Close()

❌ Evitar:
defer antes de validar erro (pode dar panic nil pointer)

--------------------------------

5️⃣ Defer e Panics

Mesmo se acontecer panic:
➡ Todos os defers são executados
➡ Depois o panic continua propagando

Isso é MUITO importante em serviços críticos.

--------------------------------

6️⃣ Casos Reais no Mercado (Backend / Bancos / Big Tech)

Muito usado em:
- APIs REST
- Consumers Kafka / RabbitMQ
- Workers de processamento
- Serviços de arquivos
- Drivers de banco

--------------------------------

Resumo Mental:

Defer = "Executa isso quando tudo terminar"

Regra de ouro:
Criou recurso → Defer Close

==============================
Este exemplo cobre:
HTTP Request -> Uso de Defer -> Segurança de recurso -> Ordem de execução -> Conceitos de runtime
==============================
*/

package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://www.google.com")
	defer req.Body.Close()
	if err != nil {
		panic(err)
	}
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))
}
