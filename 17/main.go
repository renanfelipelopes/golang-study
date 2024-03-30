package main

type Conta struct {
	saldo int
}

func (c Conta) Simular(valor int) int {
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

// A forma como está funcionando o valor da variavel saldo, está errado, pois na funcao main quando
// damos um print no valor do saldo, ele exibe 100, mas ao verificar o valor no print dentro do metodo
// Simular(), o valor é de 300, ou seja, a soma de 100 (saldo inicial) + 200 (valor inserido como parametro)
// isso ocorre porque o valor do atributo saldo na funcao main, apenas é passado uma copia dele para a funcao Simular
// o que faz com que a soma não altere o valor original. Para alterar o valor original, precisamos passar o endereco da
// memoria e inserir o valor da soma lá dentro, alterando o valor original para 300.
// O exemplo está na pasta 18
