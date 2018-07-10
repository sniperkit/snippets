package stdns

import (
	"context"
)

type Resolver interface {
	Lookup(ctx context.Context, name string) (*Entry, error)
}

type Entry struct {
	DeviceID string
}
