package contas

import (
	"fmt"
	. "golang-studies/banco/clientes"
)

type ContaCorrente struct {
	Titular       Titular
	NumeroAgencia, NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque de: " + fmt.Sprintf("%.2f", valorSaque) + " efetuado com sucesso!"
	}
	return "saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Você depositou: " + fmt.Sprintf("%.2f", valorDoDeposito), c.saldo
	}
	return "Valor para deposito deve ser maior que Zero.", c.saldo
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) string {
	if valorDaTransferencia > 0 && valorDaTransferencia <= c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return c.Titular.Nome + " você transferiu: " + fmt.Sprintf("%.2f", valorDaTransferencia) + " para a conta: " + contaDestino.
			Titular.Nome + " Agencia: " + fmt.Sprint(contaDestino.NumeroAgencia) + " Conta: " + fmt.Sprint(contaDestino.NumeroConta)
	}

	return "Não foi possível realizar a transferência, verifique seu saldo e o valor informado!"
}

func (c ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}