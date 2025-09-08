package bearer

import endpointconfig "extrator/internal/product/endpoint/endpointConfig"

type S_Bearer struct {
	Name           string                           `yaml:"name" validate:"required"`
	Prefix         string                           `yaml:"prefix" validate:"required"`
	Token          *string                          `yaml:"_token"`
	IsDynamic      bool                             `yaml:"is_dynamic"`
	EndpointConfig *endpointconfig.S_EndpointConfig `yaml:"_endpoint_config" validate:"required_if=IsDynamic true"`
}
