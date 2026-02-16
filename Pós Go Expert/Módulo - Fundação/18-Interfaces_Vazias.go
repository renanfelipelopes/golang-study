package main

import "fmt"

func main() {
	var minhaVar interface{} = "Renan Lopes"
	println(minhaVar.(string))

	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)
}

/*
Quando se usa interfaces vazias, no final das contas, o Go vai querer saber qual é o tipo
daquela variavel que vc está usando. E podemos forçar aquele tipo a ser transformado, que é
chamado de Type assertation.
O metodo da forma a seguir, o Go irá imprimir um código (0xc27040,0xc4ca90) ao inves do texto:
	func main() {
		var minhaVar interface{} = "Renan Lopes"
		println(minhaVar)
	}
Isso ocorre pq o Go não está entendendo o tipo para imprimir.
Agora quando 'forço' a tipagem na saida, ficaria assim meu print:
	println(minhaVar.(string))
E isso é chamado de Type assertation. Porém, nos casos numericos isso pode gerar problemas,
exemplo:
	func main() {
		var minhaVar interface{} = "Renan Lopes"
		println(minhaVar.(string))
		res, ok := minhaVar.(int)
		fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)
	}
O ok no caso irá validar se a conversão do tipo deu certo, mas o print mostra que não, dará
um false como resultado.

Se caso eu tentasse criar a conversao sem a variavel ok por exemplo, o Go daria um erro de
Panic, isso pq ele percebe que vc não esta querendo verificar se deu certo, e ele tenta manter
a tipagem forte.

Isso é mais comum em codigos legados, quando o Go não possuia tipos Genericos, e se usava muito
interfaces vazias.
*/
