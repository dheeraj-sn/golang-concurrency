package primitives

import (
	"fmt"
	"sync"
)

func Mutexmain() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		earlierCount := count
		count++
		fmt.Printf("Earlier Count: %d,Incremented count: %d\n", earlierCount, count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		earlierCount := count
		count--
		fmt.Printf("Earlier Count: %d,Decremented count: %d\n", earlierCount, count)
	}

	var wg sync.WaitGroup
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrement()
		}()
	}
	wg.Wait()
}
