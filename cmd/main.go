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

	for productName := range products {
		productEndpoints := products[productName]
		// TODO, adicionar controle de produtos
		utils.AsyncWithLimit(3, &wg, func() {
			for index := range productEndpoints {
				// TODO, adicionar controle de endpoints
				utils.AsyncWithLimit(10, &wg, func() {
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
