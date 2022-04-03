package channels

import "sync"

func async[T any](fn func() T) chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		out <- fn()
	}()
	return out
}

func merge[T any](chs ...chan T) chan T {
	var (
		wg  = sync.WaitGroup{}
		out = make(chan T, len(chs))
	)
	wg.Add(len(chs))
	for _, c := range chs {
		go func(ch chan T) {
			defer wg.Done()
			val := <-ch
			out <- val
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
