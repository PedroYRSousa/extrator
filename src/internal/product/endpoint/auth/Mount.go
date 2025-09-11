package auth

import "net/http"

func (a *S_Auth) Mount(request *http.Request) error {
	if a.ApiKey != nil {
		return a.ApiKey.Mount(request)
	}

	if a.Basic != nil {
		return a.Basic.Mount(request)
	}

	if a.Bearer != nil {
		return a.Bearer.Mount(request)
	}

	return nil
}
