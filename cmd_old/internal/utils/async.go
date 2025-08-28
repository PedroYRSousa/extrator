package utils

import "sync"

type S_AsyncLimiter struct {
	sem chan struct{}
	wg  *sync.WaitGroup
}

func AsyncLimiter(maxConcurrent int) *S_AsyncLimiter {
	return &S_AsyncLimiter{
		sem: make(chan struct{}, maxConcurrent),
		wg:  &sync.WaitGroup{},
	}
}

func (al *S_AsyncLimiter) Go(f func()) {
	al.wg.Add(1)
	go func() {
		defer al.wg.Done()
		al.sem <- struct{}{}
		defer func() { <-al.sem }()

		f()
	}()
}

func (al *S_AsyncLimiter) Wait() {
	al.wg.Wait()
}
