package multipart

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func (mp *S_Multipart) Mount(request *http.Request) error {
	var body io.Reader
	var buf bytes.Buffer
	var contentLength int64

	writer := multipart.NewWriter(&buf)

	for index := range mp.Fields {
		for k, v := range mp.Fields[index] {
			err := writer.WriteField(k, v)
			if err != nil {
				return err
			}
		}
	}

	for _, mappingOfFiles := range mp.Files {
		for fieldName, path := range mappingOfFiles {
			part, err := writer.CreateFormFile(fieldName, filepath.Base(path))
			if err != nil {
				return err
			}

			fileBytes, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			_, err = part.Write(fileBytes)
			if err != nil {
				return err
			}
		}

	}

	writer.Close()
	body = &buf
	contentLength = int64(buf.Len())
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Body = io.NopCloser(body)
	request.ContentLength = contentLength

	return nil
}
