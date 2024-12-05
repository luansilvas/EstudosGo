package main

import (
	"autorizador-debito/internal/transactions"
	"autorizador-debito/internal/ui"
	"fmt"
	"os"
)

func main() {
	transactionManager := transactions.NewTransactionManager()

	ui.ShowIntroduction()
	for {
		ui.ShowMenu()

		command := ui.ReadCommand()
		switch command {
		case 2:
			ui.ProcessTransaction(transactionManager)
		case 0:
			fmt.Println("exiting...")
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
			os.Exit(-1)
		}
	}
}
