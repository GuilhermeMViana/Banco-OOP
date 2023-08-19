package main

import (
	"fmt"
	"os"
	"time"

	"banco/clientes"
	"banco/contas"
)

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	conta1 := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Guilherme",
		CPF:       "546.241.138-37",
		Profissao: "Front-End Developer"},
		NumeroAgencia: 123, NumeroDaConta: 123456}

	conta2 := contas.ContaPoupanca{Titular: clientes.Titular{
		Nome:      "André",
		CPF:       "243.277.777-81",
		Profissao: "Software Engineer"},
		NumeroAgencia: 132, NumeroDaConta: 654321}

	selecionarConta()
	tipoDaConta := receiveCommand()

	switch tipoDaConta {
	case 1:
		for {
			mostrarMenu()
			command := receiveCommand()

			switch command {
			case 1:
				var saque float64
				fmt.Println("Digite o valor do saque:")
				_, err := fmt.Scan(&saque)

				if err != nil {
					fmt.Println("Erro no saque:", err)
				}

				status := conta1.Sacar(saque)

				fmt.Println("---------------------------")
				fmt.Println(status)
				fmt.Println("Novo saldo:")
				fmt.Println(conta1.ObterSaldo())
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

			case 4:
				var valor float64
				fmt.Println("Digite o valor do boleto:")
				_, err := fmt.Scan(&valor)

				if err != nil {
					fmt.Println("Erro ao fazer pagamento:", err)
				}

				PagarBoleto(&conta1, valor)

				fmt.Println("---------------------------")
				fmt.Println("Novo saldo:")
				fmt.Println(conta1.ObterSaldo())
				fmt.Println("---------------------------")

			case 0:
				os.Exit(0)

			default:
				os.Exit(-1)
			}
		}

	case 2:
		for {
			mostrarMenu()
			command := receiveCommand()

			switch command {
			case 1:
				var saque float64
				fmt.Println("Digite o valor do saque:")
				_, err := fmt.Scan(&saque)

				if err != nil {
					fmt.Println("Erro no saque:", err)
				}

				status := conta2.Sacar(saque)

				fmt.Println("---------------------------")
				fmt.Println(status)
				fmt.Println("Novo saldo:")
				fmt.Println(conta2.ObterSaldo())
				fmt.Println("---------------------------")

				restartOptions()

			case 2:
				var deposito float64
				fmt.Println("Digite o valor do depósito: ")
				_, err := fmt.Scan(&deposito)

				if err != nil {
					fmt.Println("Erro no deposito:", err)
				}

				status, novoValor := conta2.Depositar(deposito)

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
				fmt.Println(conta2)
				fmt.Println(conta1)

				restartOptions()

			case 4:
				var valor float64
				fmt.Println("Digite o valor do boleto:")
				_, err := fmt.Scan(&valor)

				if err != nil {
					fmt.Println("Erro ao fazer pagamento:", err)
				}

				PagarBoleto(&conta2, valor)

				fmt.Println("---------------------------")
				fmt.Println("Novo saldo:")
				fmt.Println(conta2.ObterSaldo())
				fmt.Println("---------------------------")

			case 0:
				os.Exit(0)

			default:
				os.Exit(-1)
			}
		}

	case 0:
		os.Exit(0)
	default:
		os.Exit(-1)
	}

}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func selecionarConta() {
	fmt.Println("Selecione o tipe da conta")
	fmt.Println("---------------------------")
	fmt.Println("1 - Conta Corrente")
	fmt.Println("2 - Conta Poupança")
	fmt.Println("0 - Finalizar programa")
	fmt.Println("---------------------------")
}

func mostrarMenu() {
	fmt.Println("Selecione uma opção")
	fmt.Println("---------------------------")
	fmt.Println("1 - Realizar saque")
	fmt.Println("2 - Realizar Deposito")
	fmt.Println("3 - Realizar Transferência")
	fmt.Println("4 - Pagar Boleto")
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
	mostrarMenu()
}
