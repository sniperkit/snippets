// Periodic tasks using time.Ticker and goroutine with a channel
package main

import (
	"fmt"
	"time"
)

type Periodic struct {
	ticker *time.Ticker
	quit chan bool
}

func New(d time.Duration) *Periodic {
	p := Periodic{}
	p.ticker = time.NewTicker(d)
	p.quit = make(chan bool)

	go func() {
		for {
			select {
			case <- p.ticker.C:
					fmt.Println("Tick!")
			case <- p.quit:
					p.ticker.Stop()
					return
			}
		}
	}()
	
	return &p
}

func (p *Periodic) Stop() {
	p.quit <- true
}

func main() {
	p := New(time.Millisecond * 250)
	time.Sleep(time.Second * 1)
	p.Stop()
}
