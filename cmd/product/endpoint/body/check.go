package body

func (body *S_Body) Check() error {
	if body == nil {
		return nil
	}

	body = body.format()

	err := body.validate()
	if err != nil {
		return err
	}

	err = body.transform()
	if err != nil {
		return err
	}

	return nil
}
