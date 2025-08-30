package paginate

import (
	"extrator/product/endpoint/paginate/link_header"
	"extrator/product/endpoint/paginate/page"
	"extrator/product/endpoint/paginate/property"
)

const (
	MODE_NONE        = "none"
	MODE_PAGE        = "page"
	MODE_PROPERTY    = "property"
	MODE_LINK_HEADER = "link_header"

	DEFAULT_MODE = MODE_NONE
)

var (
	MODE_ALLOWED = []string{MODE_NONE, MODE_PAGE, MODE_PROPERTY, MODE_LINK_HEADER}
)

type S_Paginate struct {
	Mode string `yaml:"mode,omitempty" validate:"required,printascii"`

	Page       *page.S_PaginatePage              `yaml:"_page,omitempty" validate:"required_if=Mode page"`
	Property   *property.S_PaginateProperty      `yaml:"_property,omitempty" validate:"required_if=Mode property"`
	LinkHeader *link_header.S_PaginateLinkHeader `yaml:"_link_header,omitempty" validate:"required_if=Mode link_header"`
}
