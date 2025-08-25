package pagination

func (po *S_PaginationOffset) transform() error {
	return nil
}

func (pp *S_PaginationPage) transform() error {
	return nil
}

func (pp *S_PaginationProperty) transform() error {
	return nil
}

func (plh *S_PaginationLinkHeader) transform() error {
	return nil
}

func (p *S_Pagination) Transform() error {
	if p.Mode == PAGINATION_MODE_NONE {
		return nil
	}

	switch p.Mode {
	case PAGINATION_MODE_OFFSET:
		return p.Offset.transform()
	case PAGINATION_MODE_PAGE:
		return p.Page.transform()
	case PAGINATION_MODE_PROPERTY:
		return p.Property.transform()
	case PAGINATION_MODE_LINK_HEADER:
		return p.LinkHeader.transform()
	}

	return nil
}
