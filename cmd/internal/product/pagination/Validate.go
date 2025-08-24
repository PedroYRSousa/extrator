package pagination

import (
	"errors"
	"fmt"
	"slices"
)

func (po *S_PaginationOffset) validate() error {
	if po.Direction == "" {
		return errors.New("invalid pagination | Check pagination.offset.direction | value cannot be empty")
	}

	if !slices.Contains(availableDirections, po.Direction) {
		return fmt.Errorf("invalid pagination | Check pagination.offset.direction | available options: %v", availableDirections)
	}

	if po.Location == "" {
		return errors.New("invalid pagination | Check pagination.offset.location | value cannot be empty")
	}

	if !slices.Contains(availableLocations, po.Location) {
		return fmt.Errorf("invalid pagination | Check pagination.offset.location | available options: %v", availableLocations)
	}

	if po.Offset == "" {
		return errors.New("invalid pagination | Check pagination.offset.offset | value cannot be empty")
	}

	if po.Limit == "" {
		return errors.New("invalid pagination | Check pagination.offset.limit | value cannot be empty")
	}

	return nil
}

func (pp *S_PaginationPage) validate() error {
	if pp.Direction == "" {
		return errors.New("invalid pagination | Check pagination.page.direction | value cannot be empty")
	}

	if !slices.Contains(availableDirections, pp.Direction) {
		return fmt.Errorf("invalid pagination | Check pagination.page.direction | available options: %v", availableDirections)
	}

	if pp.Location == "" {
		return errors.New("invalid pagination | Check pagination.page.location | value cannot be empty")
	}

	if !slices.Contains(availableLocations, pp.Location) {
		return fmt.Errorf("invalid pagination | Check pagination.page.location | available options: %v", availableLocations)
	}

	if pp.Page == "" {
		return errors.New("invalid pagination | Check pagination.page.page | value cannot be empty")
	}

	if pp.PageSize <= 0 {
		return errors.New("invalid pagination | Check pagination.page.page_size | value must be greater than 0")
	}

	return nil
}

func (pp *S_PaginationProperty) validate() error {
	if pp.Property == "" {
		return errors.New("invalid pagination | Check pagination.property.property | value cannot be empty")
	}

	return nil
}

func (plh *S_PaginationLinkHeader) validate() error {
	if plh.Header == "" {
		return errors.New("invalid pagination | Check pagination.link_header.header | value cannot be empty")
	}

	return nil
}

func (p *S_Pagination) Validate() error {
	if !slices.Contains(availableModes, p.Mode) {
		return fmt.Errorf("invalid pagination | Check pagination.mode | available options: %v", availableModes)
	}

	if p.Mode == PAGINATION_MODE_NONE {
		return nil
	}

	switch p.Mode {
	case PAGINATION_MODE_LINK_HEADER:
		if err := p.LinkHeader.validate(); err != nil {
			return err
		}
	case PAGINATION_MODE_PROPERTY:
		if err := p.Property.validate(); err != nil {
			return err
		}
	case PAGINATION_MODE_PAGE:
		if err := p.Page.validate(); err != nil {
			return err
		}
	case PAGINATION_MODE_OFFSET:
		if err := p.Offset.validate(); err != nil {
			return err
		}
	}

	return nil
}
