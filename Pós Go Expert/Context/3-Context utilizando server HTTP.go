package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")
	select {
	case <-time.After(5 * time.Second):
		log.Println("Request processada com sucesso")
		w.Write([]byte("Request processada com sucesso"))
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
		http.Error(w, "Request cancelada pelo cliente", http.StatusRequestTimeout)

	}
}

/*
	Aula: Context utilizando server HTTP

	🧠 1️⃣ O que é r.Context()?

	Quando uma requisição HTTP chega no servidor Go:
	O net/http automaticamente cria um context associado àquela requisição.
	Esse contexto é cancelado automaticamente quando:
	- O cliente fecha o navegador
	- A conexão TCP cai
	- O servidor finaliza a request
	- O timeout do servidor é atingido

	Ou seja:
	👉 O ciclo de vida do context = ciclo de vida da requisição.


	==========================================================================

	🔥 2️⃣ O que está acontecendo nesse exemplo?

	Temos duas possibilidades no select:
	1) A operação demora 5 segundos
	2) O cliente cancela antes

	Se o cliente esperar 5 segundos:
	✔ Request finaliza normalmente.

	Se o cliente fechar o navegador antes:
	✔ ctx.Done() é acionado.
	✔ A request é cancelada.
	✔ O handler para imediatamente.

	==========================================================================

	🎯 3️⃣ Por que isso é importante?

	Imagine que isso aqui fosse:
	- Uma query no banco
	- Uma chamada para API externa
	- Um processamento pesado
	- Geração de relatório
	- Upload grande

	Se o cliente fecha a aba do navegador…

	❌ Faz sentido continuar processando?
	❌ Faz sentido continuar usando CPU?
	❌ Faz sentido continuar ocupando conexão com banco?

	Não.

	O context resolve isso.

	==========================================================================

	🧬 4️⃣ O que aconteceria sem context?

	Se você removesse:
	case <-ctx.Done():

	O servidor continuaria:
	- Processando
	- Gastando CPU
	- Usando memória
	- Mantendo conexão aberta
	- Talvez bloqueando banco

	Mesmo que ninguém esteja mais esperando resposta.
	Isso em produção gera:

	🚨 Vazamento de recursos
	🚨 Exaustão de conexão
	🚨 Lentidão geral do sistema

	==========================================================================

	🏗 5️⃣ Fluxo real interno
	Cliente faz request:
	Browser
	↓
	Servidor Go
	↓
	net/http cria context
	↓
	handler recebe r.Context()

	Se cliente cancela:

	Browser fecha
	↓
	TCP fecha
	↓
	net/http cancela context
	↓
	ctx.Done() dispara
	↓
	handler interrompe

	Isso é cancelamento em cascata.

	==========================================================================

	🧠 6️⃣ Por que usar select?

	Porque select permite esperar múltiplos eventos ao mesmo tempo.
	Aqui ele está esperando:
	- Operação terminar (time.After)
	- Contexto ser cancelado
	O primeiro que responder vence.

	==========================================================================

	🏎 7️⃣ Em produção isso fica assim

	Exemplo real:

	func handler(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		result, err := service.Process(ctx)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write(result)
	}

	E o service:

	func (s *Service) Process(ctx context.Context) ([]byte, error) {
		return repository.Query(ctx)
	}

	E o repository:

	func (r *Repository) Query(ctx context.Context) ([]byte, error) {
		return db.QueryContext(ctx, "SELECT ...")
	}

	Se o cliente cancelar:

	👉 A query do banco é cancelada automaticamente.

	Isso é poderoso.

	==========================================================================

	🧨 8️⃣ Problema clássico resolvido

	Sem context:
	- Cliente faz 10.000 requests
	- Cancela todas
	- Servidor continua processando todas
	- CPU vai a 100%
	- Sistema cai

	Com context:
	- Cliente cancela
	- Processamento para
	- Sistema continua saudável

	==========================================================================

	📦 9️⃣ Utilidade real do context nessa aplicação

	Nesse exemplo simples ele:

	✔ Detecta cancelamento do cliente
	✔ Interrompe processamento
	✔ Evita desperdício de recurso
	✔ Permite escrever sistemas resilientes

	Mesmo que aqui seja só time.After, na vida real isso é:
	- Banco
	- Cache
	- Kafka
	- API externa

	Sistema legado

	==========================================================================

	🏛 🔟 Arquitetura mental correta

	Context em HTTP Server serve para:
		Controlar o tempo de vida da requisição.

	Ele é o "controle mestre".

	Tudo que acontece dentro do handler deve respeitar ele.
*/
