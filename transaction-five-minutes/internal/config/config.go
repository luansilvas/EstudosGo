package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	TransactionLimit  float64
	TransactionPeriod int
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("It was not possible to load .env file: %v", err)
	}

	limitStr := os.Getenv("TRANSACTION_LIMIT")
	limit, err := strconv.ParseFloat(limitStr, 64)
	if err != nil || limit <= 0 {
		log.Fatalf("the variable TRANSACTION_LIMIT is not valid: %v", err)
	}

	periodStr := os.Getenv("TRANSACTION_PERIOD_MINUTES")
	period, err := strconv.Atoi(periodStr)
	if err != nil || period <= 0 {
		log.Fatalf("the variable TRANSACTION_PERIOD_MINUTES is not valid: %v", err)
	}

	return &Config{limit, period}
}
