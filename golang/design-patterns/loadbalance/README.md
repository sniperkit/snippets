# loadbalance

This example implements the loadbalance example with multiple concurrent processes waiting for the
 first Agent resource to be acquired. When a Agent resource is acquired the other requests are canceled.

Example output of `go run main.go`:

```
-- allocating agents
 1 410ms
 2 551ms
 3 821ms
 4 51ms
 5 937ms
-- acquiring first responding agent
 3 context canceled
 5 context canceled
 1 context canceled
 2 context canceled
-- acquire finished
	took: 54.017105ms
	got agent: 4
```
