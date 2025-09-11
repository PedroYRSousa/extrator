package endpoint

import (
	"extrator/internal/product/endpoint/auth"
	"extrator/internal/product/endpoint/endpoint_config"
)

type S_Endpoint struct {
	Path    string `validate:"filepath,printascii,gt=0"`
	ToAbort bool

	Name           string                            `yaml:"name" validate:"required,printascii,gt=0"`
	EndpointConfig *endpoint_config.S_EndpointConfig `yaml:"endpoint_config" validate:"required"`
	Auth           *auth.S_Auth                      `yaml:"_auth"`
}
