package main

import (
	"extrator/internal/product"
	"extrator/internal/utils"
	"log"
	"sync"
)

func main() {
	log.Println("Start extraction")

	// TODO, adicionar controle de versão

	var wg sync.WaitGroup

	products, err := product.Load()
	if err != nil {
		log.Fatalln("ERROR", err)
	}

	// TODO, adicionar controle de produtos
	var executorProducts *utils.S_AsyncLimiter = utils.AsyncLimiter(3)
	for productName := range products {
		productEndpoints := products[productName]
		executorProducts.Go(func() {
			// TODO, adicionar controle de endpoints
			var executorEndpoints *utils.S_AsyncLimiter = utils.AsyncLimiter(10)
			for index := range productEndpoints {
				executorEndpoints.Go(func() {
					err := productEndpoints[index].Start()
					if err != nil {
						log.Fatalln("ERROR", err)
					}
				})
			}
		})
	}

	wg.Wait()
}
