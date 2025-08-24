package pagination

import "strings"

func (po *S_PaginationOffset) format() {
	if po == nil {
		panic("Offset is required")
	}

	po.Direction = paginationDirection(strings.TrimSpace(strings.ToLower(string(po.Direction))))
	po.Location = paginationLocation(strings.TrimSpace(strings.ToLower(string(po.Location))))
	po.Offset = strings.TrimSpace(po.Offset)
	po.Limit = strings.TrimSpace(po.Limit)
}

func (pp *S_PaginationPage) format() {
	if pp == nil {
		panic("Page is required")
	}

	pp.Direction = paginationDirection(strings.TrimSpace(strings.ToLower(string(pp.Direction))))
	pp.Location = paginationLocation(strings.TrimSpace(strings.ToLower(string(pp.Location))))
	pp.Page = strings.TrimSpace(pp.Page)
	if pp.PageSize <= 0 {
		pp.PageSize = 1
	}
}

func (pp *S_PaginationProperty) format() {
	if pp == nil {
		panic("Property is required")
	}

	pp.Property = strings.TrimSpace(pp.Property)
}

func (plh *S_PaginationLinkHeader) format() {
	if plh == nil {
		panic("LinkHeader is required")
	}

	plh.Header = strings.TrimSpace(plh.Header)
}

func (p *S_Pagination) format() {
	p.Mode = paginationMode(strings.TrimSpace(strings.ToLower(string(p.Mode))))
}
