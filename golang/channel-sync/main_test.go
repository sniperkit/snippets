package main

import (
	"time"
	"context"
	"testing"
)

const testResourcesCount = 1000

func newResourcesAllBusy(count int) *Resources {
	res := NewResources()
	for i := 1; i < count; i++ {
		r := NewResource(i + 1)
		r.c<-true
		res.Add(r)
	}
	return res
}

func newResourcesAllIdle(count int) *Resources {
	res := NewResources()
	for i := 1; i < count; i++ {
		r := NewResource(i + 1)
		res.Add(r)
	}
	return res
}

// Create new resources list where the last item is idle
func newResourcesLastIdle(count int) *Resources {
	res := NewResources()
	for i := 0; i < count; i++ {
		r := NewResource(i + 1)
		r.c<-true
		res.Add(r)
	}
	// last resource make it idle again
	<-res.l[count-1].c
	return res
}

func TestAcquireAllBusy(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second * 10))
	defer cancel()
	res := newResourcesAllBusy(testResourcesCount)
	_, err := res.Acquire(ctx)
	if err != ErrResourceBusy {
		t.Error("unexpected err", err)
	}
}

func TestAcquireAllIdle(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second * 10))
	defer cancel()
	resources := newResourcesAllIdle(testResourcesCount)

	// Acquire a single resource from the resource group
	res, err := resources.Acquire(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	if res == nil {
		t.Error("res should not be nil")
		return
	}
	res.Release(context.TODO())

	// Check if all resources idle by Acquiring them all
	for _, r := range resources.l {
		err := r.Acquire(context.TODO(), false)
		if err != nil {
			t.Error("unable to acquire resource:", err)
			return
		}
		err = r.Release(context.TODO())
		if err != nil {
			t.Error("unable to release resource:", err)
			return
		}
	}
}

func TestAcquireLastIdle(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second * 10))
	defer cancel()
	resources := newResourcesLastIdle(testResourcesCount)
	res, err := resources.Acquire(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	if res == nil {
		t.Error("res should not be nil")
		return
	}
	if res.id != testResourcesCount {
		t.Error("unexpected res.id", res.id)
	}
}
