package main

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"strconv"
	"time"
)

const transactionLimit5m float64 = 1000

type Transaction struct {
	UserId   string
	DateTime time.Time
	Value    float64
}

func main() {
	transactions := make(map[string][]Transaction)

	showIntroduction()
	for {
		showMenu()

		command := readCommand()
		switch command {
		case 1:
			generateUserId()
		case 2:
			processTransaction(transactions)
		case 0:
			fmt.Println("exiting")
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
			os.Exit(-1)
		}
	}

}

func processTransaction(transactions map[string][]Transaction) {
	fmt.Println("Please type the userId that wants to authorize")

	var userId string
	_, _ = fmt.Scan(&userId)
	fmt.Println("The userId typed was ", userId)

	fmt.Println("Now tell how much money you want to transfer")
	var value float64
	_, _ = fmt.Scan(&value)
	fmt.Println("The value you are sending is", value)

	checkValues(transactions, userId, value)
}

func checkValues(transactions map[string][]Transaction, id string, value float64) {
	available := transactionLimit5m
	transferredByUser, exists := transactions[id]

	if !exists {
		if value <= available {
			transaction := Transaction{DateTime: time.Now(), UserId: id, Value: value}
			transactions[id] = []Transaction{transaction}
			fmt.Println("First transaction for user: ", strconv.FormatFloat(value, 'f', 2, 64))
		} else {
			fmt.Println("Transaction exceeds the available limit for new user.")
		}
		return
	}

	available -= checkSentValuesInFiveMinutes(transferredByUser)

	if available >= value {
		transaction := Transaction{DateTime: time.Now(), UserId: id, Value: value}
		transactions[id] = append(transferredByUser, transaction)
		fmt.Println("The user has a limit! Transferring.")
	} else {
		fmt.Println("The user has no limit. Available limit: ", strconv.FormatFloat(available, 'f', 2, 64))
	}
}

func checkSentValuesInFiveMinutes(transactions []Transaction) float64 {
	var transferred float64 = 0
	fiveMinutesAgo := time.Now().Add(-5 * time.Minute)

	for _, transaction := range transactions {
		if transaction.DateTime.After(fiveMinutesAgo) {
			transferred += transaction.Value
		}
	}

	fmt.Println("Value sent by user the last five minutes " + strconv.FormatFloat(transferred, 'f', 2, 64))

	return transferred
}

func showIntroduction() {
	fmt.Println("Welcome to the debit authorizer!")
}

func showMenu() {
	fmt.Println("1 - Register / generate user-id")
	fmt.Println("2 - Transfer money")
	fmt.Println("0 - exit")
}

func readCommand() int {
	var command int
	_, _ = fmt.Scan(&command)
	fmt.Println("The chosen command was ", command)

	return command
}

func generateUserId() {
	userId, _ := uuid.NewUUID()

	fmt.Println("The user id is: ", userId.String())
}
