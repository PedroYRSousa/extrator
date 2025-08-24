package auth

import (
	endpointconfig "extrator/internal/product/endpoint_config"
	"reflect"
	"testing"
)

func TestS_BasicFields(t *testing.T) {
	b := S_Basic{
		Username: "user",
		Password: "pass",
	}
	if b.Username != "user" {
		t.Errorf("expected Username to be 'user', got '%s'", b.Username)
	}
	if b.Password != "pass" {
		t.Errorf("expected Password to be 'pass', got '%s'", b.Password)
	}
}

func TestS_DynamicDetailsFields(t *testing.T) {
	ec := &endpointconfig.S_EndpointConfig{}
	d := S_DynamicDetails{
		Endpoint:       "/token",
		Extract:        "access_token",
		EndpointConfig: ec,
	}
	if d.Endpoint != "/token" {
		t.Errorf("expected Endpoint to be '/token', got '%s'", d.Endpoint)
	}
	if d.Extract != "access_token" {
		t.Errorf("expected Extract to be 'access_token', got '%s'", d.Extract)
	}
	if d.EndpointConfig != ec {
		t.Errorf("expected EndpointConfig to match")
	}
}

func TestS_BearerFields(t *testing.T) {
	dyn := true
	name := "Auth"
	prefix := "Bearer"
	details := &S_DynamicDetails{Endpoint: "/token", Extract: "token"}
	b := S_Bearer{
		Token:          "tok",
		Dynamic:        &dyn,
		Name:           &name,
		Prefix:         &prefix,
		DynamicDetails: details,
	}
	if b.Token != "tok" {
		t.Errorf("expected Token to be 'tok', got '%s'", b.Token)
	}
	if b.Dynamic == nil || *b.Dynamic != true {
		t.Errorf("expected Dynamic to be true")
	}
	if b.Name == nil || *b.Name != "Auth" {
		t.Errorf("expected Name to be 'Auth'")
	}
	if b.Prefix == nil || *b.Prefix != "Bearer" {
		t.Errorf("expected Prefix to be 'Bearer'")
	}
	if b.DynamicDetails != details {
		t.Errorf("expected DynamicDetails to match")
	}
}

func TestS_ApiKeyFields(t *testing.T) {
	loc := AUTH_API_KEY_LOCATION_QUERY_PARAM
	a := S_ApiKey{
		Name:     "api_key",
		Value:    "123",
		Location: &loc,
	}
	if a.Name != "api_key" {
		t.Errorf("expected Name to be 'api_key', got '%s'", a.Name)
	}
	if a.Value != "123" {
		t.Errorf("expected Value to be '123', got '%s'", a.Value)
	}
	if a.Location == nil || *a.Location != AUTH_API_KEY_LOCATION_QUERY_PARAM {
		t.Errorf("expected Location to be AUTH_API_KEY_LOCATION_QUERY_PARAM")
	}
}

func TestS_CookieFields(t *testing.T) {
	ec := &endpointconfig.S_EndpointConfig{}
	c := S_Cookie{
		Endpoint:       "/login",
		Extract:        []string{"sessionid"},
		EndpointConfig: ec,
	}
	if c.Endpoint != "/login" {
		t.Errorf("expected Endpoint to be '/login', got '%s'", c.Endpoint)
	}
	if !reflect.DeepEqual(c.Extract, []string{"sessionid"}) {
		t.Errorf("expected Extract to be ['sessionid']")
	}
	if c.EndpointConfig != ec {
		t.Errorf("expected EndpointConfig to match")
	}
}

func TestS_AuthFields(t *testing.T) {
	basic := &S_Basic{Username: "u", Password: "p"}
	bearer := &S_Bearer{Token: "t"}
	apiKey := &S_ApiKey{Name: "n", Value: "v"}
	cookie := &S_Cookie{Endpoint: "/e", Extract: []string{"c"}}
	a := S_Auth{
		Mode:   AUTH_MODE_BASIC,
		Basic:  basic,
		Bearer: bearer,
		ApiKey: apiKey,
		Cookie: cookie,
	}
	if a.Mode != AUTH_MODE_BASIC {
		t.Errorf("expected Mode to be AUTH_MODE_BASIC")
	}
	if a.Basic != basic {
		t.Errorf("expected Basic to match")
	}
	if a.Bearer != bearer {
		t.Errorf("expected Bearer to match")
	}
	if a.ApiKey != apiKey {
		t.Errorf("expected ApiKey to match")
	}
	if a.Cookie != cookie {
		t.Errorf("expected Cookie to match")
	}
}

func TestConstantsAndVars(t *testing.T) {
	if AUTH_MODE_NONE != "none" {
		t.Errorf("expected AUTH_MODE_NONE to be 'none'")
	}
	if AUTH_API_KEY_LOCATION_HEADER != "header" {
		t.Errorf("expected AUTH_API_KEY_LOCATION_HEADER to be 'header'")
	}
	if AUTH_BEARER_NAME_DEFAULT != "Authorization" {
		t.Errorf("expected AUTH_BEARER_NAME_DEFAULT to be 'Authorization'")
	}
	if !containsAuthMode(MODES_AVAILABLE, AUTH_MODE_BASIC) {
		t.Errorf("MODES_AVAILABLE should contain AUTH_MODE_BASIC")
	}
	if !containsApiKeyLocation(AVAILABLE_LOCATIONS, AUTH_API_KEY_LOCATION_QUERY_PARAM) {
		t.Errorf("AVAILABLE_LOCATIONS should contain AUTH_API_KEY_LOCATION_QUERY_PARAM")
	}
}

func containsAuthMode(modes []authMode, mode authMode) bool {
	for _, m := range modes {
		if m == mode {
			return true
		}
	}
	return false
}

func containsApiKeyLocation(locs []authApiKeyLocation, loc authApiKeyLocation) bool {
	for _, l := range locs {
		if l == loc {
			return true
		}
	}
	return false
}
