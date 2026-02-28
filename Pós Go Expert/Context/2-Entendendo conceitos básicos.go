package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	// Criamos um contexto com timeout de 3 segundos
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel book cancelled. Timeout reached.")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Hotel booked.")
	}
}

/*
Aula: Context + Timeout + Select

========================================
1️⃣ O QUE É "BOOK DE HOTEL"?
========================================

"Book de hotel" significa realizar a reserva de um hotel.

Em sistemas reais isso pode envolver:
- Validar disponibilidade
- Cobrar cartão
- Confirmar com sistema externo
- Registrar no banco

Ou seja:
É uma operação que pode demorar.

----------------------------------------

2️⃣ O PROBLEMA QUE ESTAMOS SIMULANDO
========================================

Imagine:

Temos 10 minutos para confirmar uma reserva.
Se passar desse tempo, a reserva expira.
Não faz sentido continuar tentando reservar.

Aqui no exemplo:
- A operação demora 5 segundos
- Mas o timeout é de 3 segundos

Logo:
A operação deve ser cancelada antes de terminar.

----------------------------------------

3️⃣ ENTENDENDO O CONTEXT COM TIMEOUT
========================================

ctx, cancel := context.WithTimeout(ctx, 3*time.Second)

Isso significa:

"Essa operação só pode viver por 3 segundos."

Depois disso:
ctx.Done() será acionado automaticamente.

Internamente:

- WithTimeout cria um timer
- Quando o tempo expira
- Ele cancela o contexto
- Fecha o canal Done()

----------------------------------------

4️⃣ O QUE É ctx.Done()?
========================================

ctx.Done() retorna um canal (<-chan struct{})

Esse canal:

✔ Fica bloqueado normalmente
✔ É fechado quando o contexto é cancelado
✔ Pode ser usado dentro de select

----------------------------------------

5️⃣ ENTENDENDO O SELECT
========================================

Select NÃO é função.
É um statement (estrutura de controle).

Ele serve para esperar múltiplos canais ao mesmo tempo.

Exemplo simplificado:

select {
case <-canal1:
    // faz algo
case <-canal2:
    // faz outra coisa
}

Ele executa o primeiro canal que responder.

----------------------------------------

6️⃣ O QUE ESTÁ ACONTECENDO NO EXEMPLO
========================================

select {
case <-ctx.Done():
    fmt.Println("Cancelado")
case <-time.After(5 * time.Second):
    fmt.Println("Reservado")
}

Temos duas possibilidades:

1) O timeout de 3 segundos dispara primeiro
2) O time.After(5s) dispara primeiro

Como 3 < 5,
o contexto cancela antes.

Resultado:
"Hotel book cancelled. Timeout reached."

----------------------------------------

7️⃣ FLUXO VISUAL DA EXECUÇÃO
========================================

main
  ↓
cria context com 3s
  ↓
bookHotel()
  ↓
select esperando:
    - ctx.Done()
    - time.After(5s)

Após 3 segundos:
ctx é cancelado
  ↓
select executa primeiro case
  ↓
retorna

----------------------------------------

8️⃣ SIMULAÇÃO REAL DE PRODUÇÃO
========================================

Esse padrão é usado para:

- Chamadas HTTP externas
- Queries no banco
- Pagamentos
- Sistemas de reserva
- Integrações com terceiros

Exemplo real:

func bookHotel(ctx context.Context) error {
	req, _ := http.NewRequestWithContext(ctx, "POST", "...", nil)
	resp, err := http.DefaultClient.Do(req)
	return err
}

Se o timeout estourar:
A requisição HTTP é cancelada automaticamente.

----------------------------------------

9️⃣ ERRO COMUM DE INICIANTES
========================================

Não usar context na assinatura:

Errado:
func bookHotel() { }

Correto:
func bookHotel(ctx context.Context) { }

Sempre propague o context.

----------------------------------------

🔟 VISÃO SENIOR
========================================

Context cria cancelamento em cascata.

Se tivermos:

Request HTTP
   ↓
Service
   ↓
Repository
   ↓
Driver de banco

Se o timeout expirar no topo:
Tudo abaixo para automaticamente.

Isso evita:

❌ Vazamento de goroutines
❌ Conexões presas
❌ Processamento inútil
❌ Consumo desnecessário de CPU

----------------------------------------

RESUMO FINAL
========================================

Neste exemplo aprendemos:

✔ Como criar um contexto com timeout
✔ Como o select escuta múltiplos canais
✔ Como ctx.Done() sinaliza cancelamento
✔ Como controlar tempo de vida de operações

Context não é opcional.
Ele é parte da arquitetura Go moderna.

========================================
*/
