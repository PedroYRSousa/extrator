package product

import "extrator/product/endpoint"

func New() S_Product {
	return S_Product{
		Endpoints: []endpoint.S_Endpoint{},
	}
}
