package endpoint

import (
	"extrator/config"
	"extrator/modules"
	"os"
)

func Load(conf *config.S_Config, productPath string) ([]S_Endpoint, error) {
	endpoints := []S_Endpoint{}

	entries, err := os.ReadDir(productPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		endpoint := S_Endpoint{
			Name: entry.Name(),
			Path: modules.JoinPaths(productPath, entry.Name()),
		}

		err := modules.YamlToStruct(endpoint.Path, &endpoint)
		if err != nil {
			return nil, err
		}

		err = endpoint.check()
		if err != nil {
			return nil, err
		}

		endpoints = append(endpoints, endpoint)
	}

	return endpoints, nil
}
