package main

import (
	"banco/contas"
	"fmt"
	"os"
	"time"
)

func main() {
	conta1 := contas.ContaCorrente{Titular: "Guilherme", NumeroAgencia: 123, NumeroDaConta: 123456, Saldo: 500}
	conta2 := contas.ContaCorrente{Titular: "Gabriel", NumeroAgencia: 132, NumeroDaConta: 654321, Saldo: 250}
	for {
		showMenu()

		command := receiveCommand()

		switch command {
		case 1:
			var saque float64
			fmt.Println("Digite o valor do saque:")
			_, err := fmt.Scan(&saque)

			if err != nil {
				fmt.Println("Erro no saque:", err)
			}

			status, novoValor := conta1.Sacar(saque)

			fmt.Println("---------------------------")
			fmt.Println(status)
			fmt.Println("Novo saldo:", novoValor)
			fmt.Println("---------------------------")

			restartOptions()

		case 2:
			var deposito float64
			fmt.Println("Digite o valor do depósito: ")
			_, err := fmt.Scan(&deposito)

			if err != nil {
				fmt.Println("Erro no deposito:", err)
			}

			status, novoValor := conta1.Depositar(deposito)

			fmt.Println("---------------------------")
			fmt.Println(status)
			fmt.Println("Novo saldo:", novoValor)
			fmt.Println("---------------------------")

			restartOptions()

		case 3:
			var transferencia float64
			fmt.Println("Digite o valor da transferência:")
			_, err := fmt.Scan(&transferencia)

			if err != nil {
				fmt.Println("Erro com a transferência:", err)
			}

			conta2.Transfere(transferencia, &conta1)
			fmt.Println(conta1)
			fmt.Println(conta2)

			restartOptions()

		case 0:
			os.Exit(0)

		default:
			os.Exit(-1)
		}
	}

}

func showMenu() {
	fmt.Println("Selecione uma opção")
	fmt.Println("---------------------------")
	fmt.Println("1 - Realizar saque")
	fmt.Println("2 - Realizar Deposito")
	fmt.Println("3 - Realizar Transferência")
	fmt.Println("0 - Finalizar programa")
	fmt.Println("---------------------------")

}

func receiveCommand() int {
	var command int

	fmt.Println("Escolha uma opção: ")
	_, err := fmt.Scan(&command)

	if err != nil {
		fmt.Println(err)
	}

	return command
}

func restartOptions() {
	for i := 3; i > 0; i-- {
		fmt.Println("Voltando para o menu em", i, "segundos")
		time.Sleep(1 * time.Second)
	}
	showMenu()
}
