package primitives

import (
	"fmt"
	"sync"
)

func Wgmain() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Println("Hello Goroutine with id:%d", id)
	}

	const numberOfGoroutines = 15
	var wg sync.WaitGroup
	wg.Add(numberOfGoroutines)
	for i := 0; i < numberOfGoroutines; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}
