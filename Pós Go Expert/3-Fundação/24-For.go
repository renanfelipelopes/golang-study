package main

/*
	Aula: Loops (for / range)

	Go não possui:
	- while
	- do while
	- foreach

	Go possui apenas:
	- for
	- for + range (para iterar coleções)

	Obs:
	map NÃO é laço.
	map é uma estrutura de dados.
	O laço continua sendo o for, usando range para iterar.
*/

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "três"}
	for k, _ := range numeros {
		println(k)
	}

	// podemos usar o for parecido como um while
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// temos o loop infinito que tbm é comum em Go
	// porque é usado? imagine que vc queira consumir mensagens de uma fila,
	// vc pode criar um metodo operador que fique sempre escutando e executando o que vier da fila
	for {
		print("Hello, World!")
	}
}

/*

🧠 Loop tradicional (estilo C / Java)
for i := 0; i < 10; i++ {
	println(i)
}


🔥 Foreach do Go → for + range
numeros := []string{"um", "dois", "três"}

for k, v := range numeros {
	println(k, v)
}

🔍 O que range retorna?
Estrutura	Primeiro valor	Segundo valor
slice		índice			valor
array		índice			valor
map			chave			valor
string		índice			rune
channel		valor			-

--------------------------------------------------

🧪 Ignorando valores
Você faz isso:
for k, _ := range numeros

Forma mais idiomática:
for k := range numeros

--------------------------------------------------

🧠 Range com Map
m := map[string]int{
	"Renan": 1000,
	"Maria": 2000,
}

for k, v := range m {
	println(k, v)
}

--------------------------------------------------

🧠 Range com String
s := "Renan"

for i, r := range s {
	println(i, r)
}

👉 r é rune (Unicode)

--------------------------------------------------

🔥 For como while
i := 0

for i < 10 {
	println(i)
	i++
}

--------------------------------------------------

🧠 Uso real no mercado
Worker Queue
for {
	msg := <-fila
	processa(msg)
}

Servidor escutando eventos
for {
	conn := aceitaConexao()
	go handle(conn)
}

--------------------------------------------------

⚠️ Loop infinito sem controle = problema

Boa prática:

for {
	select {
	case msg := <-fila:
		processa(msg)
	case <-ctx.Done():
		return
	}
}

👉 Muito usado com goroutines.
*/
