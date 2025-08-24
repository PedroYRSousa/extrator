package utils

import (
	"fmt"
	"log"
	"sync"
)

func AsyncWithLimit(maxConcurrent int, wg *sync.WaitGroup, f func()) {
	sem := make(chan struct{}, maxConcurrent)

	wg.Add(1)
	go func() {
		defer wg.Done()
		sem <- struct{}{}        // ocupa uma "vaga"
		defer func() { <-sem }() // libera a "vaga"

		f()
	}()
}

func Async(wg *sync.WaitGroup, f func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		f()
	}()
}

func ExecWithAttempts(maxAtt int, f func() error) error {
	var err error

	for att := 1; att <= maxAtt; att++ {
		if err = f(); err == nil {
			return nil
		}
		log.Printf("ERROR | Attempt %d/%d | %v", att, maxAtt, err)
	}

	return fmt.Errorf("max attempts reached | last error: %w", err)
}
