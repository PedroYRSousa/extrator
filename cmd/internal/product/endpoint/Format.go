package endpoint

import (
	"strings"
)

func (e *S_Endpoint) format() {
	if e == nil {
		panic("TODO, Dizer aqui que endpoint é obrigatório")
	}

	e.URI = strings.TrimSpace(strings.ToLower(e.URI))
	e.Method = strings.TrimSpace(strings.ToUpper(e.Method))
	e.ResponseFormat = endpointResponseFormat(strings.TrimSpace(strings.ToLower(string(e.ResponseFormat))))

	if e.LimitExtract == nil {
		e.LimitExtract = new(int)
		*e.LimitExtract = ENDPOINT_LIMIT_EXTRACT_DEFAULT
	}
}
