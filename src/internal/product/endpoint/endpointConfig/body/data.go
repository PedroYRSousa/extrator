package body

import (
	multipart "extrator/internal/product/endpoint/endpointConfig/body/multiPart"
)

type S_Body struct {
	MultiPart      *multipart.S_Multipart `yaml:"_multi_part" validate:"required_without=FormURLEncoded File Text"`
	FormURLEncoded *[]map[string]string   `yaml:"_form_url_encoded" validate:"required_without=MultiPart File Text"`
	File           *string                `yaml:"_file" validate:"required_without=MultiPart FormURLEncoded Text"`
	Text           *string                `yaml:"_text" validate:"required_without=MultiPart FormURLEncoded File"`
}
