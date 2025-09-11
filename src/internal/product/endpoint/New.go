package endpoint

import "extrator/internal/product/endpoint/endpoint_config"

func New(nameFile string, pathFile string) *S_Endpoint {
	return &S_Endpoint{
		ToAbort:        false,
		Name:           nameFile,
		Path:           pathFile,
		EndpointConfig: endpoint_config.New(),
		Auth:           nil,
	}
}
