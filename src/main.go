package main

import (
	"extrator/internal/configs"
	"extrator/internal/product"
	"extrator/pkg/async"
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

	products1, err := product.Loads(conf)
	if err != nil {
		log.Fatalln(err)
	}

	products2, err := product.Loads(conf)
	if err != nil {
		log.Fatalln(err)
	}

	products = append(products, products1...)
	products = append(products, products2...)

	asyncProducts := async.AsyncLimiter(conf.ProductsParallelExecutionCount)

	for _, product := range products {
		_product := product
		utilsstructs.Show(product, 0)

		asyncProducts.Go(func() {

		})
	}

	asyncProducts.Wait()
}
