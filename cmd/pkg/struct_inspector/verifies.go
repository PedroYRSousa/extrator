package struct_inspector

import (
	"fmt"
	"log"
	"regexp"
)

func isSecret(target string) bool {
	secretRegex, err := regexp.Compile(fmt.Sprintf(`secret\(%s\)`, BASE_REGEX))
	if err != nil {
		log.Fatalf("failed to compile regex for secret and env validation: %v", err)
		return false
	}

	return secretRegex.Match([]byte(target))
}

func isEnv(target string) bool {
	envRegex, err := regexp.Compile(fmt.Sprintf(`env\(%s\)`, BASE_REGEX))
	if err != nil {
		log.Fatalf("failed to compile regex for secret and env validation: %v", err)
		return false
	}

	return envRegex.Match([]byte(target))
}
