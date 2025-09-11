package bearer

import "extrator/internal/product/endpoint/endpoint_config"

type S_Bearer struct {
	Name           string                            `yaml:"name" validate:"required"`
	Prefix         string                            `yaml:"prefix" validate:"required"`
	Token          *string                           `yaml:"_token"`
	IsDynamic      bool                              `yaml:"is_dynamic"`
	EndpointConfig *endpoint_config.S_EndpointConfig `yaml:"_endpoint_config" validate:"required_if=IsDynamic true"`
}
