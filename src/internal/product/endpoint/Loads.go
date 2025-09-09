package endpoint

import (
	"errors"
	utilsstructs "extrator/pkg/utils_structs"
	"extrator/pkg/validator"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func (e *S_Endpoint) load() error {
	if e == nil {
		panic("Internal error | product.load")
	}

	fileData, err := os.ReadFile(e.Path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(fileData, e); err != nil {
		return err
	}

	if err := utilsstructs.Start(e); err != nil {
		return err
	}

	if err := validator.New().Struct(e); err != nil {
		return errors.New(validator.FormatErrors(err, e))
	}

	return nil
}

func Loads(pathProduct string) ([]*S_Endpoint, error) {
	endpoints := []*S_Endpoint{}

	entriesEndpoints, err := os.ReadDir(pathProduct)
	if err != nil {
		return nil, err
	}

	for _, entryEndpoint := range entriesEndpoints {
		if entryEndpoint.IsDir() {
			continue
		}

		pathEndpoint := fmt.Sprintf("%s/%s", pathProduct, entryEndpoint.Name())
		_endpoint := New(entryEndpoint.Name(), pathEndpoint)
		err := _endpoint.load()
		if err != nil {
			return nil, err
		}

		endpoints = append(endpoints, _endpoint)
	}

	return endpoints, nil
}
