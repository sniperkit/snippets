package main

import (
	"fmt"
	"time"
	"context"
)

func main() {
	d := time.Now().Add(100 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	go func() {
		time.Sleep(time.Second)
		fmt.Println("did some stuff")
		cancel()
	}()

	<-ctx.Done()
	if ctx.Err() == context.Canceled {
		fmt.Println("it is done yuri")
	} else {
		fmt.Println("timeout ...")
	}
}
