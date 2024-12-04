package transactions

import (
	"autorizador-debito/internal/config"
	"context"
	"fmt"
	"sync"
	"time"
)

type TransactionManager struct {
	data   sync.Map
	config *config.Config
}

func NewTransactionManager() *TransactionManager {
	cfg := config.LoadConfig()
	return &TransactionManager{
		data:   sync.Map{},
		config: cfg,
	}
}

func (tm *TransactionManager) ProcessTransaction(ctx context.Context, userId string, value float64) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("operation canceled or timed out: %w", ctx.Err())
	default:
	}

	transactionsByUser := tm.getTransactions(userId)
	available, err := CalculateAvailable(transactionsByUser, tm.config.TransactionLimit, tm.config.TransactionPeriod)
	if err != nil {
		return err
	}

	if available < value {
		return fmt.Errorf("insufficient limit")
	}

	tm.appendTransaction(userId, Transaction{
		UserId:   userId,
		DateTime: time.Now(),
		Value:    value,
	})

	return nil
}

func (tm *TransactionManager) getTransactions(userId string) []Transaction {
	if existing, exists := tm.data.Load(userId); exists {
		return existing.([]Transaction)
	}
	return []Transaction{}
}

func (tm *TransactionManager) appendTransaction(userId string, transaction Transaction) {
	for {
		existing, _ := tm.data.Load(userId)

		var transactions []Transaction
		if existing != nil {
			transactions = append(existing.([]Transaction), transaction)
		} else {
			transactions = []Transaction{transaction}
		}

		tm.data.Store(userId, transactions)
		return
	}
}
