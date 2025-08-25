package env

import (
	"fmt"
	"os"
)

func Get(envName string) (string, error) {
	env, founded := os.LookupEnv(envName)
	if !founded {
		return "", fmt.Errorf("environment %s not found", envName)
	}

	return env, nil
}
