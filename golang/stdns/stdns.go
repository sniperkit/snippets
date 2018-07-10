package stdns

import (
	"net/url"
	"context"
)

type Resolver interface {
	Lookup(ctx context.Context, url string) (Entries, error)
}

type EntryType uint16

const (
	EntryDeviceID EntryType = 0
)

type Entry struct {
	DeviceID string
	URL *url.URL
}

type Entries []*Entry

func (e *Entry) Username() string {
	if e.URL == nil {
		return ""
	}
	return e.URL.User.Username()
}

func (e *Entry) Hostname() string {
	if e.URL == nil {
		return ""
	}
	return e.URL.Hostname()
}
