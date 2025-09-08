package product

import "extrator/internal/product/endpoint"

func New(nameFile string, pathFile string) *S_Product {
	return &S_Product{
		Name:      nameFile,
		Path:      pathFile,
		Endpoints: []*endpoint.S_Endpoint{},
	}
}
