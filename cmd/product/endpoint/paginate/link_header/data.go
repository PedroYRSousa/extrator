package link_header

type S_PaginateLinkHeader struct {
	Header string `yaml:"header,omitempty" validate:"required,printascii"`
}
