package env

import (
	"fmt"
	"os"
	"strings"
)

func Get(target string) (string, error) {
	_, after, _ := strings.Cut(target, "env(")
	key, _, _ := strings.Cut(after, ")")

	env, founded := os.LookupEnv(key)
	if !founded {
		return "", fmt.Errorf("environment %s not found", key)
	}

	return strings.ReplaceAll(target, fmt.Sprintf(ENV, key), env), nil
}
