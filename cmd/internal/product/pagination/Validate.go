package pagination

import (
	"errors"
	"fmt"
	"slices"
)

func (po *S_PaginationOffset) validate() error {
	po.format()

	if po.Direction == "" {
		return errors.New("check pagination.offset.direction | value cannot be empty")
	}

	if !slices.Contains(AVAILABLE_DIRECTIONS, po.Direction) {
		return fmt.Errorf("check pagination.offset.direction | available options: %v", AVAILABLE_DIRECTIONS)
	}

	if po.Location == "" {
		return errors.New("check pagination.offset.location | value cannot be empty")
	}

	if !slices.Contains(AVAILABLE_LOCATIONS, po.Location) {
		return fmt.Errorf("check pagination.offset.location | available options: %v", AVAILABLE_LOCATIONS)
	}

	if po.Offset == "" {
		return errors.New("check pagination.offset.offset | value cannot be empty")
	}

	if po.Limit == "" {
		return errors.New("check pagination.offset.limit | value cannot be empty")
	}

	return nil
}

func (pp *S_PaginationPage) validate() error {
	pp.format()

	if pp.Direction == "" {
		return errors.New("check pagination.page.direction | value cannot be empty")
	}

	if !slices.Contains(AVAILABLE_DIRECTIONS, pp.Direction) {
		return fmt.Errorf("check pagination.page.direction | available options: %v", AVAILABLE_DIRECTIONS)
	}

	if pp.Location == "" {
		return errors.New("check pagination.page.location | value cannot be empty")
	}

	if !slices.Contains(AVAILABLE_LOCATIONS, pp.Location) {
		return fmt.Errorf("check pagination.page.location | available options: %v", AVAILABLE_LOCATIONS)
	}

	if pp.Page == "" {
		return errors.New("check pagination.page.page | value cannot be empty")
	}

	if pp.PageSize <= 0 {
		return errors.New("check pagination.page.page_size | value must be greater than 0")
	}

	return nil
}

func (pp *S_PaginationProperty) validate() error {
	pp.format()

	if pp.Property == "" {
		return errors.New("check pagination.property.property | value cannot be empty")
	}

	return nil
}

func (plh *S_PaginationLinkHeader) validate() error {
	plh.format()

	if plh.Header == "" {
		return errors.New("check pagination.link_header.header | value cannot be empty")
	}

	return nil
}

func (p *S_Pagination) Validate() error {
	p.format()

	if !slices.Contains(AVAILABLE_MODES, p.Mode) {
		return fmt.Errorf("check pagination.mode | available options: %v", AVAILABLE_MODES)
	}

	if p.Mode == PAGINATION_MODE_NONE {
		return nil
	}

	switch p.Mode {
	case PAGINATION_MODE_OFFSET:
		if p.Offset == nil {
			return errors.New("check pagination.offset | offset cannot be null when mode is 'offset'")
		}
		return p.Offset.validate()
	case PAGINATION_MODE_PAGE:
		if p.Page == nil {
			return errors.New("check pagination.page | page cannot be null when mode is 'page'")
		}
		return p.Page.validate()
	case PAGINATION_MODE_PROPERTY:
		if p.Property == nil {
			return errors.New("check pagination.property | property cannot be null when mode is 'property'")
		}
		return p.Property.validate()
	case PAGINATION_MODE_LINK_HEADER:
		if p.LinkHeader == nil {
			return errors.New("check pagination.link_header | link_header cannot be null when mode is 'link_header'")
		}
		return p.LinkHeader.validate()
	}

	return nil
}
