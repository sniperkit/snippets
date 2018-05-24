// http://stackoverflow.com/questions/16466320/is-there-a-way-to-do-repetitive-tasks-at-intervals-in-golang
package main

import (
    "fmt"
    "time"
)

type Task struct {
	name string
	running chan bool
	action func()
	ch chan bool
	interval time.Duration
}

func NewTask(name string, action func(), interval time.Duration) *Task {
	var t Task
	t.name = name
	t.running = make(chan bool)
	t.ch      = make(chan bool)
	t.action = action
	t.interval = interval
	return &t
}

func (t *Task) Start() {
	go func() {
	for {
		t.action()
		select {
		case <-time.After(t.interval):
		case <-t.ch:
			fmt.Println("Stopped", t.name)
			t.running <- false
			return
		}
	}
	}()
}

func (t *Task) Stop() {
	fmt.Println("Stop signal");
	t.ch <- true

	<-t.running

	close(t.ch)
	close(t.running)
}

func main() {
	t1 := NewTask("t1", func() { fmt.Println("(t1) #") }, 250 * time.Millisecond)
	defer t1.Stop()
	t1.Start()

	t2 := NewTask("t2", func() { fmt.Println("(t2) #") }, time.Second)
	defer t2.Stop()
	t2.Start()

	t3 := NewTask("t3", func() { fmt.Println("(t3) #") }, time.Second)
	defer t3.Stop()
	t3.Start()

	t4 := NewTask("t4", func() { fmt.Println("(t4) #") }, 100 * time.Millisecond)
	defer t4.Stop()
	t4.Start()

	time.Sleep(3 * time.Second)
}
