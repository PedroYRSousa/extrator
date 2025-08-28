package config

import "extrator/modules"

func (configProducts *S_ConfigProducts) _transform() error {
	err := modules.TransformString(configProducts)
	if err != nil {
		return err
	}

	return nil
}

func (config *S_Config) transform() error {
	err := config.Products._transform()
	if err != nil {
		return err
	}

	err = modules.TransformString(config)
	if err != nil {
		return err
	}

	return nil
}
