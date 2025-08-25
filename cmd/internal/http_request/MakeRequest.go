package http_request

import (
	"extrator/internal/product/endpoint_config"
	"net/http"
	"slices"
)

func MakeRequest(request *http.Request, endpointConfig endpoint_config.S_EndpointConfig) (*http.Response, error) {
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	if slices.Contains(*endpointConfig.StatusCodeWaiting, uint(res.StatusCode)) {
		return nil, ErrStatusCodeWaiting
	} else if slices.Contains(*endpointConfig.StatusCodeSkip, uint(res.StatusCode)) {
		return nil, ErrStatusCodeSkip
	} else if !slices.Contains(*endpointConfig.StatusCodeSuccess, uint(res.StatusCode)) {
		return nil, ErrStatusCodeError
	}

	return res, nil
}
