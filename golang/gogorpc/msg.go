package gogorpc

import (
	"time"
	"strings"
	"unicode"
	"encoding/json"
)

const DinetRPCVersion1 = 1
var ClassMethod_JSON_name map[int32]string
var ClassMethod_JSON_value map[string]int32

var Errno_descr = map[int32]string{
	int32(Errno_Ok): "Ok",
	int32(Errno_OpNotSupp): "Operation not supported",
}

func init() {
	ClassMethod_JSON_name = make(map[int32]string)
	for idx, name := range ClassMethod_name {
		ClassMethod_JSON_name[idx] = classMethodJSONString(name)
	}

	ClassMethod_JSON_value = make(map[string]int32)
	for idx, name := range ClassMethod_name {
		ClassMethod_JSON_value[classMethodJSONString(name)] = idx
	}
}

// Convert camelcase protobuf ClassMethod enum name into into DI-Net RPC semicolon separated
func classMethodJSONString(name string) string {
        var words []string
        l := 0
        for s := name; s != ""; s = s[l:] {
                l = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
                if l <= 0 {
                        l = len(s)
                }
                words = append(words, strings.ToLower(s[:l]))
        }
	return strings.Join(words, ":")
}

func (cm ClassMethod) MarshalJSON() ([]byte, error) {
	name, ok := ClassMethod_JSON_name[int32(cm)]
	if !ok {
		name = ClassMethod_JSON_name[int32(ClassMethod_UnknownUnknown)]
	}
	return json.Marshal(name)
}

func (cm *ClassMethod) UnmarshalJSON(b []byte) error {
	var name string

	if err := json.Unmarshal(b, &name); err != nil {
		return err
	}

	enum, ok := ClassMethod_JSON_value[name]
	if ok {
		*cm = ClassMethod(enum)
	} else {
		*cm = ClassMethod_UnknownUnknown
	}

	return nil
}

func ErrorFromErrno(errno Errno) *Error {
	descr, ok := Errno_descr[int32(errno)]
	if !ok {
		return nil
	}
	return &Error{Code: errno, Descr: descr}
}

func (p *Value) MarshalJSON() ([]byte, error) {
	var value interface{}

	if x, ok := p.GetValue().(*Value_Bool); ok {
		value = x.Bool
	} else if x, ok := p.GetValue().(*Value_Number); ok {
		value = x.Number
	} else if x, ok := p.GetValue().(*Value_Str); ok {
		value = x.Str
	}

	return json.Marshal(value)
}

type Header struct {
	*HeaderMsg
}

func (r *Header) Validate() bool {
	hdr := r.HeaderMsg
	if hdr == nil {
		return false
	}
	if hdr.DinetRPC != DinetRPCVersion1 {
		return false
	}
	if hdr.Time == 0 {
		return false
	}
	deviceUIDLen := len(hdr.DeviceUID)
	if deviceUIDLen != 0 && deviceUIDLen != 32 {
		return false
	}
	return true
}

type Request struct {
	*RequestMsg
}

func (r *Request) String() string {
	j, _ := json.Marshal(r)
	return string(j)
}

func (r *Request) Validate() bool {
	if r.RequestMsg == nil {
		return false
	}
	hdr := &Header{&r.RequestMsg.HeaderMsg}
	if !hdr.Validate() {
		return false
	}
	if r.ClassMethod == ClassMethod_UnknownUnknown {
		return false
	}
	return true
}

func (r *Request) Param(key string, value interface{}) {
	var pv Value

	switch v := value.(type) {
	case string:
		pv.Value = isValue_Value(&Value_Str{v})
	case int:
		pv.Value = isValue_Value(&Value_Number{float64(v)})
	default:
		return
	}

	if r.Params == nil {
		r.Params = make(map[string]*Value)
	}

	r.Params[key] = &pv
}

func (r *Request) Reply() *Reply {
	rep := &Reply{&ReplyMsg{}}
	rep.ID = r.ID
	rep.DeviceUID = r.DeviceUID
	rep.ClassMethod = r.ClassMethod
	return rep
}

type Reply struct {
	*ReplyMsg
}

func (r *Reply) String() string {
	j, _ := json.Marshal(r)
	return string(j)
}

func (r *Reply) Validate() bool {
	if r.ReplyMsg == nil {
		return false
	}
	hdr := &Header{&r.ReplyMsg.HeaderMsg}
	if !hdr.Validate() {
		return false
	}
	if r.ClassMethod == ClassMethod_UnknownUnknown {
		return false
	}
	return true
}

type Publish struct {
	*PublishMsg
}

func (p *Publish) String() string {
	j, _ := json.Marshal(p)
	return string(j)
}

func TimeNow() uint64 {
	return uint64(time.Now().UnixNano() / int64(time.Millisecond))
}

func (p *PublishMsg) MarshalJSON() ([]byte, error) {
	var result interface{}
	type Alias PublishMsg

	switch r := p.Result.Result.(type) {
	case *Result_ResultValueItems:
		result = r.ResultValueItems.GetItems()
	}

	return json.Marshal(&struct {
		Result interface{} `json:"result,omitempty"`
	        *Alias
	    }{
		Result: result,
	        Alias:    (*Alias)(p),
	    })
}
