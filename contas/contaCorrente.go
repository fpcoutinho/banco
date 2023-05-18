package contas

import (
	"banco/clientes"
	"fmt"
)

type ContaCorrente struct {
	titular *clientes.Titular
	saldo   float64
	numero  int
	agencia int
}

func (c *ContaCorrente) New(titular *clientes.Titular, numero int, agencia int) {
	c.titular = titular
	c.saldo = 0
	c.numero = numero
	c.agencia = agencia
}

func (c *ContaCorrente) Saldo() float64 {
	return c.saldo
}

func (c *ContaCorrente) Titular() *clientes.Titular {
	return c.titular
}

func (c *ContaCorrente) Numero() int {
	return c.numero
}

func (c *ContaCorrente) Agencia() int {
	return c.agencia
}

func (c *ContaCorrente) Depositar(valor float64) bool {
	if valor <= 0 {
		return false
	}
	c.saldo += valor
	return true
}

func (c *ContaCorrente) Sacar(valor float64) bool {
	if c.saldo >= valor && valor > 0 {
		c.saldo -= valor
		return true
	}
	return false
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	if valor <= 0 || c.saldo < valor {
		return false
	}
	c.saldo -= valor
	contaDestino.Depositar(valor)
	return true
}

func (c *ContaCorrente) ToString() string {
	return fmt.Sprintf("Titular: %s\nAgência: %d - Conta: %d\nSaldo: %0.2f\n", c.titular.Nome(), c.agencia, c.numero, c.saldo)
}
