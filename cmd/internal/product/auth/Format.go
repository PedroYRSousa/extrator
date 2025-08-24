package auth

import (
	"strings"
)

func (a *S_Auth) Format() {
	a.AuthMode = strings.ToLower(a.AuthMode)

	if a.AuthApiKey != nil {
		a.AuthApiKey.format()
	}
	if a.AuthBasic != nil {
		a.AuthBasic.format()
	}
	if a.AuthBearer != nil {
		a.AuthBearer.format()
	}
	if a.AuthBearerDynamic != nil {
		a.AuthBearerDynamic.format()
	}
	if a.AuthCookie != nil {
		a.AuthCookie.format()
	}
}

func (ab *S_AuthBasic) format() {
	ab.Username = strings.TrimSpace(ab.Username)
	ab.Password = strings.TrimSpace(ab.Password)
}

func (ab *S_AuthBearer) format() {
	ab.Name = strings.TrimSpace(ab.Name)
	ab.Prefix = strings.TrimSpace(ab.Prefix)
	ab.Token = strings.TrimSpace(ab.Token)
}

func (apk *S_AuthApiKey) format() {
	apk.Name = strings.TrimSpace(apk.Name)
	apk.Value = strings.TrimSpace(apk.Value)
}

func (abd *S_AuthBearerDynamic) format() {
	abd.Name = strings.TrimSpace(abd.Name)
	abd.Prefix = strings.TrimSpace(abd.Prefix)
	abd.Endpoint = strings.TrimSpace(abd.Endpoint)
	abd.Extract = strings.TrimSpace(abd.Extract)
}

func (ab *S_AuthCookie) format() {
	ab.Endpoint = strings.TrimSpace(ab.Endpoint)
	ab.Method = strings.ToUpper(ab.Method)
}
