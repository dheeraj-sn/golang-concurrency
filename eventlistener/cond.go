package eventlistener

import (
	"fmt"
	"sync"
	"time"
)

// Shared resource
var queue []int
var cond *sync.Cond

func producer() {
	for i := 0; i < 5; i++ {
		cond.L.Lock()
		queue = append(queue, i)
		fmt.Printf("Produced: %d\n", i)
		cond.Signal() // Wake up one consumer
		cond.L.Unlock()
		time.Sleep(time.Second)
	}
}

func consumer() {
	for {
		cond.L.Lock()
		// Wait while queue is empty
		for len(queue) == 0 {
			fmt.Println("Consumer waiting...")
			cond.Wait() // Key point of explanation
		}

		// Process item
		item := queue[0]
		queue = queue[1:]
		fmt.Printf("Consumed: %d\n", item)
		cond.L.Unlock()
	}
}

func Xmain() {
	cond = sync.NewCond(&sync.Mutex{})
	queue = make([]int, 0)

	go producer()
	go consumer()

	// Let the program run for a while
	time.Sleep(6 * time.Second)
}
