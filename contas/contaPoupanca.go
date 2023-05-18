package contas

import (
	"banco/clientes"
	"fmt"
)

type ContaPoupanca struct {
	titular  *clientes.Titular
	saldo    float64
	numero   int
	agencia  int
	operacao int
}

// Construtor
func (c *ContaPoupanca) New(titular *clientes.Titular, numero int, agencia int, operacao int) {
	c.titular = titular
	c.saldo = 0
	c.numero = numero
	c.agencia = agencia
	c.operacao = operacao
}

// Getters
func (c *ContaPoupanca) Saldo() float64 {
	return c.saldo
}

func (c *ContaPoupanca) Titular() clientes.Titular {
	return *c.titular
}

func (c *ContaPoupanca) Numero() int {
	return c.numero
}

func (c *ContaPoupanca) Agencia() int {
	return c.agencia
}

// Métodos
func (c *ContaPoupanca) Depositar(valor float64) bool {
	if valor <= 0 {
		return false
	}
	c.saldo += valor
	return true
}

func (c *ContaPoupanca) Sacar(valor float64) bool {
	if c.saldo >= valor && valor > 0 {
		c.saldo -= valor
		return true
	}
	return false
}

func (c *ContaPoupanca) ToString() string {
	return fmt.Sprintf("Titular: %s\nAgência: %d - Conta: %d\nSaldo: %0.2f\n", c.titular.Nome(), c.agencia, c.numero, c.saldo)
}
