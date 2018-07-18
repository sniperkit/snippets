package stdns

import (
	"fmt"
	"net/url"
	"strings"
	"strconv"
)

func DecodeTXT(uri *url.URL, txtRecord string) (*Entry, error) {
	fields := strings.Fields(txtRecord)

	if len(fields) < 3 {
		return nil, fmt.Errorf("invalid")
	}

	if fields[0] != "st" {
		return nil, fmt.Errorf("invalid")
	}

	entryType, err := strconv.ParseUint(fields[1], 10, 16)
	if err != nil {
		return nil, fmt.Errorf("unexpected entry type %d", entryType)
	}

	e := &Entry{}

	switch EntryType(entryType) {
	case EntryDeviceID:
		if len(fields) != 3 {
			return nil, fmt.Errorf("unexpected field length")
		}
		e.DeviceID = fields[2]
	default:
		return nil, fmt.Errorf("unexpected entry type %d", entryType)
	}

	e.URL = uri
	return e, nil
}

func DecodeTXTRecords(uri *url.URL, txtRecords []string) (Entries, error) {
	var entries Entries

	for _, txt := range txtRecords {
		e, err := DecodeTXT(uri, txt)
		if err != nil {
			fmt.Println(err)
			continue
		}
		entries = append(entries, e)
	}
	return entries, nil
}
