package main

type Environment struct {
	name string
}

func NewEnvironment(name string) *Environment {
	e := &Environment{name: name}
	return e
}
