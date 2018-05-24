package main

import (
	"fmt"
	"io/ioutil"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/dop251/goja_nodejs/console"
	"github.com/xor-gate/go-by-example/goja/sensor"
)

type requestCallback struct {
	goja.Callable
	classMethod string
	uid uint16
}

var callbacks []requestCallback

func registercallback(cb goja.Callable, classMethod string, uid uint16) {
	fmt.Println("registerCallback:", classMethod, uid)
	callbacks = append(callbacks, requestCallback{Callable: cb, uid: uid, classMethod: classMethod})
}

func Request(call goja.FunctionCall) goja.Value {
	switch len(call.Arguments) {
	case 2:
		// string, function
		cb, ok := goja.AssertFunction(call.Argument(1));
		if !ok {
			// TODO error...
			return goja.Null()
		}
		registercallback(cb, call.Argument(0).String(), 0)
	case 3:
		// string, integer, function
		cb, ok := goja.AssertFunction(call.Argument(2));
		if !ok {
			// TODO error...
			return goja.Null()
		}
		registercallback(cb, call.Argument(0).String(), uint16(call.Argument(1).ToInteger()))
	default:
		// TODO error
	}
	return goja.Null()
}


func main() {
	vm := goja.New()

	new(require.Registry).Enable(vm)
	console.Enable(vm)
	sensor.Enable(vm)

	vm.Set("Request", Request)

	b, err := ioutil.ReadFile("test.js")
	if err != nil {
		fmt.Print(err)
	}

	_, err = vm.RunString(string(b))
	if err != nil {
		fmt.Println(err)
	}

	for _, cb := range callbacks {
		this := vm.NewObject()
		this.Set("SayHello", func(call goja.FunctionCall) goja.Value {
			return vm.ToValue("World!")
		})
		cb.Callable(this, vm.ToValue(map[string]interface{}{"uid":cb.uid,"value":true}))
	}
}
