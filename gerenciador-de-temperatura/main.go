package main

import (
	"log/slog"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	numberChan := make(chan int64)

	chanTratador1 := make(chan int64)
	chanTratador2 := make(chan int64)

	var wg sync.WaitGroup

	wg.Add(1)
	go numberGenerator(numberChan, &wg)

	wg.Add(1)
	go watcher(numberChan, chanTratador1, chanTratador2, &wg)

	wg.Add(2)
	go keeper("tratador1", chanTratador1, &wg)
	go keeper("tratador2", chanTratador2, &wg)

	wg.Wait()

}

func numberGenerator(numberChan chan int64, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		number := rand.Int64N(100)
		slog.Info("Gerador: Enviando número:", "number", number)

		numberChan <- number

		time.Sleep(500 * time.Millisecond)
	}
}

func watcher(numberChan, ch1, ch2 chan int64, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range numberChan {
		if num <= 40 {
			slog.Info("Observer: Sending to keeper1\n", "number", num)
			ch1 <- num
		} else if num > 50 {
			slog.Info("Observer: Sending  to keeper2\n", "number", num)
			ch2 <- num
		}
	}
}

func keeper(name string, ch chan int64, wg *sync.WaitGroup) {
	defer wg.Done()

	counter := 0

	for num := range ch {
		slog.Info("Processando número", "name", name, "number", num)
		counter++

		if counter == 10 {
			slog.Info("** Atingiu 10 processamentos! **", "name", name)
			counter = 0
		}

		time.Sleep(1 * time.Second)
	}
}
