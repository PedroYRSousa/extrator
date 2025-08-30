package page

const (
	DIRECTION_NEXT     = "next"
	DIRECTION_PREVIOUS = "previous"

	LOCATION_HEADER      = "header"
	LOCATION_BODY        = "body"
	LOCATION_QUERY_PARAM = "query_param"

	DEFAULT_DIRECTION = DIRECTION_NEXT
	DEFAULT_LOCATION  = LOCATION_QUERY_PARAM
)

var (
	DIRECTION_ALLOWED = []string{DIRECTION_NEXT, DIRECTION_PREVIOUS}
	LOCATION_ALLOWED  = []string{LOCATION_HEADER, LOCATION_BODY, LOCATION_QUERY_PARAM}
)

type S_PaginatePage struct {
	Page      string  `yaml:"page,omitempty" validate:"required,printascii"`
	PageSize  string  `yaml:"page_size,omitempty" validate:"required,printascii"`
	Direction *string `yaml:"_direction,omitempty"`
	Location  *string `yaml:"_location,omitempty"`
}
