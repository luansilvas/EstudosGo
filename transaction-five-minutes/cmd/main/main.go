package main

import (
	"autorizador-debito/internal/transactions"
	"autorizador-debito/internal/ui"
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	transactionManager := transactions.NewTransactionManager()

	ui.ShowIntroduction()
	for {
		ui.ShowMenu()

		command := ui.ReadCommand()
		switch command {
		case 2:
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			ui.ProcessTransaction(ctx, transactionManager)
			cancel()
		case 0:
			fmt.Println("exiting...")
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
			os.Exit(-1)
		}
	}
}
