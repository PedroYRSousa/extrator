package auth

import (
	apikey "extrator/internal/product/endpoint/auth/apiKey"
	"extrator/internal/product/endpoint/auth/basic"
	"extrator/internal/product/endpoint/auth/bearer"
	"extrator/internal/product/endpoint/auth/digest"
)

type S_Auth struct {
	ApiKey *apikey.S_ApiKey `yaml:"_api_key" validate:"required_without=Basic Digest Bearer"`
	Basic  *basic.S_Basic   `yaml:"_basic" validate:"required_without=ApiKey Digest Bearer"`
	Digest *digest.S_Digest `yaml:"_digest" validate:"required_without=Basic ApiKey Bearer"`
	Bearer *bearer.S_Bearer `yaml:"_bearer" validate:"required_without=Basic Digest ApiKey"`
}
