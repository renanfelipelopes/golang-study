package main

/*

Exemplo 1
Forma de alterar valor apenas passando "cópia" do valor da variável.
'Bug silencioso', pois não altera o valor original da var a;

func alteraValor(x int) {
	x = 20
}

func main() {
	a := 10
	alteraValor(a)
	println(a) // 10
}
*/

/*

Exemplo 2: forma certa de alterar o valor, usando ponteiro evitamos de passar a cópia
do valor original, e alteramos direto na memória.

type Valor struct {
	saldo int
}

func (v *Valor) alteraValor() {
	v.saldo = 20
}

func main() {
	a := Valor{saldo: 10}
	a.alteraValor()
	println(a.saldo) // 20
}
*/

/*
Exemplo 3
*/

type Conta struct {
	saldo int
}

func (c *Conta) Simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

// Ponteiros e Structs aula F025
func main() {
	conta := Conta{saldo: 100}
	conta.Simular(200)
	println(conta.saldo)
}

// Agora sim o valor original está sendo alterado para 300 no saldo, pois se repararmos,
// no metodo Simular, agora recebe como parametro a Struct Conta com um asteristico, o que
// significa que eu estou passando o ponteiro e manipulando o valor original contido dentro
// da memoria, e não uma copia.

// É comum em Go criar Structs com apontamentos para a memoria porque geralmente existe a
// intenção de alterar o valor original.
