package link_header

import (
	"extrator/modules"
)

func (paginateLinkHeader *S_PaginateLinkHeader) format() *S_PaginateLinkHeader {
	modules.FormatString(paginateLinkHeader)

	return paginateLinkHeader
}
