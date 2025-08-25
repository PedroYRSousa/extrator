package main

import (
	"extrator/internal/product"
	"log"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	log.Println("Start extraction")

	// TODO, adicionar controle de versão

	var wg sync.WaitGroup

	_, err := product.Load()
	if err != nil {
		log.Fatalln("ERROR", err)
	}

	// for productName := range products {
	// 	productEndpoints := products[productName]
	// 	// TODO, adicionar controle de produtos
	// 	utils.AsyncWithLimit(3, &wg, func() {
	// 		for index := range productEndpoints {
	// 			// TODO, adicionar controle de endpoints
	// 			utils.AsyncWithLimit(10, &wg, func() {
	// 				productEndpoints[index].Start()
	// 			})
	// 		}
	// 	})
	// }

	wg.Wait()

	log.Println("End extraction", time.Since(startTime))
}
