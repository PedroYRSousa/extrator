package endpoint

import (
	"extrator/config"
	"io"
	"log"
	"net/http"
)

func (endpoint *S_Endpoint) Start(conf *config.S_Config) error {
	client := http.DefaultClient
	res, err := client.Do(endpoint.Request)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	log.Println(string(body))

	return nil
}
