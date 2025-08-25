package product

import (
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func readEndpoints(product s_ProductConfigFile) (s_ProductConfigFile, error) {
	endpointDirs, err := os.ReadDir(product.Path)
	if err != nil {
		return product, err
	}

	for _, endpointDir := range endpointDirs {
		if endpointDir.IsDir() {
			continue
		}

		product.Endpoints = append(product.Endpoints, s_EndpointConfigFile{
			Name: strings.ReplaceAll(endpointDir.Name(), ".yaml", ""),
			Path: product.Path + "/" + endpointDir.Name(),
		})
	}

	return product, nil
}

func loadConfigFile() (s_ConfigFile, error) {
	log.Println("Read file configuration")

	config := s_ConfigFile{
		products: []s_ProductConfigFile{},
	}

	dirs, err := os.ReadDir(configPath)
	if err != nil {
		return config, err
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			product, err := readEndpoints(s_ProductConfigFile{
				Name:      dir.Name(),
				Path:      configPath + "/" + dir.Name(),
				Endpoints: []s_EndpointConfigFile{},
			})
			if err != nil {
				return config, err
			}

			config.products = append(config.products, product)
		}
	}

	return config, nil
}

func parseConfig(configFile s_ConfigFile) (map[string][]S_ProductEndpoint, error) {
	products := make(map[string][]S_ProductEndpoint)

	for _, product := range configFile.products {
		products[product.Name] = []S_ProductEndpoint{}

		for _, endpoint := range product.Endpoints {
			data, err := os.ReadFile(endpoint.Path)
			if err != nil {
				return nil, err
			}

			var productEndpoint S_ProductEndpoint
			err = yaml.Unmarshal(data, &productEndpoint)
			if err != nil {
				return nil, err
			}

			productEndpoint.ProductName = product.Name
			products[product.Name] = append(products[product.Name], productEndpoint)
		}
	}

	return products, nil
}

func validateConfig(configs map[string][]S_ProductEndpoint) error {
	log.Println("Validate configuration")

	for _, productConfigs := range configs {
		for index := range productConfigs {
			err := productConfigs[index].validate()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func transformConfig(configs map[string][]S_ProductEndpoint) error {
	for _, productConfigs := range configs {
		for index := range productConfigs {
			err := productConfigs[index].transform()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func Load() (map[string][]S_ProductEndpoint, error) {
	log.Println("Loading configuration")

	s_ConfigFile, err := loadConfigFile()
	if err != nil {
		return nil, err
	}

	config, err := parseConfig(s_ConfigFile)
	if err != nil {
		return nil, err
	}

	err = validateConfig(config)
	if err != nil {
		return nil, err
	}

	err = transformConfig(config)
	if err != nil {
		return nil, err
	}

	log.Printf("Configuration Read")

	return config, nil
}
