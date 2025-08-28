package auth

import (
	"fmt"
	"net/http"
)

func (ab *S_Basic) mount(request *http.Request) error {
	request.SetBasicAuth(ab.Username, ab.Password)

	return nil
}

func (ab *S_Bearer) mount(request *http.Request) error {
	if *ab.Dynamic {
		newToken, err := ab.DynamicDetails.getToken()
		if err != nil {
			return err
		}
		ab.Token = *newToken
	}

	if request.Header == nil {
		request.Header = http.Header{}
	}

	request.Header.Add(*ab.Name, fmt.Sprintf("%s %s", *ab.Prefix, ab.Token))

	return nil
}

func (apk *S_ApiKey) mount(request *http.Request) error {
	switch *apk.Location {
	case AUTH_API_KEY_LOCATION_HEADER:
		if request.Header == nil {
			request.Header = http.Header{}
		}

		request.Header.Add(apk.Name, apk.Value)
	case AUTH_API_KEY_LOCATION_QUERY_PARAM:
		query := request.URL.Query()

		query.Add(apk.Name, apk.Value)

		request.URL.RawQuery = query.Encode()
	}

	return nil
}

func (c *S_Cookie) mount(request *http.Request) error {
	cookies, err := c.getCookies()
	if err != nil {
		return err
	}

	for _, cookie := range cookies {
		request.AddCookie(cookie)
	}

	return nil
}

func (a *S_Auth) Mount(request *http.Request) error {
	switch a.Mode {
	case AUTH_MODE_BASIC:
		return a.Basic.mount(request)
	case AUTH_MODE_BEARER:
		return a.Bearer.mount(request)
	case AUTH_MODE_API_KEY:
		return a.ApiKey.mount(request)
	case AUTH_MODE_COOKIE:
		return a.Cookie.mount(request)
	}

	return nil
}
