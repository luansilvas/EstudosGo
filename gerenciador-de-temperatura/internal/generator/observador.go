package generator

import (
	"gerenciador-de-temperatura/internal/config"
	"math/rand"
	"sync"
	"time"
)

func Start(condTrat1, condTrat2 *sync.Cond, cfg config.Config, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		number := rand.Int63n(100)

		if number <= 40 {
			condTrat1.Signal()
		} else if number > 50 {
			condTrat2.Signal()
		}

		time.Sleep(time.Duration(cfg.GeneratorDelay) * time.Millisecond)
	}
}
