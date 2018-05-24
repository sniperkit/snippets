package main

import (
	"io"
	"bytes"
	"encoding/json"
	"github.com/ghodss/yaml"
	"github.com/BurntSushi/toml"
)

type Object struct {
	A int `json:"a"`
	B string `json:"b"`
}

func configureSingle(data interface{}, value interface{}) {
	var buf []byte
	switch d := data.(type) {
	case []byte:
		buf = d
	case string:
		buf = []byte(d)
	default:
		return
	}
	// try json
	if err := json.Unmarshal(buf, value); err == nil {
		return
	}
	// try yaml
	if err := yaml.Unmarshal(buf, value); err == nil {
		return
	}
	// try toml
	if _, err := toml.Decode(string(buf), value); err == nil {
		return
	}
}

func configure(value interface{}, config ...interface{}) {
	for _, cfg := range config {
		switch option := cfg.(type) {
		case io.Reader:
			var buf bytes.Buffer

			io.Copy(&buf, option)
			configureSingle(buf.Bytes(), value)
		case []byte:
			configureSingle(option, value)
		case string:
			configureSingle(option, value)
		case []string:
			for _, opt := range option {
				configureSingle(opt, value)
			}
		}
	}
}

func (o *Object) Configure(config ...interface{}) error {
	configure(o, config...)
	return nil
}

func main() {
	// empty
}
