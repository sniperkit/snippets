package stdns

import (
	"net/url"
	"strings"
)

func ParseURI(uri string) (*url.URL, error) {
	if !strings.HasPrefix(uri, URLScheme) {
		uri = URLScheme + uri
	}
	return url.Parse(uri)
}
