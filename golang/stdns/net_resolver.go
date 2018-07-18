package stdns

import (
	"net"
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

func (sr *StdResolver) Lookup(ctx context.Context, uri string) (Entries, error) {
	u, err := ParseURI(uri)
	if err != nil {
		return nil, err
	}

        txt, err := sr.LookupTXT(ctx, u.Hostname())
	if err != nil {
		return nil, err
	}

	return DecodeTXTRecords(u, txt)
}
