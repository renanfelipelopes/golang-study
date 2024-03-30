package main

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
