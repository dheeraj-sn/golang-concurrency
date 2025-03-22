package primitives

import (
	"fmt"
	"sync"
	"time"
)

func CMain() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]int, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		item := queue[0]
		queue = queue[1:]
		fmt.Println("Removed from queue:%d", item)
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Adding to queue:%d", i)
		queue = append(queue, i)
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}

}
