package async

func (al *S_AsyncLimiter) Wait() {
	al.wg.Wait()
}
