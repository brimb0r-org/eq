package worker_pool

import "sync"

var W int = 5

var wg sync.WaitGroup

func Worker(task func(interface{})) chan interface{} {
	input := make(chan interface{}, W)
	for i := 0; i < W; i++ {
		go func() {
			for {
				v, ok := <-input
				if ok {
					task(v)
				} else {
					return
				}
			}
		}()
	}
	return input
}
