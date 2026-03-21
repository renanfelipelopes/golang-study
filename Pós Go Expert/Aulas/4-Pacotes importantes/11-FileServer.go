package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from blog"))
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}

/*
Aula: Servindo arquivos estáticos + Rotas HTTP

==============================
VISÃO RÁPIDA
==============================

Este código cria um servidor HTTP que:

✔ Serve arquivos da pasta ./public
✔ Possui rota /blog com resposta customizada

--------------------------------

Fluxo:

Browser → Go Server → ServeMux → Handler correto → Response

==============================
CONCEITO 1 — FileServer
==============================

fileServer := http.FileServer(http.Dir("./public"))

Cria um handler que serve arquivos estáticos.

Exemplo estrutura:
public/
 ├ index.html
 ├ style.css
 └ logo.png

Se acessar:
http://localhost:8080/index.html

Go retorna:
./public/index.html

--------------------------------

💡 Uso real:
Frontends estáticos
Landing pages
Arquivos públicos
Downloads

==============================
CONCEITO 2 — http.Dir()
==============================

http.Dir("./public")

Define raiz de arquivos que podem ser acessados.

É um filesystem seguro que impede acesso fora da pasta.

==============================
CONCEITO 3 — ServeMux como Router
==============================

mux := http.NewServeMux()

mux.Handle("/", fileServer)

Tudo que começar com "/" cai no FileServer.

--------------------------------

mux.HandleFunc("/blog", ...)

Rota específica sobrescreve comportamento do "/".

Ordem não importa.
Mux escolhe rota mais específica.

==============================
CONCEITO 4 — Static + Dynamic juntos
==============================

Muito comum em produção:

/ → arquivos estáticos
/api → backend API
/admin → painel

==============================
CONCEITO 5 — log.Fatal()
==============================

log.Fatal(http.ListenAndServe(":8080", mux))

Se servidor falhar:
✔ Loga erro
✔ Finaliza aplicação

Boa prática em entrypoint.

==============================
VISÃO SENIOR (ARQUITETURA REAL)
==============================

🔹 Esse padrão é MUITO usado em:

Microservices simples
BFFs
Serviços internos

--------------------------------

🔹 Em produção normalmente vira:

NGINX / CDN → Static Files
Go Server → API

--------------------------------

🔹 Problema real:

FileServer não tem:
Cache control
Compressão
ETag
Versionamento

Normalmente isso fica no CDN.

--------------------------------

🔹 Segurança

Nunca servir:
.env
config files
secrets

==============================
RESUMO MENTAL
==============================

FileServer → Serve arquivos estáticos
http.Dir → Pasta raiz segura
ServeMux → Decide rota
Handle "/" → Static
Handle "/blog" → Dinâmico
log.Fatal → Fail fast

==============================
Isso já é padrão real de backend Go simples.
==============================
*/
