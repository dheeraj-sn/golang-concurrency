package patterns

import (
	"fmt"
	"math/rand"
	"time"
)

// protect read
func DoneChannelForRead() {

	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		workTerminated := make(chan interface{})
		defer fmt.Println("doWork: hello")
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(workTerminated)
			for {
				select {
				case <-done:
					return
				case s := <-strings:
					fmt.Printf("work %v\n", s)
				}
			}
		}()
		return workTerminated
	}

	done := make(chan interface{})
	workTerminated := doWork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("close done, exit do work now")
		close(done)
	}()
	<-workTerminated

}

// protect write
func DoneChannelForWrite() {
	randomStream := func(done <-chan interface{}) <-chan int {
		rstream := make(chan int)
		go func() {
			defer fmt.Println("randomStream: stops sending")
			for {
				select {
				case <-done:
					fmt.Printf("randomStream: done\n")
					return
				default:
					rstream <- rand.Int()
				}
			}
		}()
		return rstream
	}

	done := make(chan interface{})
	rstream := randomStream(done)
	for i := 0; i < 3; i++ {
		fmt.Printf("randomStream: %d\n", <-rstream)
	}
	close(done)
	time.Sleep(2 * time.Second)
}
