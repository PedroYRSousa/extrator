package bearer

import (
	"fmt"
	"net/http"
)

func (b *S_Bearer) Mount(request *http.Request) error {
	var token = b.Token

	if b.IsDynamic {
		err := b.EndpointConfig.Mount()
		if err != nil {
			return err
		}
	}

	// Monta header
	prefix := b.Prefix
	if prefix == "" {
		prefix = "Bearer"
	}
	request.Header.Set(b.Name, fmt.Sprintf("%s %s", prefix, (*token)))

	return nil
}
