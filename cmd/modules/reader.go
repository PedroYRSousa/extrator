package modules

import (
	"encoding/json"
	"encoding/xml"
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

func JsonToStruct(pathJson string, dataEntry any) error {
	dataBytes, err := os.ReadFile(pathJson)
	if err != nil {
		return err
	}

	err = json.Unmarshal(dataBytes, dataEntry)
	if err != nil {
		return err
	}

	return err
}

func XMLToStruct(pathXML string, dataEntry any) error {
	dataBytes, err := os.ReadFile(pathXML)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(dataBytes, dataEntry)
	if err != nil {
		return err
	}

	return err
}
