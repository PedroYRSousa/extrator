package body

import (
	"bytes"
	"extrator/config"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func (body *S_Body) fromURLEncoded() ([]byte, string, error) {
	var data string = ""

	for k, v := range *body.FormFields {
		if data != "" {
			data = data + "&"
		}
		data = data + fmt.Sprintf("%s=%s", k, v)
	}

	return []byte(data), body.ContentType, nil
}

func (body *S_Body) fromFormData() ([]byte, string, error) {
	data := &bytes.Buffer{}
	writer := multipart.NewWriter(data)

	for k, v := range body.MultiPart.Fields {
		err := writer.WriteField(k, v)
		if err != nil {
			return nil, body.ContentType, err
		}
	}

	for k, v := range body.MultiPart.Files {
		for _, filePath := range v {
			file, err := os.Open(filePath)
			if err != nil {
				return nil, body.ContentType, err
			}
			defer file.Close()

			part, err := writer.CreateFormFile(k, filePath)
			if err != nil {
				return nil, body.ContentType, err
			}

			_, err = io.Copy(part, file)
			if err != nil {
				return nil, body.ContentType, err
			}
		}
	}

	err := writer.Close()
	if err != nil {
		return nil, body.ContentType, err
	}

	return data.Bytes(), writer.FormDataContentType(), nil
}

func (body *S_Body) fromGeneric() ([]byte, string, error) {
	var data []byte

	if body.File != nil {
		fileData, err := os.ReadFile(*body.File)
		if err != nil {
			return nil, body.ContentType, err
		}

		data = fileData
	} else {
		data = []byte(*body.Content)
	}

	return data, body.ContentType, nil
}

func (body *S_Body) Mount(conf *config.S_Config) ([]byte, string, error) {
	// if body.ContentType == "application/json" {
	// 	return body.fromJson()
	// }
	// if body.ContentType == "application/xml" {
	// 	return body.fromXML()
	// }

	if body.ContentType == "application/x-www-form-urlencoded" {
		return body.fromURLEncoded()
	}

	if body.ContentType == "multipart/form-data" {
		return body.fromFormData()
	}

	return body.fromGeneric()
}
