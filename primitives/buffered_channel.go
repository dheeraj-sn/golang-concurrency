package primitives

import (
	"bytes"
	"fmt"
	"os"
)

func TestBuffered() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)
	intStream := make(chan int, 10)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(os.Stdout, "Producer Done.")

		for i := 0; i < 11; i++ {
			fmt.Fprintf(os.Stdout, "Sending: %d\n", i)
			intStream <- i
		}
	}()
	for integer := range intStream {
		fmt.Fprintf(os.Stdout, "Received %v.\n", integer)
	}
}
