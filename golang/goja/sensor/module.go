package sensor

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	_ "github.com/dop251/goja_nodejs/util"
)

type Sensor struct {
	runtime *goja.Runtime
	Callback goja.Callable
	Info struct {
		Uid uint16
		Label string
		Type string
	}
	This map[string]interface{}
	Stats struct {
		Calls uint
	}
}

type Sensors map[uint16]*Sensor

var list Sensors

func Call(uid uint16) {
	s, ok := list[uid]
	if !ok {
		return
	}
	s.Callback(s.runtime.ToValue(&s.This))
	s.Stats.Calls++
	fmt.Printf("%+s\n", s)
}

func (s Sensor) String() string {
	return fmt.Sprintf("{\"uid\":%d,\"value\":%+v}", s.Info.Uid, s.This["value"])
}

func (c *Sensor) new(call goja.FunctionCall) goja.Value {
	s := &Sensor{}

	if cb, ok := goja.AssertFunction(call.Argument(3)); ok {
		s.Callback = cb
	} else {
		panic(c.runtime.NewTypeError("Not a function"))
		return nil
	}

	s.Info.Uid = uint16(call.Argument(0).ToInteger())
	s.Info.Label = call.Argument(1).String()
	s.Info.Type  = call.Argument(2).String()

	s.This = make(map[string]interface{})
	s.This["value"] = call.Argument(4)
	s.This["uid"]   = s.Info.Uid
	s.runtime = c.runtime

	list[s.Info.Uid] = s

	fmt.Printf("new: %s\n", s)

	return call.Argument(0)
}

func (c *Sensor) publish(call goja.FunctionCall) goja.Value {
	fmt.Println("publish")
	return nil
}

func (c *Sensor) update(call goja.FunctionCall) goja.Value {
	Call(uint16(call.Argument(0).ToInteger()))
	return nil
}

func Require(runtime *goja.Runtime, module *goja.Object) {
	c := &Sensor{
		runtime: runtime,
	}

	o := module.Get("exports").(*goja.Object)
	o.Set("new",     c.new)
	o.Set("publish", c.publish)
	o.Set("update",  c.update)
}

func Enable(runtime *goja.Runtime) {
	runtime.Set("sensor", require.Require(runtime, "sensor"))
}

func init() {
	list = Sensors{}
	require.RegisterNativeModule("sensor", Require)
}
