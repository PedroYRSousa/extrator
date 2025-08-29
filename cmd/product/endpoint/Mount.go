package endpoint

import (
	"extrator/config"
	"net/http"
	"net/url"
)

func (endpoint *S_Endpoint) Mount(conf *config.S_Config) error {
	if endpoint.Request == nil {
		endpoint.Request = new(http.Request)
		*endpoint.Request = http.Request{}
	}

	urlParsed, err := url.Parse(endpoint.URL)
	if err != nil {
		return err
	}
	endpoint.Request.URL = urlParsed
	endpoint.Request.Method = *endpoint.Method
	endpoint.Request.Body = http.NoBody

	if endpoint.QueryParams != nil {
		query := urlParsed.Query()
		for k, v := range *endpoint.QueryParams {
			query.Add(k, v)
		}
		urlParsed.RawQuery = query.Encode()
	}

	if endpoint.Headers != nil {
		endpoint.Request.Header = http.Header{}
		for k, v := range *endpoint.Headers {
			endpoint.Request.Header.Add(k, v)
		}
	}

	return nil
}
