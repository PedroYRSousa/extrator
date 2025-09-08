package main

import (
	"extrator/internal/configs"
	"extrator/internal/product"
	utilsstructs "extrator/pkg/utils_structs"
	"log"
)

func main() {
	conf, err := configs.Load()
	if err != nil {
		log.Fatalln(err)
	}

	utilsstructs.Show(conf, 0)

	products, err := product.Loads(conf)
	if err != nil {
		log.Fatalln(err)
	}

	for _, product := range products {
		utilsstructs.Show(product, 0)
	}
}
