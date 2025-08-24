package main

import (
	"extrator/internal/product"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	products, err := product.Load()
	if err != nil {
		log.Fatalln(err)
	}

	for productName, productData := range products {
		log.Println(productName)
		log.Println(productData)
	}

	wg.Wait()
}
