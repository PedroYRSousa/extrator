package utils

import (
	"fmt"
	"log"
	"regexp"
)

func IsSecret(target string) bool {
	secretRegex, err := regexp.Compile(fmt.Sprintf(`^secret\(%s\)$`, baseRegex))
	if err != nil {
		log.Fatalf("failed to compile regex for secret and env validation: %v", err)
		return false
	}

	return secretRegex.Match([]byte(target))
}

func IsEnv(target string) bool {
	envRegex, err := regexp.Compile(fmt.Sprintf(`^env\(%s\)$`, baseRegex))
	if err != nil {
		log.Fatalf("failed to compile regex for secret and env validation: %v", err)
		return false
	}

	return envRegex.Match([]byte(target))
}

func IsHardcoded(target string) bool {
	hardCodedRegex, err := regexp.Compile(baseRegex)
	if err != nil {
		log.Fatalf("failed to compile regex for secret and env validation: %v", err)
		return false
	}

	return hardCodedRegex.Match([]byte(target))
}
