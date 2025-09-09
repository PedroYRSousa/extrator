package body

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func (b *S_Body) Mount(request *http.Request) error {
	var body io.Reader
	var contentLength int64

	if b.File != nil {
		fileBytes, err := os.ReadFile(*b.File)
		if err != nil {
			return err
		}
		body = bytes.NewReader(fileBytes)
		contentLength = int64(len(fileBytes))
	}

	if b.FormURLEncoded != nil {
		data := url.Values{}
		for index := range *b.FormURLEncoded {
			for k, v := range (*b.FormURLEncoded)[index] {
				data.Set(k, v)
			}
		}
		encoded := data.Encode()
		body = strings.NewReader(encoded)
		contentLength = int64(len(encoded))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	if b.MultiPart != nil {
		err := b.MultiPart.Mount(request)
		if err != nil {
			return err
		}

		return nil
	}

	if b.Text != nil {
		txt := *b.Text
		body = strings.NewReader(txt)
		contentLength = int64(len(txt))
	}

	request.Body = io.NopCloser(body)
	request.ContentLength = contentLength

	return nil
}
