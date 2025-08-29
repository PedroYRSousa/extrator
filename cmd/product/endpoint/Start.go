package endpoint

import (
	"extrator/config"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

func (endpoint *S_Endpoint) Start(conf *config.S_Config) error {
	log.Println(endpoint.Headers)
	log.Println(endpoint.QueryParams)

	dump, err := httputil.DumpRequest(endpoint.Request, true)
	fmt.Println("====================")
	fmt.Println(string(dump), err)
	fmt.Println("====================")

	client := http.DefaultClient
	res, err := client.Do(endpoint.Request)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// log.Println(string(body))

	return nil
}
