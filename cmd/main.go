package main

import (
	"extrator/config"
	"extrator/product"
	"log"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatalf("ERROR | Load config | %+v", err)
	}

	log.Printf("%+v", *conf)
	log.Printf("%+v", conf.Products)

	products, err := product.Load(*conf)
	if err != nil {
		log.Fatalf("ERROR | Load products | %+v", err)
	}
	log.Printf("%+v", products)

	// for _, product := range products {
	// 	err := product.Mount(conf)
	// 	if err != nil {
	// 		log.Fatalf("ERROR | Mount product (%s) | %+v", product.Name, err)
	// 	}
	// }

	// productsAsyncCount := modules.AsyncLimiter(int(*conf.Products.ProductAsyncCount))
	// for _, product := range products {
	// 	_product := product
	// 	productsAsyncCount.Go(func() {
	// 		err := modules.ExecWithAttempts(*conf.Products.MaxAttempts, *conf.Products.DelayAttemptsInSeconds, func() error {
	// 			err := _product.Start(conf)
	// 			if err != nil {
	// 				return fmt.Errorf("ERROR | Start product (%s) | %+v", _product.Name, err)
	// 			}

	// 			return nil
	// 		})
	// 		if err != nil {
	// 			log.Println(err)
	// 		}
	// 	})

	// }
	// productsAsyncCount.Wait()
}
