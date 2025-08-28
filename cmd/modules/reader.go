package modules

import (
	"os"

	"gopkg.in/yaml.v3"
)

func YamlToStruct(pathYaml string, dataEntry any) error {
	dataBytes, err := os.ReadFile(pathYaml)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(dataBytes, dataEntry)
	if err != nil {
		return err
	}

	return err
}
