package stdns

import (
	"net"
	"fmt"
	"context"
)

type StdResolver struct {
	*net.Resolver
}

var _ Resolver = &StdResolver{}

func NewStdResolver() *StdResolver {
	return &StdResolver{
		Resolver: &net.Resolver{},
	}
}

func (sr *StdResolver) Lookup(ctx context.Context, name string) (*Entry, error) {
        txt, err := sr.LookupTXT(ctx, name)
    fmt.Println(err)
    fmt.Println(txt)
	return nil, nil
}
