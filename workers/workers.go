package workers

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Work() {
	channel := make(chan int)

	// Criando os workers
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(channel)
	}

	// Adicionando valor ao channel
	for i := 0; i <= 100; i++ {
		channel <- i
	}

	close(channel)
	wg.Wait()
	fmt.Println("Fim...")
}

func worker(channel chan int) {
	defer wg.Done()
	for i := range channel {
		fmt.Println(i)
		time.Sleep(time.Second * 5)
	}
}
