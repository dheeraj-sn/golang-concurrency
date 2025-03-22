package primitives

import (
	"fmt"
	"sync"
)

type Button struct {
	Clicked *sync.Cond
}

func CondBroadcast() {
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func(*sync.WaitGroup), extWg *sync.WaitGroup) {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			wg.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn(extWg)
		}()
		wg.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)

	maximizeWindows := func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("Maximizing Windows")
	}

	displayDialogBox := func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("Displaying dialog box")
	}

	mouseClicked := func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("Mouse clicked")
	}

	subscribe(button.Clicked, maximizeWindows, &clickRegistered)
	subscribe(button.Clicked, displayDialogBox, &clickRegistered)
	subscribe(button.Clicked, mouseClicked, &clickRegistered)

	button.Clicked.Broadcast()
	clickRegistered.Wait()
}
