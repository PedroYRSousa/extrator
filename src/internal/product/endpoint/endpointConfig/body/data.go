package body

import (
	multipart "extrator/internal/product/endpoint/endpointConfig/body/multiPart"
)

type S_Body struct {
	MultiPart      *multipart.S_Multipart
	FormURLEncoded *map[string]string
	File           *string
	Text           *string
}
