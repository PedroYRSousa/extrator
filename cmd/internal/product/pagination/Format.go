package pagination

import "strings"

func (po *S_PaginationOffset) Format() {
	if po != nil {
		po.Direction = paginationDirection(strings.TrimSpace(strings.ToLower(string(po.Direction))))
		po.Location = paginationLocation(strings.TrimSpace(strings.ToLower(string(po.Location))))
		po.Offset = strings.TrimSpace(po.Offset)
		po.Limit = strings.TrimSpace(po.Limit)
	}
}

func (pp *S_PaginationPage) Format() {
	if pp != nil {
		pp.Direction = paginationDirection(strings.TrimSpace(strings.ToLower(string(pp.Direction))))
		pp.Location = paginationLocation(strings.TrimSpace(strings.ToLower(string(pp.Location))))
		pp.Page = strings.TrimSpace(pp.Page)
		if pp.PageSize <= 0 {
			pp.PageSize = 1
		}
	}
}

func (pp *S_PaginationProperty) Format() {
	if pp != nil {
		pp.Property = strings.TrimSpace(pp.Property)
	}
}

func (plh *S_PaginationLinkHeader) Format() {
	if plh != nil {
		plh.Header = strings.TrimSpace(plh.Header)
	}
}

func (p *S_Pagination) Format() {
	p.Mode = paginationMode(strings.TrimSpace(strings.ToLower(string(p.Mode))))

	p.Offset.Format()
	p.Page.Format()
	p.Property.Format()
	p.LinkHeader.Format()
}
