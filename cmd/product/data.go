package product

import "extrator/product/endpoint"

type S_Product struct {
	// Obrigatório
	Name      string
	Path      string
	Endpoints []endpoint.S_Endpoint

	// Opcionais
}
