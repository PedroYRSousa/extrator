package modules

import (
	"io"
	"net/http"
)

func MakeRequest(request *http.Request) (*http.Response, []byte, error) {
	client := http.DefaultClient
	res, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	return res, body, nil
}
