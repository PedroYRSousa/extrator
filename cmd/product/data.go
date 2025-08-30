package product

import "extrator/product/endpoint"

type S_Product struct {
	Path      string                `validate:"required,dirpath,printascii"`
	Name      string                `validate:"required,printascii"`
	Endpoints []endpoint.S_Endpoint `validate:"required,gt=0"`
}
