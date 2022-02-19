package workers2

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Run() {
	// Poderia ser um channel de uma struct de 'Funcionario'
	requestId := make(chan int)

	// Quantidade de processos paralelos
	councurrency := 500
	for i := 1; i <= councurrency; i++ {
		go worker(requestId, i)
	}

	// Poderia ser um array de struct de 'Funcionario'
	for i := 1; i <= 1000; i++ {
		requestId <- i
	}
}

// Poderia ser uma função que faz algum processamento em uma struct de 'Funcionario'
func worker(requestId chan int, w int) {
	for r := range requestId {
		// Adicionado função anônima para que ao termino da execução da função seja executado o 'defer'
		// Sem a função a execução do 'defer' só ocorreria ao termino da função 'worker'
		// Poderia ser uma função nomeada sendo chamada no loop
		func() {
			res, err := http.Get("http://localhost:3000/products")
			if err != nil {
				log.Fatal("Erro")
			}

			// defer -> executado após todos os processos da função acabarem (como se fosse um dispose)
			defer res.Body.Close()

			content, _ := ioutil.ReadAll(res.Body)
			fmt.Printf("Worker %d, RequestId %d, Content %s", w, r, string(content))
			time.Sleep(time.Second * 2)
		}()
	}
}
