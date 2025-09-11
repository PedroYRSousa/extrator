package auth

import (
	"extrator/internal/product/endpoint/auth/api_key"
	"extrator/internal/product/endpoint/auth/basic"
	"extrator/internal/product/endpoint/auth/bearer"
)

type S_Auth struct {
	ApiKey *api_key.S_ApiKey `yaml:"_api_key" validate:"required_without=Basic Digest Bearer"`
	Basic  *basic.S_Basic    `yaml:"_basic" validate:"required_without=ApiKey Digest Bearer"`
	Bearer *bearer.S_Bearer  `yaml:"_bearer" validate:"required_without=Basic Digest ApiKey"`
}
