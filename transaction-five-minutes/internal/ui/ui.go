package ui

import (
	"autorizador-debito/internal/transactions"
	"context"
	"fmt"
	"log/slog"
	"time"
)

func ShowIntroduction() {
	fmt.Println("Welcome to the debit authorizer!")
}

func ShowMenu() {
	fmt.Println("2 - Transfer money")
	fmt.Println("0 - Exit")
}

func ReadCommand() int {
	var command int
	_, _ = fmt.Scan(&command)
	slog.Debug("Command selected", "command", command)
	return command
}

func ProcessTransaction(ctx context.Context, tm *transactions.TransactionManager) {
	fmt.Println("Please type the userId that wants to authorize")

	var userId string
	_, _ = fmt.Scan(&userId)
	slog.Debug("User id successfully typed", "userId", userId)

	fmt.Println("Now tell how much money you want to transfer")
	var value float64
	_, _ = fmt.Scan(&value)
	slog.Debug("Value successfully sent", "value", value)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := tm.ProcessTransaction(ctx, userId, value)
	if err != nil {
		fmt.Printf("Error processing transaction. error: %v\n", err)
	} else {
		fmt.Println("Transaction successful.")
	}
}
