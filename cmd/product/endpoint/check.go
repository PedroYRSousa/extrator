package endpoint

import "log"

func (endpoint *S_Endpoint) check() error {
	log.Printf("%+v | %+v | %+v | %+v | %+v", endpoint, endpoint.QueryParams, endpoint.Headers, endpoint.Body, endpoint.Body.File)

	endpoint = endpoint.format()

	err := endpoint.validate()
	if err != nil {
		return err
	}

	err = endpoint.transform()
	if err != nil {
		return err
	}

	err = endpoint.Body.Check()
	if err != nil {
		return err
	}

	log.Printf("%+v | %+v | %+v | %+v | %+v", endpoint, endpoint.QueryParams, endpoint.Headers, endpoint.Body, endpoint.Body.File)

	return nil
}
