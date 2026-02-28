package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}

/*
	Aula: Context no lado do Client

	Se antes vimos Context no Server, agora vamos entender:
	🔥 Context no lado do CLIENT HTTP
	🔥 Como cancelar requisição externa
	🔥 Como evitar travar sua aplicação
	🔥 Como timeout funciona de verdade

	==========================================================================

	🧠 1️⃣ O que está acontecendo aqui?

	Você está criando uma requisição HTTP com um tempo máximo de vida de 10 segundos.

	Se a requisição:
	- Demorar mais que 10 segundos
	- Travar
	- O servidor não responder
	- A conexão ficar pendurada
	Ela será cancelada automaticamente.

	==========================================================================

	🎯 2️⃣ O que é http.NewRequestWithContext?

	Essa função liga o ciclo de vida da requisição ao context.

	Internamente ela faz:
	Request
	↓
	Associa ctx
	↓
	Se ctx cancelar
	↓
	Fecha conexão TCP
	↓
	Interrompe leitura/escrita

	Isso é extremamente importante.

	==========================================================================

	🚨 3️⃣ O problema que isso resolve

	Imagine que você chama uma API externa:
	- API está lenta
	- API travou
	- API nunca responde
	- DNS demorando
	- TLS demorando
	- Conexão ficou aberta

	Sem context:

	❌ Sua aplicação pode ficar presa
	❌ Goroutine bloqueada
	❌ Pool de conexões esgota
	❌ Sistema para de responder

	Com context:

	✔ Timeout automático
	✔ Cancelamento limpo
	✔ Recursos liberados
	✔ Sistema continua saudável

	==========================================================================

	🧬 4️⃣ Fluxo interno real
	Sua aplicação
		↓
	Cria ctx com timeout
		↓
	Cria request com ctx
		↓
	http.Client inicia conexão TCP
		↓
	Envia request
		↓
	Espera resposta

	Se passar de 10 segundos:
	Timer dispara
		↓
	ctx é cancelado
		↓
	Cliente fecha conexão
		↓
	Do() retorna erro

	Erro típico:
		context deadline exceeded

	==========================================================================

	🏗 5️⃣ Onde isso é usado em produção?

	Esse padrão é obrigatório em:
	- Microservices
	- Gateway API
	- Chamadas REST internas
	- Integração com terceiros
	- Sistemas de pagamento
	- Comunicação entre containers
	Se você não colocar timeout, você está aceitando travamento infinito.

	==========================================================================

	🧨 6️⃣ Diferença importante
	Existe também isso:
		http.Client{
			Timeout: 10 * time.Second,
		}

	Mas isso é diferente de context.

	Client.Timeout:
	Timeout global da requisição inteira.

	Context:

	Controle fino e propagável.

	Exemplo real:

	func Service(ctx context.Context) {
		req, _ := http.NewRequestWithContext(ctx, ...)
	}

	Se o request HTTP original for cancelado,
	a chamada externa também será cancelada.

	Isso cria cancelamento em cascata.

	==========================================================================

	🔁 7️⃣ O que acontece com o io.Copy?
	io.Copy(os.Stdout, res.Body)

	res.Body é um stream.

	Se o context for cancelado:
	- A conexão é fechada
	- Read() retorna erro
	- io.Copy para automaticamente

	Sem vazamento.
	Sem travamento.

	==========================================================================

	🧠 8️⃣ Arquitetura mental correta

	Context no client serve para:
		Definir o tempo máximo aceitável para esperar outra aplicação.

		Isso é engenharia de resiliência.

	🧩 9️⃣ Comparação Server vs Client
	| Lado   | Quem cancela  | Para quê                         |
	| ------ | ------------- | -------------------------------- |
	| Server | Cliente       | Evitar processar sem necessidade |
	| Client | Sua aplicação | Evitar ficar preso esperando     |

	==========================================================================

	🔥 🔥 🔥 10️⃣ Cenário real de desastre (sem context)

	Sistema A chama Sistema B.
	Sistema B trava.
	Sistema A:
	- Fica esperando
	- Bloqueia goroutines
	- Pool de conexões esgota
	- CPU sobe
	- Kubernetes mata o container
	- Efeito cascata

	Com context:
		Após 10s:
		- Cancela chamada
		- Libera recurso
		- Retorna erro controlado
		- Sistema continua estável

	==========================================================================
*/
