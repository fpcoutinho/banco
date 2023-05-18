package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {

	cliente := clientes.Titular{}
	cliente.New("Filipe", "123.123.123-12", "Desenvolvedor")
	conta := contas.ContaCorrente{}
	conta.New(&cliente, 93345, 456)
	conta2 := contas.ContaCorrente{}
	conta2.New(&cliente, 88642, 456)

	conta.Depositar(500)
	conta.Transferir(130, &conta2)

	fmt.Println(conta.ToString())
	fmt.Println(conta2.ToString())

	PagarBoleto(&conta, 60)
	fmt.Println(conta.ToString())
	fmt.Println(conta2.ToString())
}

func PagarBoleto(c contas.Conta, valorBoleto float64) {
	c.Sacar(valorBoleto)
}
