package processor

import (
	"log/slog"
	"strings"
	"sync"
	"time"
)

func Start(name string, cond *sync.Cond, log *slog.Logger, wg *sync.WaitGroup) {
	defer wg.Done()

	counter := 0

	for {
		cond.L.Lock()
		cond.Wait()
		cond.L.Unlock()

		log.Debug("Processando número", "name", name)
		counter++

		if counter == 10 {
			var msg strings.Builder
			msg.WriteString("O tratador ")
			msg.WriteString(name)
			msg.WriteString(" atingiu 10 processamentos!")

			log.Info(msg.String())
			counter = 0
		}

		time.Sleep(1 * time.Second)
	}
}
