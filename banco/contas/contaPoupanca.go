package contas

import (
	"fmt"
	. "golang-studies/banco/clientes"
)

type ContaPoupanca struct {
	Titular Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo float64
}


func (c *ContaPoupanca) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque de: " + fmt.Sprintf("%.2f", valorSaque) + " efetuado com sucesso!"
	}
	return "saldo insuficiente"
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Você depositou: " + fmt.Sprintf("%.2f", valorDoDeposito), c.saldo
	}
	return "Valor para deposito deve ser maior que Zero.", c.saldo
}

func (c ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}