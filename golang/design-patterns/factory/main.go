package main

import (
	"fmt"
	"reflect"
)

type Object interface {
	String() string
}

type ObjectOne struct{}

func (o *ObjectOne) String() string {
	return "ObjectOne"
}

type ObjectTwo struct{}

func (o *ObjectTwo) String() string {
	return "ObjectTwo"
}

var registry []Object = []Object{&ObjectOne{}, &ObjectTwo{}}

func New(name string) (Object, error) {
	var object interface{}

	for _, obj := range registry {
		rname := reflect.ValueOf(obj).Type().Elem().Name()
		if rname == name {
			object = obj
			break
		}
	}

	if object == nil {
		return nil, fmt.Errorf("not found")
	}

	return reflect.New(reflect.ValueOf(object).Type()).Elem(), nil
}

func main() {
	for _, name := range []string{"ObjectOne", "ObjectTwo", "ObjectUnknown"} {
		obj, err := New(name)
		fmt.Println(obj, err)
	}
}
