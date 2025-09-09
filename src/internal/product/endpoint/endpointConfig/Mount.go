package endpointconfig

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func (ec *S_EndpointConfig) Mount() error {
	ec.Request = nil
	ec.Response = nil

	request, err := http.NewRequest(ec.Method, ec.URL, nil)
	if err != nil {
		return err
	}
	request.Body = http.NoBody
	request.Header = http.Header{}

	if ec.Headers != nil {
		for index := range *ec.Headers {
			for k, v := range (*ec.Headers)[index] {
				request.Header.Add(k, v)
			}
		}
	}

	if ec.QueryParams != nil {
		query := request.URL.Query()

		for index := range *ec.QueryParams {
			for k, v := range (*ec.QueryParams)[index] {
				query.Add(k, v)
			}
		}

		request.URL.RawQuery = query.Encode()
	}

	if ec.Body != nil {
		err = ec.Body.Mount(request)
		if err != nil {
			return err
		}
	}

	ec.Request = request

	dumpRequest, err := httputil.DumpRequest(ec.Request, true)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println(string(dumpRequest))
	fmt.Println()

	return nil
}
