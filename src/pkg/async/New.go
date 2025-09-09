package async

import "sync"

func AsyncLimiter(maxConcurrent int) *S_AsyncLimiter {
	return &S_AsyncLimiter{
		sem: make(chan struct{}, maxConcurrent),
		wg:  &sync.WaitGroup{},
	}
}
