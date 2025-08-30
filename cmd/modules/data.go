package modules

import "sync"

const (
	BASE_REGEX = `[0-9a-zA-Z-_\.]*`
)

type S_AsyncLimiter struct {
	sem chan struct{}
	wg  *sync.WaitGroup
}
