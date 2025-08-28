package endpoint_config

import (
	"errors"
	"extrator/internal/utils"
	"fmt"
	"slices"
	"strings"
)

func (po *S_PaginationOffset) validate(ec *S_EndpointConfig) error {
	if po == nil {
		panic("check pagination.offset | offset cannot be null when mode is 'offset'")
	}

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

func (pp *S_PaginationPage) validate(ec *S_EndpointConfig) error {
	if pp == nil {
		panic("check pagination.page | page cannot be null when mode is 'page'")
	}

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

func (pp *S_PaginationProperty) validate(ec *S_EndpointConfig) error {
	if pp == nil {
		panic("check pagination.property | property cannot be null when mode is 'property'")
	}

	pp.format()

	if pp.Property == "" {
		return errors.New("check pagination.property.property | value cannot be empty")
	}

	return nil
}

func (plh *S_PaginationLinkHeader) validate() error {
	if plh == nil {
		panic("check pagination.link_header | link_header cannot be null when mode is 'link_header'")
	}

	plh.format()

	if plh.Header == "" {
		return errors.New("check pagination.link_header.header | value cannot be empty")
	}

	return nil
}

func (p *S_Pagination) Validate(ec *S_EndpointConfig) error {
	p.format()

	if !slices.Contains(AVAILABLE_MODES, p.Mode) {
		return fmt.Errorf("check pagination.mode | available options: %v", AVAILABLE_MODES)
	}

	switch p.Mode {
	case PAGINATION_MODE_OFFSET:
		return p.Offset.validate(ec)
	case PAGINATION_MODE_PAGE:
		return p.Page.validate(ec)
	case PAGINATION_MODE_PROPERTY:
		return p.Property.validate(ec)
	case PAGINATION_MODE_LINK_HEADER:
		return p.LinkHeader.validate()
	}

	return nil
}

func (r *S_Retry) validate() error {
	r.format()

	return nil
}

func (ec *S_EndpointConfig) Validate() error {
	ec.format()

	if !slices.Contains(METHODS_AVAILABLE, *ec.Method) {
		return fmt.Errorf("check endpoint_config.method | available options: %v", METHODS_AVAILABLE)
	}

	if ec.Headers != nil {
		for key, value := range *(ec.Headers) {
			if strings.TrimSpace(key) == "" {
				return errors.New("check endpoint_config.headers | header names cannot be empty")
			}

			if value != "" {
				if !utils.CheckIsHardCodedSecretOrEnv(value) {
					return errors.New("check endpoint_config.headers | body parameter values must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
				}
			}
		}
	}

	if ec.Body != nil {
		for key, value := range *(ec.Body) {
			if strings.TrimSpace(key) == "" {
				return errors.New("check endpoint_config.body | body parameter names cannot be empty")
			}

			if value != "" {
				if !utils.CheckIsHardCodedSecretOrEnv(value) {
					return errors.New("check endpoint_config.body | body parameter values must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
				}
			}
		}
	}

	if ec.QueryParams != nil {
		for key, value := range *(ec.QueryParams) {
			if strings.TrimSpace(key) == "" {
				return errors.New("check endpoint_config.query_params | query_params parameter names cannot be empty")
			}

			if value != "" {
				if !utils.CheckIsHardCodedSecretOrEnv(value) {
					return errors.New("check endpoint_config.query_params | query_params parameter values must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
				}
			}
		}
	}

	err := ec.Retry.validate()
	if err != nil {
		return err
	}

	err = ec.Pagination.Validate(ec)
	if err != nil {
		return err
	}

	return nil
}
