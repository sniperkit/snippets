package main

import (
	"bytes"
	"encoding/json"
)

type JSONMapString map[string]string

func (m JSONMapString) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("{")
	i := 0
	for k, v := range m {
		if i > 0 {
			buf.WriteString(",")
		}
		i++
		kjs, err := json.Marshal(k)
		if err != nil {
			return nil,err
		}
		buf.Write(kjs)
		buf.WriteString(":")
		buf.WriteString(v)

	}
	buf.WriteString("}")
	return buf.Bytes(),nil
}


func main() {
}
