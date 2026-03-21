package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	// create category
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	category2 := Category{Name: "Cozinha"}
	db.Create(&category2)

	// create product
	db.Create(&Product{
		Name:       "Panela",
		Price:      99.00,
		Categories: []Category{category, category2},
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			println("- ", product.Name)
		}
	}
}

/*
	Aula: Many-to-Many com GORM (Product ↔ Category)

	Nesta aula vamos entender:

	🔥 O que é relacionamento Many-to-Many
	🔥 Como o GORM cria as tabelas automaticamente
	🔥 Como struct vira tabela SQL
	🔥 O que é tabela de junção (join table)
	🔥 Como funciona o Preload
	🔥 Como o banco realmente armazena essa relação

	==========================================================================

	🧠 1️⃣ O que é Many-to-Many?

	Many-to-Many (N:N) significa:

	- Um produto pode ter várias categorias
	- Uma categoria pode ter vários produtos

	Exemplo real:
	Produto: "Panela"
	Categorias: "Cozinha" e "Promoção"

	Categoria: "Eletrônicos"
	Pode ter:
	TV, Notebook, Fone, etc.

	Ou seja:
	Múltiplos ↔ Múltiplos

	==========================================================================

	🧱 2️⃣ Como as structs viram tabelas SQL?

	Struct Category:

	type Category struct {
		ID       int `gorm:"primaryKey"`
		Name     string
		Products []Product `gorm:"many2many:products_categories;"`
	}

	Struct Product:

	type Product struct {
		ID         int `gorm:"primaryKey"`
		Name       string
		Price      float64
		Categories []Category `gorm:"many2many:products_categories"`
		gorm.Model
	}

	O GORM usa reflection para ler as structs
	e gerar as tabelas automaticamente com AutoMigrate.

	==========================================================================

	🗄 3️⃣ Quais tabelas são criadas no banco?

	db.AutoMigrate(&Product{}, &Category{})

	Isso gera:

	1) products
	2) categories
	3) products_categories (tabela de junção)

	------------------------------------------

	📌 Tabela: products

	id (PK)
	name
	price
	created_at
	updated_at
	deleted_at

	------------------------------------------

	📌 Tabela: categories

	id (PK)
	name

	------------------------------------------

	📌 Tabela: products_categories (JOIN TABLE)

	product_id (FK)
	category_id (FK)

	Essa tabela NÃO tem ID próprio.
	Ela só guarda os relacionamentos.

	==========================================================================

	🔗 4️⃣ Como funciona a relação internamente?

	Quando você faz:

	db.Create(&Product{
		Name:  "Panela",
		Price: 99.00,
		Categories: []Category{category, category2},
	})

	O GORM faz:

	1) Insere produto na tabela products
	2) Pega ID do produto
	3) Pega ID das categorias
	4) Insere na tabela products_categories:

	Exemplo:

	product_id | category_id
	-----------|------------
	1          | 1
	1          | 2

	Ou seja:
	O produto 1 pertence às categorias 1 e 2.

	==========================================================================

	🧬 5️⃣ Relação visual no banco

	products
	--------------------------------
	id | name
	1  | Panela

	categories
	--------------------------------
	id | name
	1  | Eletronicos
	2  | Cozinha

	products_categories
	--------------------------------
	product_id | category_id
	1          | 1
	1          | 2

	Isso é o coração do Many-to-Many.

	==========================================================================

	🔎 6️⃣ O que faz o Preload?

	err = db.Model(&Category{}).
		Preload("Products").
		Find(&categories).Error

	Preload faz EAGER LOADING.

	Significa:
	Carregar as categorias
	E já carregar os produtos relacionados.

	Internamente o GORM faz:

	SELECT * FROM categories;
	SELECT * FROM products
	JOIN products_categories
	ON products.id = products_categories.product_id
	WHERE products_categories.category_id IN (...);

	Depois ele monta os structs automaticamente.

	==========================================================================

	🔁 7️⃣ Como o GORM reconstrói a relação?

	Ele pega os resultados do JOIN
	E popula:

	category.Products

	Então quando você faz:

	for _, category := range categories {
		fmt.Println(category.Name)
		for _, product := range category.Products {
			println("-", product.Name)
		}
	}

	Ele já tem tudo carregado em memória.

	==========================================================================

	🏗 8️⃣ Por que usamos many2many:"products_categories"?

	Esse trecho:

	`gorm:"many2many:products_categories;"`

	Informa ao GORM:

	"Crie ou use uma tabela de junção chamada products_categories"

	Se você não especificar,
	o GORM cria um nome automático.

	Boa prática:
	Sempre definir explicitamente.

	==========================================================================

	🧨 9️⃣ Pontos importantes

	✔ Many-to-Many sempre precisa de tabela intermediária
	✔ Essa tabela guarda apenas chaves estrangeiras
	✔ O GORM gerencia isso automaticamente
	✔ Preload evita múltiplas queries manuais
	✔ gorm.Model adiciona created_at, updated_at, deleted_at

	==========================================================================

	🏛 🔟 Visão arquitetural

	Esse padrão é usado em:

	- Produtos ↔ Categorias
	- Usuários ↔ Permissões
	- Alunos ↔ Cursos
	- Filmes ↔ Atores
	- Usuários ↔ Grupos

	Sempre que:
	Muitos se relacionam com muitos.

	==========================================================================

	🧠 RESUMO FINAL

	Structs definem o modelo.
	GORM transforma em tabelas.
	Many-to-Many cria tabela de junção.
	Essa tabela guarda apenas IDs.
	Preload carrega relacionamento automaticamente.

	Você modela em Go.
	O GORM traduz para SQL.
	O banco mantém a integridade via Foreign Keys.

	Isso é ORM mapeando objeto → relacional.
*/
