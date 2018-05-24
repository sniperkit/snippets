package main

import (
	"os"
	"fmt"
	"time"
	"errors"
	"context"
//	"runtime"
	"runtime/pprof"
)

const resourceCount = 1000000

var ErrResourceBusy = errors.New("resource busy")
var ErrResourceIdle = errors.New("resource idle")

// Resource which can be acquired and released
type Resource struct {
	id int
	c chan bool
}

// NewResource creates a new "idle" resource
func NewResource(id int) *Resource {
	return &Resource{
		id: id,
		c: make(chan bool, 1),
	}
}

// Acquire the resource. When already acquired it blocks forever, or when the context is finished
func (r *Resource) Acquire(ctx context.Context, nonBlock bool) error {
	if nonBlock {
		select {
		case r.c <-true:
			break
		default:
			return ErrResourceBusy
		}
	} else {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case r.c <-true:
			break
		}
	}
	return nil
}

// Release the resource. When it is idle it returns immediate with ErrResourceIdle
func (r *Resource) Release(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-r.c:
		break
	}
	return nil
}

// Resources is holds a group of resource items
type Resources struct {
	l []*Resource
}

func NewResources() *Resources {
	return &Resources{}
}

func (r *Resources) Add(res *Resource) {
	r.l = append(r.l, res)
}

// Acquire gets a "idle" resource from the resources group
func (r *Resources) Acquire(ctx context.Context) (*Resource, error) {
	for _, res := range r.l {
		if err := res.Acquire(ctx, true); err == nil {
			return res, nil
		}
	}
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return nil, ErrResourceBusy
	}
}

func main() {
	f, _ := os.Create("profile.pb.gz")
	defer f.Close()


	// Allocate new resources
	res := NewResources()
	for i := 0; i < resourceCount; i++ {
		r := NewResource(i + 1)
		// TODO: all resources are busy now
		// Resource 1 and 4 are faked to be busy...
		//if r.id == 1 || r.id == 4 {
		if i != (resourceCount - 1) {
			r.c<-true
		}
		//}
		res.Add(r)
	}

	fmt.Println(len(res.l), "resources allocated")
//	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	ctx := context.Background()
	pprof.StartCPUProfile(f)
	now := time.Now()
	r, err := res.Acquire(ctx)
	fmt.Println("took:", time.Since(now))
	pprof.StopCPUProfile()
	if err == nil {
		fmt.Println("got:",r.id)
	} else {
		fmt.Println("err:",err)
	}

	//runtime.GC()
	//pprof.WriteHeapProfile(f);
	//pprof.StopCPUProfile()
}
