package body

import (
	"encoding/json"
	"extrator/config"
	"fmt"
)

func (body *S_Body) fromJson() ([]byte, error) {
	var i interface{}
	err := json.Unmarshal([]byte(*body.Content), &i)
	if err != nil {
		return nil, err
	}

	if body.Content != nil {
		return []byte(*body.Content), nil
	}
	return nil, fmt.Errorf("no content for JSON body")
}

func (body *S_Body) Mount(conf *config.S_Config) ([]byte, error) {
	if body.ContentType == "application/json" {
		return body.fromJson()
	}

	return []byte{}, nil
}
