package utilsstructs

import (
	"extrator/pkg/env"
	"extrator/pkg/secret"
	"fmt"
	"log"
	"regexp"
)

func is(target, expr string) bool {
	secretRegex, err := regexp.Compile(expr)
	if err != nil {
		log.Fatalf("failed to compile regex for secret and env validation: %v", err)
		return false
	}

	return secretRegex.Match([]byte(target))
}

func isSecret(target string) bool {
	return is(target, fmt.Sprintf(secret.SECRET, BASE_REGEX))
}

func isEnv(target string) bool {
	return is(target, fmt.Sprintf(env.ENV, BASE_REGEX))
}
