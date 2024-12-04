package transactions

import (
	"time"
)

type Transaction struct {
	UserId   string
	DateTime time.Time
	Value    float64
}

func CalculateAvailable(transactions []Transaction, limit float64, periodMinutes int) (float64, error) {
	periodAgo := time.Now().Add(-time.Duration(periodMinutes) * time.Minute)

	var totalTransferred float64
	for _, t := range transactions {
		if t.DateTime.After(periodAgo) {
			totalTransferred += t.Value
		}
	}

	available := limit - totalTransferred

	return available, nil
}
