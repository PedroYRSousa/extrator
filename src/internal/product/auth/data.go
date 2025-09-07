package auth

import (
	apikey "extrator/internal/product/auth/apiKey"
	"extrator/internal/product/auth/basic"
	"extrator/internal/product/auth/bearer"
	"extrator/internal/product/auth/digest"
)

type S_Auth struct {
	ApiKey *apikey.S_ApiKey
	Basic  *basic.S_Basic
	Digest *digest.S_Digest
	Bearer *bearer.S_Bearer
}
