package endpoint_config

import "io"

func (ec *S_EndpointConfig) Execute() error {
	if ec.Request == nil {
		panic("Internal error | endpointconfig.Request")
	}

	response, err := ec.Client.HttpClient.Do(ec.Request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	ec.Response = response
	ec.BodyResponse.Body = body

	return nil
}
