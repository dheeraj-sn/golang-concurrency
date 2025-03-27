package generators

import (
	"fmt"
	"time"
)

func RepeatGenerator(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			for _, value := range values {
				select {
				case <-done:
					return
				case valueStream <- value:
				}
			}
		}
	}()
	return valueStream
}

func CheckRepeatGenerator() {
	done := make(chan interface{})
	go func() {
		time.Sleep(1 * time.Second)
		close(done)
	}()

	repeatStream := RepeatGenerator(done, "Hello", "World")
	for value := range repeatStream {
		fmt.Println(value)
	}
}
