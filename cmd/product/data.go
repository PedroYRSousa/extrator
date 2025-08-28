package product

import "extrator/product/endpoint"

type S_Product struct {
	Name      string                `validate:"required,printascii"`
	Path      string                `validate:"required,dir,printascii"`
	Endpoints []endpoint.S_Endpoint `validate:"required,gt=0"`
}
