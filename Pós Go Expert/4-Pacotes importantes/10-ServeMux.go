package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.Handler("/blog", blog{title: "My Blog"})
	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}

/*
Aula: ServeMux + Interface http.Handler + Método ServeHTTP

==============================
README DIDÁTICO
==============================

Este programa cria:

1️⃣ Um multiplexer (roteador de rotas HTTP)
2️⃣ Duas rotas:
   "/" → Função simples
   "/blog" → Struct que implementa Handler

--------------------------------

Fluxo:

HTTP Request
   ↓
ServeMux (roteador)
   ↓
Handler correto
   ↓
Response HTTP

--------------------------------

🔹 O que é ServeMux?

ServeMux = HTTP Router do Go.

Ele decide:
➡ Qual handler executar
➡ Baseado no path da URL

--------------------------------

🔹 Criando ServeMux

mux := http.NewServeMux()

Cria roteador independente.
(Não usa DefaultServeMux global)

--------------------------------

🔹 Registrando rotas

mux.HandleFunc("/", HomeHandler)

Registra função como handler.

--------------------------------

🔹 Registrando struct como handler

mux.Handle("/blog", blog{title: "My Blog"})

Aqui entra conceito MUITO IMPORTANTE:

👉 blog implementa interface http.Handler

--------------------------------

🔹 Interface http.Handler

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

Se uma struct tiver ServeHTTP → ela vira Handler automaticamente.

--------------------------------

🔹 Por isso existe:

func (b blog) ServeHTTP(...)

Isso é um METHOD RECEIVER.

--------------------------------

🔹 O que é (b blog)?

Significa:
➡ Método pertence ao tipo blog
➡ b = instância atual (tipo this/self)

--------------------------------

🔹 Analogia:

Java:
this.title

Go:
b.title

--------------------------------

🔹 Como ServeHTTP é chamado automaticamente?

Passo mágico do Go:

1️⃣ Você registra handler:
mux.Handle("/blog", blog{})

2️⃣ ServeMux salva handler

3️⃣ Chega request /blog

4️⃣ ServeMux executa:
handler.ServeHTTP(w, r)

--------------------------------

Você NÃO chama ServeHTTP manualmente.

O net/http chama para você.

==============================
VISÃO DE ARQUITETURA (NÍVEL SENIOR)
==============================

🔹 ServeMux é Pattern:

Router + Dispatcher

--------------------------------

🔹 Interface-Based Design

Go favorece:
Composition + Interfaces

--------------------------------

🔹 Isso permite:

Middleware
Decorators
Chain of Responsibility

--------------------------------

🔹 Exemplo real:

AuthMiddleware → LoggingMiddleware → Handler Final

--------------------------------

🔹 Performance

ServeMux é:
✔ Muito rápido
✔ Thread safe
✔ Sem reflection pesada

--------------------------------

🔹 Por que não usar DefaultServeMux?

Default:
Global state

NewServeMux:
Isolado
Testável
Seguro

--------------------------------

🔹 Pattern usado aqui

Dependency Injection implícita:
http.ListenAndServe(..., mux)

==============================
EXPLICAÇÃO PROFUNDA DO (b blog)
==============================

func (b blog) ServeHTTP(...)

Significa:

👉 Método pertence ao tipo blog
👉 blog agora implementa Handler

--------------------------------

🔹 Receiver = objeto atual

b é cópia do struct blog.

--------------------------------

🔹 Poderia ser ponteiro:

func (b *blog) ServeHTTP()

Quando usar ponteiro?
➡ Struct grande
➡ Precisa modificar estado

--------------------------------

🔹 Aqui poderia ser valor porque:

blog é pequeno
Só leitura

==============================
COMO O SERVEHTTP É EXECUTADO
==============================

Internamente:

ListenAndServe →
Accept conexão →
Cria Request →
Passa pro mux →
Mux acha rota →
Mux chama handler.ServeHTTP()

--------------------------------

Pseudo código interno:

handler := mux.match(path)
handler.ServeHTTP(w, r)

==============================
POR QUE ISSO É PODEROSO?
==============================

Porque você pode criar:

Structs com estado
Middlewares
Routers customizados
API frameworks

==============================
EXEMPLO MENTAL REAL
==============================

Imagine:

type AuthHandler struct {
	next http.Handler
}

func (a AuthHandler) ServeHTTP(...) {
	validar token
	chamar next.ServeHTTP()
}

Isso é middleware real.

==============================
EVOLUÇÃO REAL NO MERCADO
==============================

Junior:
HandleFunc simples

Pleno:
Handlers customizados

Senior:
Middleware chain
Router modular
DI
Observability

Staff:
Framework interno
Tracing distribuído
Gateway pattern

==============================
RESUMO MENTAL
==============================

ServeMux → Router
HandleFunc → Função vira handler
Handle → Struct vira handler
ServeHTTP → Método obrigatório
Receiver (b blog) → Instância atual
net/http → Chama ServeHTTP automaticamente

==============================
Essa aula é FUNDAMENTO de frameworks como:
Gin
Echo
Chi
Fiber (indiretamente)
==============================
*/
