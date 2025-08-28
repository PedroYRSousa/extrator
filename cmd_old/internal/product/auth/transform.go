package auth

import (
	"extrator/internal/utils"
	"extrator/pkg/env"
)

func (ab *S_Basic) transform() error {
	if utils.IsEnv(ab.Username) {
		newValue, err := env.Get(ab.Username)
		if err != nil {
			return err
		}
		ab.Username = newValue
	} else if utils.IsSecret(ab.Username) {
		// TODO, Pegar da aws secretmanager
	}

	if utils.IsEnv(ab.Password) {
		newValue, err := env.Get(ab.Password)
		if err != nil {
			return err
		}
		ab.Password = newValue
	} else if utils.IsSecret(ab.Password) {
		// TODO, Pegar da aws secretmanager
	}

	return nil
}

func (dd *S_BearerDynamicDetails) transform() error {
	err := dd.EndpointConfig.Transform()
	if err != nil {
		return err
	}

	return nil
}

func (ab *S_Bearer) transform() error {
	if utils.IsEnv(ab.Token) {
		newValue, err := env.Get(ab.Token)
		if err != nil {
			return err
		}
		ab.Token = newValue
	} else if utils.IsSecret(ab.Token) {
		// TODO, Pegar da aws secretmanager
	}

	err := ab.DynamicDetails.transform()
	if err != nil {
		return err
	}

	return nil
}

func (apk *S_ApiKey) transform() error {
	if utils.IsEnv(apk.Name) {
		newValue, err := env.Get(apk.Name)
		if err != nil {
			return err
		}
		apk.Name = newValue
	} else if utils.IsSecret(apk.Name) {
		// TODO, Pegar da aws secretmanager
	}

	if utils.IsEnv(apk.Value) {
		newValue, err := env.Get(apk.Value)
		if err != nil {
			return err
		}
		apk.Value = newValue
	} else if utils.IsSecret(apk.Value) {
		// TODO, Pegar da aws secretmanager
	}

	return nil
}

func (c *S_Cookie) transform() error {
	err := c.EndpointConfig.Transform()
	if err != nil {
		return err
	}

	return nil
}

func (auth *S_Auth) Transform() error {
	var transforms = map[authMode]func(*S_Auth) error{
		AUTH_MODE_BASIC:   func(_auth *S_Auth) error { return _auth.Basic.transform() },
		AUTH_MODE_BEARER:  func(_auth *S_Auth) error { return _auth.Bearer.transform() },
		AUTH_MODE_API_KEY: func(_auth *S_Auth) error { return _auth.ApiKey.transform() },
		AUTH_MODE_COOKIE:  func(_auth *S_Auth) error { return _auth.Cookie.transform() },
	}

	return transforms[auth.Mode](auth)
}
