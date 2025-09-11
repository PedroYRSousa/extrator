package main

import (
	"extrator/internal/configs"
	"extrator/internal/product"
	"extrator/pkg/async"
	utils_structs "extrator/pkg/utils_structs"
	"log"
)

func main() {
	conf, err := configs.Load()
	if err != nil {
		log.Fatalln(err)
	}

	utils_structs.Show(conf, 0)

	products, err := product.Loads(conf)
	if err != nil {
		log.Fatalln(err)
	}

	asyncProducts := async.AsyncLimiter(conf.ProductsParallelExecutionCount)

	for _, product := range products {
		_product := product
		_conf := conf
		utils_structs.Show(product, 0)

		asyncProducts.Go(func() { _product.Mount(_conf) })

		asyncProducts.Wait()

		// asyncProducts.Go(_product.Start)
	}

	asyncProducts.Wait()
}
