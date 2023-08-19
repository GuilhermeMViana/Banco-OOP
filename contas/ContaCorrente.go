package contas

import (
	"banco/Clientes"
	"fmt"
)

type ContaCorrente struct {
	Titular       Clientes.Titular
	NumeroAgencia int
	NumeroDaConta int
	Saldo         float64
}

func (conta *ContaCorrente) Sacar(saque float64) (string, float64) {

	if saque <= conta.Saldo && saque > 0 {
		conta.Saldo -= saque
		return "Saque realizado com sucesso!", conta.Saldo
	} else {
		return "Saldo insuficiente", conta.Saldo

	}

}

func (conta *ContaCorrente) Depositar(deposito float64) (string, float64) {
	if deposito > 0 {
		conta.Saldo += deposito
		return "Depósito realizado com sucesso!", conta.Saldo
	} else {
		return "Valor ínvalido!", conta.Saldo
	}

}

func (conta *ContaCorrente) Transfere(transferencia float64, contaDestino *ContaCorrente) bool {

	if conta.Saldo > transferencia && transferencia > 0 {
		conta.Saldo -= transferencia
		contaDestino.Depositar(transferencia)
		return true
	} else {
		fmt.Println("Saldo insuficiente!")
		return false
	}
}
