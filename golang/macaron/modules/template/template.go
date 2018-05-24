package template

import (
	"html/template"
	"time"

	"github.com/dustin/go-humanize"
)

func NewFuncMap() []template.FuncMap {
	return []template.FuncMap{map[string]interface{}{
		"AppVer": func() string {
			return "1337"
		},
		"DateFmtShort": func(t time.Time) string {
			return t.Format("Jan 02, 2006")
		},
		"DateFmtLong": func(t time.Time) string {
			return t.Format(time.RFC1123Z)
		},
		"TimeFmtShort": func(t time.Time) string {
			return t.Format("15:04:05")
		},
		"NumCommas": humanize.Comma,
	}}
}
