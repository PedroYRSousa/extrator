package utils

import (
	"fmt"
	"log"
	"regexp"
)

const baseRegex = `[0-9a-zA-Z-_\.]*`

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

func CheckIsHardCodedSecretOrEnv(target string) bool {
	return IsHardcoded(target) || IsEnv(target) || IsSecret(target)
}
