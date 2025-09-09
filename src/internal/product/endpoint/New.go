package endpoint

import endpointconfig "extrator/internal/product/endpoint/endpointConfig"

func New(nameFile string, pathFile string) *S_Endpoint {
	return &S_Endpoint{
		ToAbort:        false,
		Name:           nameFile,
		Path:           pathFile,
		EndpointConfig: endpointconfig.New(),
		Auth:           nil,
	}
}
