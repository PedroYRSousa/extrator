package async

func (al *S_AsyncLimiter) Go(f func()) {
	al.wg.Add(1)
	go func() {
		defer al.wg.Done()
		al.sem <- struct{}{}
		defer func() { <-al.sem }()

		f()
	}()
}
