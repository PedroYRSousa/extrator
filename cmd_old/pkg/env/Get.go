package env

import (
	"fmt"
	"os"
	"strings"
)

func Get(envName string) (string, error) {
	envName = strings.ReplaceAll(strings.ReplaceAll(envName, "env(", ""), ")", "")

	env, founded := os.LookupEnv(envName)
	if !founded {
		return "", fmt.Errorf("environment %s not found", envName)
	}

	return env, nil
}
