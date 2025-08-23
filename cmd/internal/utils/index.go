package utils

import (
	"fmt"
	"log"
	"regexp"
)

func CheckIsHardCodedSecretOrEnv(target string) bool {
	var baseRegex = `[0-9a-zA-Z-_\.]*`

	hardCodedRegex, err := regexp.Compile(baseRegex)
	if err != nil {
		log.Fatalf("failed to compile regex for secret and env validation: %v", err)
		return false
	}

	secretRegex, err := regexp.Compile(fmt.Sprintf(`secret\(%s\)`, baseRegex))
	if err != nil {
		log.Fatalf("failed to compile regex for secret and env validation: %v", err)
		return false
	}

	envRegex, err := regexp.Compile(fmt.Sprintf(`env\(%s\)`, baseRegex))
	if err != nil {
		log.Fatalf("failed to compile regex for secret and env validation: %v", err)
		return false
	}

	return hardCodedRegex.Match([]byte(target)) || secretRegex.Match([]byte(target)) || envRegex.Match([]byte(target))
}
