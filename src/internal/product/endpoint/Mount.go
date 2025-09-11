package endpoint

import (
	"fmt"
	"net/http/httputil"
)

func (e *S_Endpoint) Mount() error {
	err := e.EndpointConfig.Mount()
	if err != nil {
		return err
	}

	if e.Auth != nil {
		err := e.Auth.Mount(e.EndpointConfig.Request)
		if err != nil {
			return err
		}
	}

	dumpRequest, err := httputil.DumpRequest(e.EndpointConfig.Request, true)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println(string(dumpRequest))
	fmt.Println()

	return nil
}
