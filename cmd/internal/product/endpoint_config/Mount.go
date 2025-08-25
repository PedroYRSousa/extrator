package endpoint_config

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

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
