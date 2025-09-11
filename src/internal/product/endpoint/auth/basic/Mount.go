package basic

import (
	"net/http"
)

func (b *S_Basic) Mount(request *http.Request) error {
	request.SetBasicAuth(b.Username, b.Password)

	return nil
}
