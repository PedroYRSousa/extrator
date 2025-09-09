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

	asyncProducts := async.AsyncLimiter(conf.ProductsParallelExecutionCount)

	for _, product := range products {
		_product := product
		_conf := conf
		utilsstructs.Show(product, 0)

		asyncProducts.Go(func() { _product.Mount(_conf) })

		asyncProducts.Wait()

		// asyncProducts.Go(_product.Start)
	}

	asyncProducts.Wait()
}
