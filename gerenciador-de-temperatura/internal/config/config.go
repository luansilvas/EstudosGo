package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	LogLevel       string
	GeneratorDelay int
}

func LoadConfig(filepath string) Config {
	err := godotenv.Load(filepath)
	if err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %v", err)
	}

	delay, _ := strconv.Atoi(os.Getenv("GENERATOR_DELAY"))
	return Config{
		LogLevel:       os.Getenv("LOG_LEVEL"),
		GeneratorDelay: delay,
	}
}
