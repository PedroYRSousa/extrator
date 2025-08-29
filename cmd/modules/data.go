package modules

import "sync"

type S_AsyncLimiter struct {
	sem chan struct{}
	wg  *sync.WaitGroup
}

const baseRegex = `[0-9a-zA-Z-_\.]*`
