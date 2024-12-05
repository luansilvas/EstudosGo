package ui

import (
	"autorizador-debito/internal/transactions"
	"fmt"
	"log/slog"
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

func ProcessTransaction(tm *transactions.TransactionManager) {
	fmt.Println("Digite o ID do usuário que deseja operar:")

	var userId string
	_, _ = fmt.Scan(&userId)
	slog.Debug("Id de usuário recebido com sucesso", "userId", userId)

	fmt.Println("Digite o valor que deseja transferir:")
	var value float64
	_, _ = fmt.Scan(&value)
	slog.Debug("Valor de transferencia recebido", "value", value)

	err := tm.ProcessTransaction(userId, value)
	if err != nil {
		fmt.Printf("Erro ao processar transação. erro: %v\n", err)
	} else {
		fmt.Println("Transação bem-sucedida!")
	}
}
