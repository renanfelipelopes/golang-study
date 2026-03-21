package main

/*
	Aula: Condicionais

	Go possui:
	- if
	- else if (SIM, existe em Go)
	- else
	- switch
	- switch sem expressão (switch true implícito)

	Go NÃO possui:
	- elif (Python)
	- ternário (condicao ? x : y)

	Obs:
	Go força uso de chaves {}
	Não existe if sem {}
*/

func main() {
	a := 1
	b := 2
	c := 3

	if a > b {
		println(a)
	} else {
		println(b)
	}

	// ✅ Existe sim else if
	if a > b {
		println("a")
	} else if b > c {
		println("b")
	} else {
		println("c")
	}

	// ------------------------------------------------
	if a > b || c < b && b == a {
		println("Amostra de operadores logicos")
	}

	// temos tbm o switch case
	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default:
		println("default")
	}

}

/*
🔥 If com variável dentro (muito Go idiomático)
if err := salvar(); err != nil {
	println("erro")
}

------------------------------------------------

🧠 Diferença para Java / C#

👉 Go NÃO tem fallthrough automático.

✅ Só se usar fallthrough
case 1:
	println("1")
	fallthrough
case 2:
	println("2")

"Fallthrough" no contexto de programação (especialmente em estruturas switch-case de linguagens como C, C++, Java e JavaScript) ocorre quando a execução de um caso (case) passa automaticamente para o próximo, sem interromper com um comando break. Isso permite que múltiplos casos executem o mesmo código, mas, se não intencional, pode causar bugs.
Principais detalhes sobre Fallthrough:
Funcionamento: Quando um case corresponde ao valor, mas não tem um break no final, o programa "cai" para o próximo case e executa o código dele também.
Intencional vs. Acidental: Pode ser usado para agrupar casos que executam a mesma ação. No entanto, é frequentemente considerado um "erro de programação" se o break foi esquecido acidentalmente, tornando o código difícil de manter.
Em Go (Golang): Diferente de outras linguagens, o Go não tem fallthrough automático. Ele utiliza a palavra-chave explícita fallthrough para transferir o controle para o próximo case.

------------------------------------------------

🔥 Switch com múltiplos valores
switch a {
case 1, 2, 3:
	println("baixo")
case 4, 5, 6:
	println("medio")
}

------------------------------------------------

🔥 Type Switch
switch v := x.(type) {
case int:
	println("int")
case string:
	println("string")
}

------------------------------------------------

🧠 Go NÃO tem operador ternário

❌ Não existe:
x := cond ? a : b

✅ Forma Go
var x int

if cond {
	x = a
} else {
	x = b
}

------------------------------------------------

🧠 Boa prática Go

Preferir:

if err != nil {
	return err
}


Evitar:

if err == nil {
	// codigo gigante
}
*/
