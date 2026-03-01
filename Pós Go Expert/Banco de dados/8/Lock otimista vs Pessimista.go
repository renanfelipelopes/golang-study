package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}
	c.Name = "Eletronicos"
	tx.Debug().Save(&c)
	tx.Commit()
}

/*
	Aula: Lock Otimista vs Lock Pessimista (Controle de Concorrência)

	Nesta aula vamos entender:

	🔥 O que é problema de concorrência no banco
	🔥 O que é Lost Update
	🔥 O que é Lock Pessimista
	🔥 O que é Lock Otimista
	🔥 Como o GORM aplica Lock Pessimista
	🔥 Quando usar cada abordagem

	==========================================================================

	🧠 1️⃣ O PROBLEMA QUE LOCK RESOLVE

	Imagine:

	Duas pessoas atualizando a mesma categoria ao mesmo tempo.

	Processo A lê:
	Name = "Eletronicos"

	Processo B lê:
	Name = "Eletronicos"

	Processo A altera para:
	"Eletronicos e Tecnologia"

	Processo B altera para:
	"Aparelhos"

	Se não houver controle:
	O último que salvar sobrescreve o anterior.

	Isso se chama:
	❌ Lost Update (perda de atualização)

	==========================================================================

	🔒 2️⃣ LOCK PESSIMISTA (Pessimistic Lock)

	Ideia:

	"Eu acredito que alguém pode alterar esse dado,
	então vou bloquear ele enquanto trabalho."

	No seu código:

	tx := db.Begin()

	tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&c, 1)

	O que isso gera no SQL (MySQL):

	SELECT * FROM categories
	WHERE id = 1
	FOR UPDATE;

	FOR UPDATE significa:

	🔐 Bloqueia essa linha até o commit ou rollback.

	Enquanto essa transação não finalizar:
	Outras transações não podem alterar essa linha.

	==========================================================================

	🔁 Fluxo real do seu código

	1) Inicia transação
	2) Executa SELECT ... FOR UPDATE
	3) Linha fica bloqueada
	4) Atualiza Name
	5) Save()
	6) Commit()
	7) Lock é liberado

	Enquanto isso:
	Outro processo tentando UPDATE vai ficar esperando.

	==========================================================================

	📌 Características do Lock Pessimista

	✔ Bloqueia imediatamente
	✔ Garante exclusividade
	✔ Evita conflito
	❌ Pode gerar espera
	❌ Pode causar deadlock
	❌ Reduz escalabilidade

	==========================================================================

	🚀 3️⃣ LOCK OTIMISTA (Optimistic Lock)

	Ideia:

	"Eu acredito que ninguém vai alterar.
	Se alguém alterar, eu detecto o conflito."

	Ele NÃO bloqueia a linha.

	Em vez disso, usa uma coluna de controle,
	geralmente chamada:

	- version
	ou
	- updated_at

	Exemplo clássico:

	type Category struct {
		ID      int
		Name    string
		Version int
	}

	Fluxo:

	1) Processo A lê versão = 1
	2) Processo B lê versão = 1

	Processo A salva:
	UPDATE categories
	SET name = "Novo", version = 2
	WHERE id = 1 AND version = 1;

	Processo B tenta salvar:
	UPDATE categories
	SET name = "Outro", version = 2
	WHERE id = 1 AND version = 1;

	Mas agora version já é 2.
	O update não afeta nenhuma linha.

	Isso indica conflito.

	==========================================================================

	📌 Características do Lock Otimista

	✔ Não bloqueia linha
	✔ Melhor para alta concorrência
	✔ Mais escalável
	❌ Precisa tratar conflito manualmente
	❌ Pode exigir retry

	==========================================================================

	⚔️ 4️⃣ Comparação Direta

	| Lock Pessimista | Lock Otimista |
	|------------------|---------------|
	| Bloqueia linha   | Não bloqueia  |
	| Usa FOR UPDATE   | Usa version   |
	| Seguro imediato  | Detecta conflito depois |
	| Pode travar      | Pode falhar update |
	| Menos escalável  | Mais escalável |

	==========================================================================

	🏗 5️⃣ Quando usar cada um?

	🔒 Use Pessimista quando:

	- Transações financeiras
	- Estoque crítico
	- Baixa concorrência
	- Não pode haver conflito

	🚀 Use Otimista quando:

	- Alta concorrência
	- Sistemas distribuídos
	- Edição de conteúdo
	- APIs REST comuns

	==========================================================================

	🧨 6️⃣ E o Deadlock?

	Lock pessimista pode causar deadlock.

	Exemplo:

	Transação A bloqueia linha 1
	Transação B bloqueia linha 2

	A tenta linha 2
	B tenta linha 1

	🔥 Deadlock.

	O banco precisa abortar uma delas.

	==========================================================================

	🧬 7️⃣ Relação com ACID

	Locks ajudam a garantir:

	Isolation (I do ACID)

	Garantem que transações não interfiram
	de maneira inconsistente.

	==========================================================================

	🏛 8️⃣ Visão arquitetural moderna

	Em sistemas modernos:

	- Monólitos costumam usar mais pessimista
	- Microservices preferem otimista
	- Sistemas altamente distribuídos evitam locks longos
	- Sistemas financeiros usam pessimista estratégico

	==========================================================================

	🧠 RESUMO FINAL

	Concorrência gera conflitos.

	Lock Pessimista:
		Bloqueia antes de alterar.

	Lock Otimista:
		Não bloqueia.
		Detecta conflito depois.

	Seu código usa:

	🔒 Lock Pessimista
		SELECT ... FOR UPDATE

	Isso garante que ninguém altere a Category
	até o commit da transação.

	Controle de concorrência é fundamental
	para sistemas robustos e consistentes.
*/
