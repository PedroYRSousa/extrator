package aws

import (
	"strings"
)

func Get(envName string) (string, error) {
	envName = strings.ReplaceAll(strings.ReplaceAll(envName, "secret(", ""), ")", "")

	panic("TODO, não implementado")

	return envName, nil
}
