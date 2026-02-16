package main

// func alteraValor(x int) {
// 	x = 20
// }

// func main() {
// 	a := 10
// 	alteraValor(a)
// 	println(a) // 10
// }

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
