package main

import (
	"fmt"
	. "golang-studies/banco/clientes"
	. "golang-studies/banco/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64)  {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valorDoSaque float64) string
}

func main() {

	titularArnaud := Titular{ Nome: "Arnaud Santana Alves", CPF: "589.666.220-38"}
	contaDoArnaud := ContaCorrente{ Titular: titularArnaud, NumeroAgencia: 4170, NumeroConta: 24899}
	contaDoArnaud.Depositar(777.77)

	fmt.Println(contaDoArnaud)
	fmt.Println(contaDoArnaud.Sacar(300.00))
	fmt.Println(contaDoArnaud)

	status, saldo:= contaDoArnaud.Depositar(100)
	fmt.Println(status, "seu saldo atual é de: ", saldo)

	titularLilia := Titular{ Nome: "Lilia Maria dos Santos"}
	contaDaLilia := ContaCorrente{ Titular: titularLilia, NumeroAgencia: 4170, NumeroConta: 24777}
	contaDaLilia.Depositar(250.77)

	fmt.Println(contaDaLilia)

	fmt.Println(contaDoArnaud.Transferir(50.00, &contaDaLilia))
	fmt.Println(contaDaLilia)
	fmt.Println(contaDoArnaud)
	PagarBoleto(&contaDoArnaud, 30)

	contaDaAna := ContaCorrente{}
	contaDaAna.Depositar(-100)

	fmt.Println(contaDaLilia.Titular.Nome, "seu saldo é:", contaDaLilia.ObterSaldo())
	fmt.Println(contaDoArnaud.Titular.Nome, "seu saldo é:", contaDoArnaud.ObterSaldo())
	fmt.Println(contaDaAna.Titular.Nome, "seu saldo é:", contaDaAna.ObterSaldo())

	titularGabriel := Titular{Nome: "Gabriel dos Santos Alves", CPF: "56678909899", Profissao: "Biologo"}
	contaDoGabriel := ContaPoupanca{ Titular: titularGabriel, NumeroConta: 45677, NumeroAgencia: 4170, Operacao: 157}
	contaDoGabriel.Depositar(757.00)
	fmt.Println(contaDoGabriel.ObterSaldo())
	PagarBoleto(&contaDoGabriel, 100)
	fmt.Println(contaDoGabriel.ObterSaldo())
}
