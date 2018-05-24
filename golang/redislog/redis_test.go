package redislog

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/stretchr/testify/require"
)

var uri string
var mr *miniredis.Miniredis

func init() {
	uri = os.Getenv("CICD_TEST_REDIS_URI")
	if uri == "" {
		s, err := miniredis.Run()
		if err != nil {
			panic(err)
		}
		mr = s
		uri = s.Addr()
	}
}

func TestNew(t *testing.T) {
	rl, err := New(uri, t.Name())
	require.Nil(t, err)
	rl.Info("Hello from ", t.Name(), time.Now())
}

func TestSubscribe(t *testing.T) {
	rl, err := New(uri, t.Name())
	require.Nil(t, err)

	go func() {
		for {
			rl.Info(time.Now())
			time.Sleep(time.Second)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rl.Subscribe(ctx, os.Stdout)
	<-ctx.Done()
}
