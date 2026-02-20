/*
	Aula: Trabalhando com Jsons
	Marshal é converter em json

	Podemos começar por converter uma struct em um json
	Json sempre vai retornar em bytes

	Quando uso o Marshal, eu salvo para mim o valor do json numa variavel
	Quando uso o Encoder, eu pego o valor, já faço o processo de serialização entregando
	para alguém (o Stdout por exemplo, ou um arquivo, ou webserver)

	Se eu colocar o valor do json diferente do nome dos campos da struct, irá retornar o
	valor zero, pois o go não terá conseguido fazer o bind dos campos
	O go tem um recurso muito interessante, que é o TAGS, por exemplo:
	type Conta struct {
		Numero int `json:"n"`
		Saldo  int `json:"s"`
	}
	jsonPuro := []byte(`{"s":2,"s":200}`)
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

	jsonPuro := []byte(`{"Numero":2,"Saldo":200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	println(contaX.Saldo)
}
