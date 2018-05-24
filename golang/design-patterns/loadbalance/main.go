package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const nrOfAgents = 5
const acquireDeadline = time.Second

// Agent resource
type Agent struct {
	id           int
	acquire      chan bool
	acquireDelay time.Duration
}

func (a *Agent) Acquire(ctx context.Context) error {
	// Fake the agent busy loop
	timer := time.NewTimer(a.acquireDelay).C
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer:
		break
	}

	// Trying to acquire the agent using the acquire token channel
	select {
	case <-ctx.Done():
		return ctx.Err()
	case a.acquire <- true:
		fmt.Printf("%2d acquired\n", a.id)
		break
	}
	return nil
}

func (a *Agent) Release() {
	<-a.acquire
}

// Agents groups multiple Agent resources
type Agents struct {
	l []*Agent
}

// NewAgents creates an empty agent resource group
func NewAgents() *Agents {
	return &Agents{}
}

// Add creates a new Agent and adds it to the group
func (a *Agents) Add(id int, acquireDelay time.Duration) {
	agent := &Agent{
		id:           id,
		acquire: make(chan bool, 1),
		acquireDelay: acquireDelay,
	}
	a.l = append(a.l, agent)
}

// Acquire gets the first responding agent from the group
func (a *Agents) Acquire(ctx context.Context) (*Agent, error) {
	var wg sync.WaitGroup
	var acquiredAgent *Agent

	// Concurrent try to Acquire all agents in the group
	acquireCtx, acquireCancel := context.WithCancel(ctx)

	for _, agent := range a.l {
		// Track the goroutine execution inside the WaitGroup
		wg.Add(1)
		go func(ctx context.Context, a *Agent) {
			defer wg.Done()
			if err := a.Acquire(ctx); err != nil {
				fmt.Printf("%2d %v\n", a.id, err)
				return
			}
			// The first-responder cancels the other blocking Acquire calls
			acquireCancel()
			acquiredAgent = a
		}(acquireCtx, agent)
	}

	// Wait until all goroutines are finished
	wg.Wait()

	// Errors other than canceled should move up the call chain
	if err := acquireCtx.Err(); err != nil && err != context.Canceled {
		return nil, err
	}
	return acquiredAgent, nil
}

func main() {
	// Create a group of agents with random acquire delay
	agents := NewAgents()

	fmt.Println("-- allocating agents")
	for id := 1; id <= nrOfAgents; id++ {
		// XXX: change to time.Second to see the deadline reached
		acquireDelay := time.Millisecond * time.Duration((rand.Int() % 1000))
		agents.Add(id, acquireDelay)
		fmt.Printf("%2d %v\n", id, acquireDelay)
	}

	// Create application context with a deadline
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(acquireDeadline))
	defer cancel()

	// Try to Acquire the first responding agent
	fmt.Println("-- acquiring first responding agent")
	now := time.Now()
	agent, err := agents.Acquire(ctx)
	fmt.Println("-- acquire finished")
	fmt.Printf("\ttook: %v\n", time.Since(now))
	if err != nil {
		fmt.Println("\tfailed:", err)
		return
	}
	fmt.Println("\tgot agent:", agent.id)
	agent.Release()
}
