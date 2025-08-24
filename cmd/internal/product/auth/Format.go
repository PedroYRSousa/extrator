package auth

import (
	"strings"
)

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

func (ab *S_Cookie) format() {
	ab.Endpoint = strings.TrimSpace(ab.Endpoint)
	ab.Method = strings.TrimSpace(strings.ToUpper(ab.Method))
	for index := range ab.Extract {
		ab.Extract[index] = strings.TrimSpace(ab.Extract[index])
	}
}

func (abd *S_BearerDynamic) format() {
	abd.Name = strings.TrimSpace(abd.Name)
	abd.Prefix = strings.TrimSpace(abd.Prefix)
	abd.Endpoint = strings.TrimSpace(abd.Endpoint)
	abd.Method = strings.TrimSpace(strings.ToUpper(abd.Method))
	abd.Extract = strings.TrimSpace(abd.Extract)
}

func (a *S_Auth) format() {
	a.Mode = authMode(strings.ToLower(string(a.Mode)))
}
