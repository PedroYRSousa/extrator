package product

import (
	"extrator/internal/product/endpoint"
)

type S_Product struct {
	Name string `validate:"printascii,gt=0"`
	Path string `validate:"dirpath,printascii,gt=0"`

	Endpoints []*endpoint.S_Endpoint `validate:"required,gt=0"`
}
