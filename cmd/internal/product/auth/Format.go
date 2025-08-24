package auth

import (
	"strings"
)

func (ab *S_Basic) format() {
	if ab == nil {
		panic("TODO, Dizer aqui que é obrigatório")
	}

	ab.Username = strings.TrimSpace(ab.Username)
	ab.Password = strings.TrimSpace(ab.Password)
}

func (dd *S_DynamicDetails) format() {
	if dd == nil {
		panic("TODO, Dizer aqui que é obrigatório")
	}

	dd.Endpoint = strings.TrimSpace(dd.Endpoint)
	dd.Extract = strings.TrimSpace(dd.Extract)
}

func (ab *S_Bearer) format() {
	if ab == nil {
		panic("TODO, Dizer aqui que é obrigatório")
	}

	ab.Token = strings.TrimSpace(ab.Token)

	if ab.Dynamic == nil {
		ab.Dynamic = new(bool)
		*ab.Dynamic = AUTH_BEARER_DYNAMIC_DEFAULT
	}

	if ab.Name == nil {
		ab.Name = new(string)
		*ab.Name = AUTH_BEARER_NAME_DEFAULT
	} else {
		*ab.Name = strings.TrimSpace(*ab.Name)
	}

	if ab.Prefix == nil {
		ab.Prefix = new(string)
		*ab.Prefix = AUTH_BEARER_PREFIX_DEFAULT
	} else {
		*ab.Prefix = strings.TrimSpace(*ab.Prefix)
	}
}

func (apk *S_ApiKey) format() {
	if apk == nil {
		panic("TODO, Dizer aqui que é obrigatório")
	}

	apk.Name = strings.TrimSpace(apk.Name)
	apk.Value = strings.TrimSpace(apk.Value)

	if apk.Location == nil {
		apk.Location = new(authApiKeyLocation)
		*apk.Location = AUTH_API_KEY_LOCATION_DEFAULT
	} else {
		*apk.Location = authApiKeyLocation(strings.TrimSpace(string(*apk.Location)))
	}
}

func (ab *S_Cookie) format() {
	if ab == nil {
		panic("TODO, Dizer aqui que é obrigatório")
	}
	ab.Endpoint = strings.TrimSpace(ab.Endpoint)
	for index := range ab.Extract {
		ab.Extract[index] = strings.TrimSpace(ab.Extract[index])
	}
}

func (a *S_Auth) format() {
	if a == nil {
		panic("TODO, Dizer aqui que é obrigatório")
	}

	a.Mode = authMode(strings.ToLower(string(a.Mode)))
}
