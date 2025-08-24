package product

import (
	"errors"
	"extrator/internal/utils"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func (c *S_Product) validate() error {
	if (c.Name == "") || (strings.TrimSpace(c.Name) == "") {
		return errors.New("invalid config name | Check name | name cannot be empty")
	}

	if (c.Version == "") || (strings.TrimSpace(c.Version) == "") {
		return errors.New("invalid config version | Check version | version cannot be empty")
	}

	if (c.Description == "") || (strings.TrimSpace(c.Description) == "") {
		return errors.New("invalid config description | Check description | description cannot be empty")
	}

	err := c.Endpoint.Validate()
	if err != nil {
		return err
	}

	err = c.Auth.Validate()
	if err != nil {
		return err
	}

	for _, table := range c.Tables {
		err := table.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

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

func parseConfig(configFile s_ConfigFile) (map[string][]S_Product, error) {
	configs := make(map[string][]S_Product)

	for _, product := range configFile.products {
		configs[product.Name] = []S_Product{}

		for _, endpoint := range product.Endpoints {
			data, err := os.ReadFile(endpoint.Path)
			if err != nil {
				return nil, err
			}

			var config S_Product
			err = yaml.Unmarshal(data, &config)
			if err != nil {
				return nil, err
			}

			config.Path = utils.JoinPaths(utils.GetDataPaths(), product.Name)
			configs[product.Name] = append(configs[product.Name], config)
		}
	}

	return configs, nil
}

func validateConfig(configs map[string][]S_Product) error {
	for _, productConfigs := range configs {
		for _, config := range productConfigs {
			err := config.validate()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func Load() (map[string][]S_Product, error) {
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

	return config, nil
}
