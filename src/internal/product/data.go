package product

import (
	"extrator/internal/product/auth"
	"extrator/internal/product/endpoint"
)

type S_Product struct {
	Name     string
	Endpoint endpoint.S_Endpoint
	Auth     *auth.S_Auth
}
