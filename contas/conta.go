package contas

import "banco/clientes"

type Conta interface {
	Titular() *clientes.Titular
	Numero() int
	Agencia() int
	Saldo() float64
	Depositar(valor float64) bool
	Sacar(valor float64) bool
	ToString() string
}
