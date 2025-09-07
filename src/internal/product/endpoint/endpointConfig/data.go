package endpointconfig

import (
	"extrator/internal/product/endpoint/endpointConfig/body"
)

type S_EndpointConfig struct {
	URL    string
	Method *string
	// ExtractJsonPath *string
	Headers     *map[string]string
	QueryParams *map[string]string
	Body        *body.S_Body
}
