package endpoint_config

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func (po *S_PaginationOffset) mount(request *http.Request) error {
	return nil
}

func (pp *S_PaginationPage) mount(request *http.Request) error {
	return nil
}

func (pp *S_PaginationProperty) mount(request *http.Request) error {
	return nil
}

func (plh *S_PaginationLinkHeader) mount(request *http.Request) error {
	return nil
}

func (p *S_Pagination) mount(request *http.Request) error {
	if p.Mode == PAGINATION_MODE_NONE {
		return nil
	}

	switch p.Mode {
	case PAGINATION_MODE_OFFSET:
		return p.Offset.mount(request)
	case PAGINATION_MODE_PAGE:
		return p.Page.mount(request)
	case PAGINATION_MODE_PROPERTY:
		return p.Property.mount(request)
	case PAGINATION_MODE_LINK_HEADER:
		return p.LinkHeader.mount(request)
	}

	return nil
}

func (ec *S_EndpointConfig) Mount(request *http.Request) error {
	request.Method = string(*ec.Method)

	header := http.Header{}
	if ec.Headers != nil {
		for k, v := range *ec.Headers {
			header.Add(k, v)
		}
	}
	request.Header = header

	if ec.QueryParams != nil {
		query := request.URL.Query()

		for k, v := range *ec.QueryParams {
			query.Add(k, v)
		}

		request.URL.RawQuery = query.Encode()
	}

	// Por hora só aceito body do tipo json
	// TODO, adicionar outros tipos
	request.Body = http.NoBody
	if ec.Body != nil {
		bodyBytes, err := json.Marshal(*ec.Body)
		if err != nil {
			return err
		}
		request.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		request.ContentLength = int64(len(bodyBytes))
	}

	return nil
}
