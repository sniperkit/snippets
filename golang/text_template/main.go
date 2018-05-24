package main

import (
	"os"
	"text/template"
)

func main() {
	const tmpl =`Hello {{.User}}
Your homedir is {{.Home}}`

	env := &struct {
		User string
		Home string
	}{
		User: os.Getenv("USER"),
		Home: os.Getenv("HOME"),
	}

	t := template.Must(template.New("hello").Parse(tmpl))
	t.Execute(os.Stdout, env)
}
