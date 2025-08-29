package endpoint

import (
	"bytes"
	"extrator/config"
	"io"
	"net/http"
	"net/url"
	"strconv"
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

	if endpoint.QueryParams != nil {
		query := urlParsed.Query()
		for k, v := range *endpoint.QueryParams {
			query.Add(k, v)
		}
		urlParsed.RawQuery = query.Encode()
	}

	endpoint.Request.Header = http.Header{}
	if endpoint.Headers != nil {
		for k, v := range *endpoint.Headers {
			endpoint.Request.Header.Add(k, v)
		}
	}

	if endpoint.Body != nil {
		bodyBytes, contentType, err := endpoint.Body.Mount(conf)
		if err != nil {
			return err
		}

		endpoint.Request.Header.Set("content-type", contentType)
		endpoint.Request.Header.Set("content-length", strconv.Itoa(len(bodyBytes)))
		endpoint.Request.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	} else {
		endpoint.Request.Body = http.NoBody
	}

	return nil
}
