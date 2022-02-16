package workers

import (
	"fmt"
	"time"
)

func Work() {
	channel := make(chan int)

	// Criando os workers
	for i := 1; i <= 10; i++ {
		go worker(channel)
	}

	// Adicionando valor ao channel
	for i := 0; i <= 100; i++ {
		channel <- i
	}
}

func worker(channel chan int) {
	for i := range channel {
		fmt.Println(i)
		time.Sleep(time.Second * 5)
	}
}
