package endpoint

import (
	"encoding/json"
	"encoding/xml"
	"extrator/config"
	"extrator/modules"
	"log"
)

func (endpoint *S_Endpoint) Start(conf *config.S_Config) error {
	res, body, err := modules.MakeRequest(endpoint.Request)
	if err != nil {
		return err
	}
	log.Println(res)
	log.Println("============")

	var bodyData interface{}
	if *endpoint.ResponseFormat == RESPONSE_FORMAT_JSON {
		err = json.Unmarshal(body, &bodyData)
		if err != nil {
			return err
		}
	} else if *endpoint.ResponseFormat == RESPONSE_FORMAT_XML {
		err = xml.Unmarshal(body, &bodyData)
		if err != nil {
			return err
		}
	}

	return nil
}
