package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/xor-gate/envconf"
)

func envToJson(data interface{}) string {
	var buf bytes.Buffer

	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "\t")

	if err := enc.Encode(data); err != nil {
		return ""
	}

	return buf.String()
}

func main() {
	type Env struct {
		User  string `json:"user"`
		Shell string `json:"shell"`
		Path  string `json:"path"`
		Pwd   string `json:"pwd"`
	}

	var env Env

	osEnviron := os.Environ()
	envconf.Environ(osEnviron, &env)

	fmt.Println(envToJson(env))
}
