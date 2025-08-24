package auth

import (
	"strings"
	"testing"
)

func TestS_Basic_format(t *testing.T) {
	b := &S_Basic{
		Username: " user ",
		Password: " pass ",
	}
	b.format()
	if b.Username != "user" {
		t.Errorf("expected Username to be 'user', got '%s'", b.Username)
	}
	if b.Password != "pass" {
		t.Errorf("expected Password to be 'pass', got '%s'", b.Password)
	}
}

func TestS_Basic_format_nil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on nil receiver")
		}
	}()
	var b *S_Basic
	b.format()
}

func TestS_DynamicDetails_format(t *testing.T) {
	dd := &S_DynamicDetails{
		Endpoint: " endpoint ",
		Extract:  " extract ",
	}
	dd.format()
	if dd.Endpoint != "endpoint" {
		t.Errorf("expected Endpoint to be 'endpoint', got '%s'", dd.Endpoint)
	}
	if dd.Extract != "extract" {
		t.Errorf("expected Extract to be 'extract', got '%s'", dd.Extract)
	}
}

func TestS_DynamicDetails_format_nil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on nil receiver")
		}
	}()
	var dd *S_DynamicDetails
	dd.format()
}

func TestS_Bearer_format(t *testing.T) {
	b := &S_Bearer{
		Token:   " token ",
		Dynamic: nil,
		Name:    nil,
		Prefix:  nil,
	}
	b.format()
	if b.Token != "token" {
		t.Errorf("expected Token to be 'token', got '%s'", b.Token)
	}
	if b.Dynamic == nil || *b.Dynamic != AUTH_BEARER_DYNAMIC_DEFAULT {
		t.Errorf("expected Dynamic to be set to default")
	}
	if b.Name == nil || *b.Name != AUTH_BEARER_NAME_DEFAULT {
		t.Errorf("expected Name to be set to default")
	}
	if b.Prefix == nil || *b.Prefix != AUTH_BEARER_PREFIX_DEFAULT {
		t.Errorf("expected Prefix to be set to default")
	}

	// Test trimming when not nil
	name := " name "
	prefix := " prefix "
	b2 := &S_Bearer{
		Token:   " token2 ",
		Dynamic: new(bool),
		Name:    &name,
		Prefix:  &prefix,
	}
	b2.format()
	if *b2.Name != "name" {
		t.Errorf("expected Name to be 'name', got '%s'", *b2.Name)
	}
	if *b2.Prefix != "prefix" {
		t.Errorf("expected Prefix to be 'prefix', got '%s'", *b2.Prefix)
	}
}

func TestS_Bearer_format_nil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on nil receiver")
		}
	}()
	var b *S_Bearer
	b.format()
}

func TestS_ApiKey_format(t *testing.T) {
	apk := &S_ApiKey{
		Name:     " name ",
		Value:    " value ",
		Location: nil,
	}
	apk.format()
	if apk.Name != "name" {
		t.Errorf("expected Name to be 'name', got '%s'", apk.Name)
	}
	if apk.Value != "value" {
		t.Errorf("expected Value to be 'value', got '%s'", apk.Value)
	}
	if apk.Location == nil || *apk.Location != AUTH_API_KEY_LOCATION_DEFAULT {
		t.Errorf("expected Location to be set to default")
	}

	// Test trimming when not nil
	loc := authApiKeyLocation(" header ")
	apk2 := &S_ApiKey{
		Name:     " name2 ",
		Value:    " value2 ",
		Location: &loc,
	}
	apk2.format()
	if string(*apk2.Location) != "header" {
		t.Errorf("expected Location to be 'header', got '%s'", *apk2.Location)
	}
}

func TestS_ApiKey_format_nil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on nil receiver")
		}
	}()
	var apk *S_ApiKey
	apk.format()
}

func TestS_Cookie_format(t *testing.T) {
	c := &S_Cookie{
		Endpoint: " endpoint ",
		Extract:  []string{" a ", " b "},
	}
	c.format()
	if c.Endpoint != "endpoint" {
		t.Errorf("expected Endpoint to be 'endpoint', got '%s'", c.Endpoint)
	}
	for i, v := range c.Extract {
		if v != strings.TrimSpace(v) {
			t.Errorf("expected Extract[%d] to be trimmed, got '%s'", i, v)
		}
	}
	if c.Extract[0] != "a" || c.Extract[1] != "b" {
		t.Errorf("expected Extract to be ['a','b'], got %v", c.Extract)
	}
}

func TestS_Cookie_format_nil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on nil receiver")
		}
	}()
	var c *S_Cookie
	c.format()
}

func TestS_Auth_format(t *testing.T) {
	a := &S_Auth{
		Mode: authMode("BASIC"),
	}
	a.format()
	if a.Mode != authMode("basic") {
		t.Errorf("expected Mode to be 'basic', got '%s'", a.Mode)
	}
}

func TestS_Auth_format_nil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on nil receiver")
		}
	}()
	var a *S_Auth
	a.format()
}
