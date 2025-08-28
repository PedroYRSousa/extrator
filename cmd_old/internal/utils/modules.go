package utils

import (
	"fmt"
	"log"
	"time"
)

func ExecWithAttempts(maxAtt, delayInSeconds int, f func() error) error {
	var err error

	for att := 1; att <= maxAtt; att++ {
		if err = f(); err == nil {
			return nil
		}
		log.Printf("Attempt %d/%d | Waiting %d seconds to continue | %v", att, maxAtt, delayInSeconds, err)
		time.Sleep(time.Duration(delayInSeconds) * time.Second)
	}

	return fmt.Errorf("Max attempts reached | last error: %w", err)
}
