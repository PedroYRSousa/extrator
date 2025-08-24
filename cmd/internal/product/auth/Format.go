package auth

import (
	"strings"
)

func (a *S_Auth) Format() {
	a.Mode = strings.ToLower(a.Mode)

	if a.ApiKey != nil {
		a.ApiKey.format()
	}
	if a.Basic != nil {
		a.Basic.format()
	}
	if a.Bearer != nil {
		a.Bearer.format()
	}
	if a.BearerDynamic != nil {
		a.BearerDynamic.format()
	}
	if a.Cookie != nil {
		a.Cookie.format()
	}
}

func (ab *S_Basic) format() {
	ab.Username = strings.TrimSpace(ab.Username)
	ab.Password = strings.TrimSpace(ab.Password)
}

func (ab *S_Bearer) format() {
	ab.Name = strings.TrimSpace(ab.Name)
	ab.Prefix = strings.TrimSpace(ab.Prefix)
	ab.Token = strings.TrimSpace(ab.Token)
}

func (apk *S_ApiKey) format() {
	apk.Name = strings.TrimSpace(apk.Name)
	apk.Value = strings.TrimSpace(apk.Value)
}

func (abd *S_BearerDynamic) format() {
	abd.Name = strings.TrimSpace(abd.Name)
	abd.Prefix = strings.TrimSpace(abd.Prefix)
	abd.Endpoint = strings.TrimSpace(abd.Endpoint)
	abd.Extract = strings.TrimSpace(abd.Extract)
}

func (ab *S_Cookie) format() {
	ab.Endpoint = strings.TrimSpace(ab.Endpoint)
	ab.Method = strings.ToUpper(ab.Method)
}
