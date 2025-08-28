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

	products, err := product.Load(conf)
	if err != nil {
		log.Fatalf("ERROR | Load products | %+v", err)
	}

	log.Println("config", conf, err)
	log.Println("product", products, err)
}
