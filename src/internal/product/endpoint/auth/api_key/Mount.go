package api_key

import (
	"fmt"
	"net/http"
	"strings"
)

func (ak *S_ApiKey) Mount(request *http.Request) error {
	switch strings.ToLower(ak.Location) {
	case "header":
		request.Header.Set(ak.Key, ak.Value)

	case "query":
		query := request.URL.Query()
		query.Set(ak.Key, ak.Value)
		request.URL.RawQuery = query.Encode()

	case "cookie":
		cookie := &http.Cookie{
			Name:  ak.Key,
			Value: ak.Value,
		}
		request.AddCookie(cookie)

	default:
		return fmt.Errorf("invalid api key location: %s", ak.Location)
	}

	return nil
}
