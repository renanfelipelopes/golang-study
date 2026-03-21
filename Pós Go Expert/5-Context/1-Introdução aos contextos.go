/*
Aula: Context (Conceitos Básicos)

==============================
O QUE É CONTEXT?
==============================

Context é um mecanismo do Go para:

✔ Cancelar operações
✔ Definir timeout
✔ Controlar ciclo de vida de requisições
✔ Transportar informações entre camadas (request-scoped data)

--------------------------------

ANALOGIA SIMPLES

Imagine:

Você está fazendo uma tarefa.
Alguém diz:
"Pode parar, já resolvemos."

Ou:
"Você tem 10 minutos pra terminar."

Se o tempo acabar ou alguém cancelar,
você para.

Context faz exatamente isso no código.

--------------------------------

PROBLEMA QUE ELE RESOLVE

Sem context:

- Requisição HTTP pode ficar travada
- Goroutine pode nunca parar
- Conexão pode ficar aberta indefinidamente
- Recursos podem vazar

Context permite controle.

==============================
1️⃣ CANCELAMENTO MANUAL
==============================

Exemplo básico:

ctx, cancel := context.WithCancel(context.Background())

go func() {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Parando execução")
            return
        default:
            fmt.Println("Executando...")
            time.Sleep(time.Second)
        }
    }
}()

time.Sleep(3 * time.Second)
cancel()

--------------------------------

O que acontece?

- ctx.Done() é um canal
- Quando cancel() é chamado
- Done é fechado
- Todas goroutines escutando param

--------------------------------

DESENHO MENTAL:

Main
  ↓
cancel()
  ↓
ctx.Done() fecha
  ↓
goroutines param

==============================
2️⃣ TIMEOUT
==============================

ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

select {
case <-time.After(5 * time.Second):
    fmt.Println("Terminou")
case <-ctx.Done():
    fmt.Println("Tempo esgotado")
}

--------------------------------

Se passar de 2 segundos:
ctx cancela automaticamente.

--------------------------------

Uso real:

HTTP request
DB query
Chamada externa

==============================
3️⃣ CONTEXT EM HTTP (MUITO IMPORTANTE)
==============================

Em servidores HTTP:

func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    select {
    case <-time.After(5 * time.Second):
        fmt.Println("Processou")
    case <-ctx.Done():
        fmt.Println("Cliente cancelou a requisição")
    }
}

Se cliente fechar navegador:
Context é cancelado automaticamente.

--------------------------------

Isso evita:
❌ Processamento desnecessário
❌ Vazamento de recurso

==============================
4️⃣ CONTEXT COM KEY/VALUE
==============================

Permite transportar dados entre camadas.

ctx := context.WithValue(context.Background(), "userID", 123)

valor := ctx.Value("userID")

--------------------------------

Uso real:

- Correlation ID
- User ID
- Token
- Trace ID
- Request ID

--------------------------------

Muito usado em:

Tracing distribuído
Logs estruturados
Observabilidade

==============================
⚠️ BOAS PRÁTICAS IMPORTANTES
==============================

1️⃣ Nunca usar context para dados de negócio
2️⃣ Use apenas para dados da requisição
3️⃣ Sempre propagar context entre funções
4️⃣ Sempre chamar cancel() quando usar WithCancel/WithTimeout

--------------------------------

Errado:
func BuscaCep() { }

Correto:
func BuscaCep(ctx context.Context) { }

==============================
VISÃO SENIOR (ARQUITETURA REAL)
==============================

Context é base de:

- Microservices
- Mensageria (Kafka, RabbitMQ)
- APIs REST
- gRPC
- Banco de dados

--------------------------------

Fluxo real em produção:

Request chega
  ↓
Server cria context
  ↓
Handler recebe context
  ↓
Service recebe context
  ↓
Repository recebe context
  ↓
Driver do banco usa context
  ↓
Se timeout → tudo cancela em cascata

--------------------------------

DESENHO DE CANCELAMENTO EM CASCATA:

Request
  ↓
Service
  ↓
Repository
  ↓
DB

Se cancelar no topo,
todos abaixo param.

==============================
RESUMO MENTAL
==============================

Context serve para:

✔ Cancelar
✔ Definir timeout
✔ Transportar dados da requisição
✔ Controlar ciclo de vida

Sem context:
Sistema não escala direito.

Com context:
Sistema fica previsível, controlável e seguro.

==============================
Essa é uma das peças mais importantes da arquitetura Go moderna.
==============================
*/