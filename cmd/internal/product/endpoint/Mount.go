package endpoint

import (
	"net/http"
	"net/url"
)

func (e *S_Endpoint) Mount(request *http.Request) error {
	uri, err := url.Parse(e.URI)
	if err != nil {
		return err
	}
	request.URL = uri

	err = e.EndpointConfig.Mount(request)
	if err != nil {
		return err
	}

	return nil
}
