package endpoint_config

import (
	"extrator/internal/utils"
	"extrator/pkg/env"
)

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

func (r *S_Retry) transform() error {
	return nil
}

func (ec *S_EndpointConfig) Transform() error {
	if ec.Headers != nil {
		for k, v := range *ec.Headers {
			if utils.IsEnv(v) {
				newValue, err := env.Get(v)
				if err != nil {
					return err
				}
				(*ec.Headers)[k] = newValue
			} else if utils.IsSecret(v) {
				// TODO, Pegar da aws secretmanager
			}
		}
	}

	if ec.Body != nil {
		for k, v := range *ec.Body {
			if utils.IsEnv(v) {
				newValue, err := env.Get(v)
				if err != nil {
					return err
				}
				(*ec.Body)[k] = newValue
			} else if utils.IsSecret(v) {
				// TODO, Pegar da aws secretmanager
			}
		}
	}

	if ec.QueryParams != nil {
		for k, v := range *ec.QueryParams {
			if utils.IsEnv(v) {
				newValue, err := env.Get(v)
				if err != nil {
					return err
				}
				(*ec.QueryParams)[k] = newValue
			} else if utils.IsSecret(v) {
				// TODO, Pegar da aws secretmanager
			}
		}
	}

	err := ec.Retry.transform()
	if err != nil {
		return err
	}

	return nil
}
