package main

import (
	"fmt"
	"sync"
	"time"
	"context"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Spawn a handfull of goroutines
	for i := 1; i <= 5; i++ {
		// Add the goroutine to the WaitGroup before starting
		wg.Add(1)
		go func(wid int) {
			// When the goroutine is finished remove it from the WaitGroup
			defer func() {
				wg.Done()
				fmt.Printf("worker %d done\n", wid)
			}()

			// Infinite loop with context Done channel
			for {
				select {
				case <-ctx.Done():
					// Somebody canceled us
					return
				default:
					// Dummy work which takes one second
					fmt.Printf("worker %d\n", wid)
					time.Sleep(time.Second)
				}
			}
		}(i)
	}

	// Reap all the running goroutines after 3 seconds with the cancel() method
	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 3)
		cancel()
		wg.Done()
	}()

	// Wait until all goroutines are finished including the reaper
	wg.Wait()
}
