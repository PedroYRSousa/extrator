package config

import "fmt"

func (configProducts *S_ConfigProducts) _validate() error {
	if configProducts.Path == "" {
		return fmt.Errorf("TODO, melhorar essa mensagem de erro")
	}

	return nil
}

func (config *S_Config) validate() error {
	err := config.Products._validate()
	if err != nil {
		return err
	}

	if config.Products == nil {
		return fmt.Errorf("TODO, melhorar essa mensagem de erro")
	}

	return nil
}
