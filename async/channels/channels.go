package channels

import "sync"

func async(fn func() error) chan error {
	out := make(chan error)
	go func() {
		defer close(out)
		out <- fn()
	}()
	return out
}

func merge(chs ...chan error) chan error {
	var (
		wg  = sync.WaitGroup{}
		out = make(chan error, len(chs))
	)
	wg.Add(len(chs))
	for _, c := range chs {
		go func(ch chan error) {
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
