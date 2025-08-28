package auth

import (
	"extrator/internal/utils"
	"net/http"
	"net/url"
)

func (c *S_Cookie) getCookies() ([]*http.Cookie, error) {
	request := http.Request{}

	uri, err := url.Parse(c.Endpoint)
	if err != nil {
		return nil, err
	}
	request.URL = uri

	err = c.EndpointConfig.Mount(&request)
	if err != nil {
		return nil, err
	}

	cookies := []*http.Cookie{}
	err = utils.ExecWithAttempts(int(*c.EndpointConfig.Retry.Attempts), int(*c.EndpointConfig.Retry.DelayInSeconds), func() error {
		res, err := http.DefaultClient.Do(&request)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		cookies = res.Cookies()

		return nil
	})
	if err != nil {
		return nil, err
	}

	return cookies, nil
}
