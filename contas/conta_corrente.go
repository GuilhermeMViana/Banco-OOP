package contas

import (
	"banco/clientes"
	"fmt"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroDaConta int
	saldo         float64
}

func (conta *ContaPoupanca) Sacar(saque float64) string {

	if saque <= conta.saldo && saque > 0 {
		conta.saldo -= saque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente"

	}

}

func (conta *ContaPoupanca) Depositar(deposito float64) (string, float64) {
	if deposito > 0 {
		conta.saldo += deposito
		return "Depósito realizado com sucesso!", conta.saldo
	} else {
		return "Valor ínvalido!", conta.saldo
	}

}

func (conta *ContaPoupanca) Transfere(transferencia float64, contaDestino *ContaCorrente) bool {

	if conta.saldo > transferencia && transferencia > 0 {
		conta.saldo -= transferencia
		contaDestino.Depositar(transferencia)
		return true
	} else {
		fmt.Println("Saldo insuficiente!")
		return false
	}
}

func (conta *ContaPoupanca) ObterSaldo() (string, float64) {
	return "Seu saldo atual é de:", conta.saldo
}
