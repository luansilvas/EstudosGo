package main

import (
	"gerenciador-de-temperatura/internal/config"
	"gerenciador-de-temperatura/internal/generator"
	"gerenciador-de-temperatura/internal/logger"
	"gerenciador-de-temperatura/internal/processor"
	"sync"
)

func main() {
	cfg := config.LoadConfig(".env")

	log := logger.New(cfg.LogLevel)

	var (
		mu        sync.Mutex
		condTrat1 = sync.NewCond(&mu)
		condTrat2 = sync.NewCond(&mu)
	)

	var wg sync.WaitGroup

	wg.Add(1)
	go generator.Start(condTrat1, condTrat2, cfg, &wg)

	wg.Add(2)
	go processor.Start("tratador1", condTrat1, log, &wg)
	go processor.Start("tratador2", condTrat2, log, &wg)

	wg.Wait()
}
