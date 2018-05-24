package gogorpc

import (
	"fmt"
	"testing"
	"encoding/json"
	"github.com/gogo/protobuf/proto"
)

func TestEnumMapper(t *testing.T) {
	fmt.Println(classMethodJSONString("DevicePing"))
	fmt.Println(proto.EnumName(ClassMethod_JSON_name, int32(ClassMethod_DevicePing)))
}

func TestRequestReplyMsg(t *testing.T) {
	r := &Request{RequestMsg: &RequestMsg{}}

	r.ClassMethod = ClassMethod_DevicePing
	r.ID = 1234
	r.Time = TimeNow()
	r.DeviceUID = "deadbeefdeadbeefdeadbeefdeadbeef"
	r.Param("uid", 1)
	r.Param("value", 1337)

	fmt.Println("r.Validate:", r.Validate())
	fmt.Println(r)

	rep := r.Reply()
	rep.Time = TimeNow()
	rep.Error = ErrorFromErrno(Errno_OpNotSupp)

	fmt.Println("rep.Validate:", rep.Validate())
	fmt.Println(rep)
}

func TestJSONUnmarshal(t *testing.T) {
	req := &Request{}
	const reqString = `{"dinetrpc":0,"time":1525901757421,"device:uid":"deadbeefdeadbeefdeadbeefdeadbeef","id":1234,"req":"device:ping","params":{"uid":1,"value":1337}}`

	json.Unmarshal([]byte(reqString), req)
	fmt.Println(req)
}

func TestPublishMsg(t *testing.T) {
	p := &Publish{PublishMsg: &PublishMsg{}}
	p.DinetRPC = DinetRPCVersion1
	p.DeviceUID = "deadbeefdeadbeefdeadbeefdeadbeef"
	p.ClassMethod = ClassMethod_DeviceUserData
	p.Time = TimeNow()

	var rvi []*ResultValueItem

	rvi = append(rvi, &ResultValueItem{Uid: 1, Time: TimeNow(), Value: &Value{Value: isValue_Value(&Value_Bool{Bool: true})}})
	rvi = append(rvi, &ResultValueItem{Uid: 2, Time: TimeNow(), Value: &Value{Value: isValue_Value(&Value_Number{Number: 1234})}})

	p.Result = &Result{Result: &Result_ResultValueItems{ResultValueItems: &ResultValueItems{Items: rvi}}}

	fmt.Println(p)
}
