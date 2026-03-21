package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 40}
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}

/*
Aula: Templates no Go (html/template)

==============================
VISÃO RÁPIDA
==============================

Templates servem para:

✔ Gerar texto dinamicamente
✔ Gerar HTML dinâmico
✔ Separar lógica de apresentação
✔ Evitar concatenação manual de strings

--------------------------------

Fluxo do código:

Struct Go → Template → Texto final renderizado

==============================
EXPLICANDO O CÓDIGO
==============================

🔹 Struct de dados

type Curso struct {
	Nome string
	CargaHoraria int
}

Representa os dados que serão enviados para o template.

--------------------------------

🔹 Criando Template

tmp := template.New("CursoTemplate")

Cria template vazio em memória.

Nome é só identificador interno.

--------------------------------

🔹 Parse do Template

tmp.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")

Define layout do template.

--------------------------------

🔹 Sintaxe {{ }}

{{ }} = marcador de dados dinâmicos.

. (ponto) = objeto atual (curso).

{{.Nome}} → curso.Nome
{{.CargaHoraria}} → curso.CargaHoraria

--------------------------------

🔹 Execute()

tmp.Execute(os.Stdout, curso)

Renderiza template substituindo valores.

Destino:
os.Stdout → terminal

Poderia ser:
HTTP Response
Arquivo
Buffer

==============================
PARA QUE TEMPLATES SERVEM NA PRÁTICA
==============================

🔹 Web Servers

Gerar páginas HTML dinâmicas.

--------------------------------

🔹 Emails dinâmicos

Ex:
Olá {{.Nome}}, seu pedido {{.Pedido}} foi enviado.

--------------------------------

🔹 Relatórios

PDF / TXT / CSV dinâmicos.

--------------------------------

🔹 Configs dinâmicas

Gerar arquivos de config baseados em dados.

==============================
POR QUE USAR html/template E NÃO text/template?
==============================

html/template:
✔ Protege contra XSS
✔ Escapa HTML automaticamente

text/template:
✔ Para texto puro

==============================
VISÃO SENIOR
==============================

Templates são parte da arquitetura de:

Presentation Layer

--------------------------------

Go segue filosofia:

Business Logic → Go Code
View → Template

--------------------------------

Em sistemas grandes normalmente usamos:

Templates server-side (SSR)
ou
API + Frontend separado (React, etc)

--------------------------------

Hoje em microservices:

Templates são comuns em:
Admin panels
Emails
Relatórios
Ferramentas internas

==============================
RESUMO MENTAL
==============================

Template → Layout com placeholders
{{ }} → Variáveis dinâmicas
Execute → Renderiza template
html/template → Seguro para HTML

==============================
Esse exemplo é a base de SSR no Go.
==============================
*/
