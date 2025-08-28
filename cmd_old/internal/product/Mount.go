package product

import "net/http"

func (p *S_ProductEndpoint) Mount() (*http.Request, error) {
	request := http.Request{}

	err := p.Endpoint.Mount(&request)
	if err != nil {
		return nil, err
	}

	err = p.Auth.Mount(&request)
	if err != nil {
		return nil, err
	}

	return &request, nil
}
