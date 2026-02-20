/*
	Aula: Trabalhando com Jsons
	Marshal é converter em json

	Podemos começar por converter uma struct em um json
	Json sempre vai retornar em bytes

	Quando uso o Marshal, eu salvo para mim o valor do json numa variavel
	Quando uso o Encoder, eu pego o valor, já faço o processo de serialização entregando
	para alguém (o Stdout por exemplo, ou um arquivo, ou webserver)
*/

package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int
	Saldo  int
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}
}
