package paginate

import "net/http"

const (
	MODE_NONE        = "none"
	MODE_PAGE        = "page"
	MODE_PROPERTY    = "property"
	MODE_LINK_HEADER = "link_header"

	DIRECTION_NEXT     = "next"
	DIRECTION_PREVIOUS = "previous"

	LOCATION_HEADER      = "header"
	LOCATION_BODY        = "body"
	LOCATION_QUERY_PARAM = "query_param"

	METHOD_GET  = http.MethodGet
	METHOD_POST = http.MethodPost

	DEFAULT_MODE      = MODE_NONE
	DEFAULT_DIRECTION = DIRECTION_NEXT
	DEFAULT_LOCATION  = LOCATION_QUERY_PARAM
)

var (
	MODE_ALLOWED      = []string{MODE_NONE, MODE_PAGE, MODE_PROPERTY, MODE_LINK_HEADER}
	DIRECTION_ALLOWED = []string{DIRECTION_NEXT, DIRECTION_PREVIOUS}
	LOCATION_ALLOWED  = []string{LOCATION_HEADER, LOCATION_BODY, LOCATION_QUERY_PARAM}
)

type S_PaginatePage struct {
	Page      string  `yaml:"page,omitempty" validate:"required,printascii"`
	PageSize  string  `yaml:"page_size,omitempty" validate:"required,printascii"`
	Direction *string `yaml:"_direction,omitempty"`
	Location  *string `yaml:"_location,omitempty"`
}

type S_PaginateProperty struct {
	Property string `yaml:"property,omitempty" validate:"required,printascii"`
}

type S_PaginateLinkHeader struct {
	Header string `yaml:"header,omitempty" validate:"required,printascii"`
}

type S_Paginate struct {
	Mode string `yaml:"mode,omitempty" validate:"required,printascii"`

	Page       *S_PaginatePage       `yaml:"_page,omitempty" validate:"required_if Mode page"`
	Property   *S_PaginateProperty   `yaml:"_property,omitempty" validate:"required_if Mode property"`
	LinkHeader *S_PaginateLinkHeader `yaml:"_link_header,omitempty" validate:"required_if Mode link_header"`
}
