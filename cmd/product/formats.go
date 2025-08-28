package product

import (
	"extrator/modules"
)

func (product *S_Product) format() *S_Product {
	if product == nil {
		panic("TODO, Melhora isso aqui")
	}

	modules.FormatString(product)

	return product
}
