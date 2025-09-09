package async

import "sync"

type S_AsyncLimiter struct {
	sem chan struct{}
	wg  *sync.WaitGroup
}
