package primitives

import (
	"fmt"
	"sync"
)

func PoolTest() {
	var numCalcCreated int

	createMemory := func() interface{} {
		numCalcCreated += 1
		mem := make([]byte, 1024)
		return &mem
	}

	calcPool := &sync.Pool{New: createMemory}

	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	work := func() {
		defer wg.Done()
		mem := calcPool.Get().(*[]byte)
		defer calcPool.Put(mem)
	}
	for i := 0; i < numWorkers; i++ {
		go work()
	}
	wg.Wait()
	fmt.Printf("%d calculator were created.", numCalcCreated)

}
