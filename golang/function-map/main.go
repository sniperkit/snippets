// Map of strings to functions
// http://stackoverflow.com/questions/6769020/go-map-of-functions
// http://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
package main

import "log"

type fn func (string)
var m map[string] fn

func foo(msg string) {
	log.Printf("foo! Message is %s", msg)
}

func bar(msg string) {
	log.Printf("bar! Message is %s", msg)
}

func execute(function, msg string) {
	if _func, ok := m[function]; ok {
		_func(msg)
		return
	}
	log.Println("unknown function:", function)
}

func init() {
	m = map[string] fn {
		"foo": foo,
		"bar": bar,
	}
}

func main() {
	log.Printf("map is %v", m)

	execute("foo", "hello")
	execute("bar", "world!")
	execute("bla", "bla")
}
