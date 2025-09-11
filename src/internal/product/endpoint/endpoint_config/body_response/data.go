package body_response

type S_BodyResponse struct {
	Body []byte

	ExtractionResponse string `yaml:"extraction_response" validate:"required"`
	FormatResponse     string `yaml:"format_response" validate:"required,lowercase,oneof=json zip/json"`
}
