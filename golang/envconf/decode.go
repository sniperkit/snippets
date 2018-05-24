package envconf

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
)

// Unmarshal key=value formatted data into v
func Unmarshal(data []byte, v interface{}) error {
	var d decodeState
	d.init(data)
	return d.unmarshal(v)
}

// Environ unmarshalls key=value formatted env items into v
func Environ(env []string, v interface{}) error {
	var d decodeState
	for _, e := range env {
		d.init([]byte(e))
		if err := d.unmarshal(v); err != nil {
			return err
		}
	}
	return nil
}

type decoderState int

const (
	decoderStateKey = iota + decoderState(1)
	decoderStateValue
)

// decodeState represents the state while decoding a env value
type decodeState struct {
	state decoderState
	data  []byte
}

type decodeKeyValue struct {
	key string
	val string
}

func (d *decodeState) init(data []byte) *decodeState {
	d.data = data
	d.state = decoderStateKey
	return d
}

func unmarshalReflect(v interface{}, dv *decodeKeyValue) {
	elem := reflect.ValueOf(v).Elem()

	for i := 0; i < elem.NumField(); i++ {
		valueField := elem.Field(i)
		if !valueField.CanSet() {
			continue
		}

		typeField := elem.Type().Field(i)
		tag := typeField.Tag.Get("envconf")

		var sep string

		switch tag {
		case "", "json":
		default:
			// TODO regex for sep fetching
			if tag == `sep=':'` {
				sep = ":"
				break
			}
			if tag == `sep=' '` {
				sep = " "
				break
			}
			continue
		}

		if !strings.EqualFold(dv.key, typeField.Name) {
			continue
		}

		if tag == "json" {
			// TODO only interface allowed for now
			var iface interface{}
			if err := json.Unmarshal([]byte(dv.val), &iface); err != nil {
				// TODO catch err
				return
			}
			if iface != nil {
				valueField.Set(reflect.ValueOf(iface))
			}
		}

		switch valueField.Interface().(type) {
		case bool:
			if b, err := strconv.ParseBool(dv.val); err == nil {
				valueField.SetBool(b)
			}
		case int:
			if i, err := strconv.ParseInt(dv.val, 10, 64); err == nil {
				valueField.SetInt(i)
			}
		case string:
			valueField.SetString(dv.val)
		case []string:
			items := strings.Split(dv.val, sep)
			valueField.Set(reflect.AppendSlice(valueField, reflect.ValueOf(interface{}(items))))
		}
	}
}

func unmarshalFinalizer(v interface{}, dv *decodeKeyValue) {
	dv.val = strings.Trim(dv.val, `"`)
	unmarshalReflect(v, dv)
}

func (d *decodeState) unmarshal(v interface{}) error {
	var val decodeKeyValue
	var buf []byte

	// collect
	for _, b := range d.data {
		switch b {
		case '_':
			if d.state == decoderStateValue {
				buf = append(buf, b)
			}
		case '=':
			val.key = string(buf)
			d.state = decoderStateValue
			buf = nil
		case '\n':
			val.val = string(buf)
			buf = nil
			unmarshalFinalizer(v, &val)
			d.state = decoderStateKey
		default:
			buf = append(buf, b)
		}
	}

	// collect last buffered item without a newline
	if len(buf) > 0 && d.state == decoderStateValue {
		val.val = string(buf)
		unmarshalFinalizer(v, &val)
	}

	return nil
}
