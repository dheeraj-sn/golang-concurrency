package generators

import "fmt"

func Take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func CheckTake() {
	done := make(chan interface{})
	defer close(done)
	for value := range Take(done, RepeatGenerator(done, 1, 2, 3), 10) {
		fmt.Println(value)
	}
}
