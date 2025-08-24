package pagination

type paginationMode string
type paginationDirection string
type paginationLocation string

const (
	PAGINATION_MODE_NONE        paginationMode = "none"
	PAGINATION_MODE_OFFSET      paginationMode = "offset"
	PAGINATION_MODE_PAGE        paginationMode = "page"
	PAGINATION_MODE_PROPERTY    paginationMode = "property"
	PAGINATION_MODE_LINK_HEADER paginationMode = "link_header"

	PAGINATION_DIRECTION_NEXT     paginationDirection = "next"
	PAGINATION_DIRECTION_PREVIOUS paginationDirection = "previous"

	PAGINATION_LOCATION_HEADER      paginationLocation = "header"
	PAGINATION_LOCATION_BODY        paginationLocation = "body"
	PAGINATION_LOCATION_QUERY_PARAM paginationLocation = "query_param"
)

var (
	AVAILABLE_LOCATIONS  = []paginationLocation{PAGINATION_LOCATION_HEADER, PAGINATION_LOCATION_BODY, PAGINATION_LOCATION_QUERY_PARAM}
	AVAILABLE_DIRECTIONS = []paginationDirection{PAGINATION_DIRECTION_NEXT, PAGINATION_DIRECTION_PREVIOUS}
	AVAILABLE_MODES      = []paginationMode{PAGINATION_MODE_NONE, PAGINATION_MODE_OFFSET, PAGINATION_MODE_PAGE, PAGINATION_MODE_PROPERTY, PAGINATION_MODE_LINK_HEADER}
)

type S_PaginationOffset struct {
	Direction paginationDirection `yaml:"direction"`
	Location  paginationLocation  `yaml:"location"`
	Offset    string              `yaml:"offset"`
	Limit     string              `yaml:"limit"`
}

type S_PaginationPage struct {
	Direction paginationDirection `yaml:"direction"`
	Location  paginationLocation  `yaml:"location"`
	Page      string              `yaml:"page"`
	PageSize  int                 `yaml:"page_size"`
}

type S_PaginationProperty struct {
	Property string `yaml:"property"`
}

type S_PaginationLinkHeader struct {
	Header string `yaml:"header"`
}

type S_Pagination struct {
	Mode paginationMode `yaml:"mode"`

	// Opcionais
	Offset     *S_PaginationOffset     `yaml:"offset,omitempty"`
	Page       *S_PaginationPage       `yaml:"page,omitempty"`
	Property   *S_PaginationProperty   `yaml:"property,omitempty"`
	LinkHeader *S_PaginationLinkHeader `yaml:"link_header,omitempty"`
}
