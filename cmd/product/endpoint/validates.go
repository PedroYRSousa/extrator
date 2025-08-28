package endpoint

import "fmt"

func (endpoint *S_Endpoint) validate() error {
	if endpoint.Name == "" {
		return fmt.Errorf("TODO, melhorar essa mensagem de erro")
	}

	if endpoint.Path == "" {
		return fmt.Errorf("TODO, melhorar essa mensagem de erro")
	}

	if endpoint.Description == "" {
		return fmt.Errorf("TODO, melhorar essa mensagem de erro")
	}

	if endpoint.Version == "" {
		return fmt.Errorf("TODO, melhorar essa mensagem de erro")
	}

	return nil
}
