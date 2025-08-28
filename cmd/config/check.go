package config

func (config *S_Config) check() error {
	config = config.format()

	err := config.validate()
	if err != nil {
		return err
	}

	err = config.transform()
	if err != nil {
		return err
	}

	return nil
}
