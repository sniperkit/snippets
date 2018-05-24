package main

import (
	"fmt"
	"encoding/json"
)

type MyParamsBoem struct {
	Boem string
}

type MyParamsBats struct {
	Bats string
}

type MyMsg struct {
	Type string
	Params interface{}
}

func (v *MyParamsBoem) UnmarshalJSON(data []byte) error {
	type Alias MyParamsBoem
	return json.Unmarshal(data, &struct{*Alias}{Alias: (*Alias)(v),})
}

func (v *MyParamsBats) UnmarshalJSON(data []byte) error {
	type Alias MyParamsBats
	return json.Unmarshal(data, &struct{*Alias}{Alias: (*Alias)(v),})
}

func (m *MyMsg) UnmarshalJSON(data []byte) error {
	type myMsgAlias MyMsg
	msg := (*myMsgAlias)(m)
	params := &json.RawMessage{}
	msg.Params = params
	if err := json.Unmarshal(data, msg); err != nil {
		return err
	}

	isarray := false
	if []byte(*params)[0] == '[' {
		isarray = true
	}

	switch msg.Type {
	case "boem": msg.Params = &MyParamsBoem{}
	case "bats":
		if isarray {
			msg.Params = &[]MyParamsBats{}
		} else {
			msg.Params = &MyParamsBats{}
		}
	default:
		msg.Params = nil
	}

	if err := json.Unmarshal(*params, &msg.Params); err != nil {
		return err
	}

	return nil
}

func (m *MyMsg) String() string {
	switch v := m.Params.(type) {
	case *MyParamsBats:
		return fmt.Sprintf("bats: %+v", v)
	case *[]MyParamsBats:
		return fmt.Sprintf("bats[]: %+v", v)
	case *MyParamsBoem:
		return fmt.Sprintf("boem: %+v", v)
	}
	return fmt.Sprintf("Unknown <%T>", m.Params)
}

func main() {
	data := []json.RawMessage{
		json.RawMessage(`{"type":"boem","params":{"boem":"knal"}}`),
		json.RawMessage(`{"type":"bats","params":{"bats":"pow"}}`),
		json.RawMessage(`{"type":"bats","params":[{"bats":"pow"},{"bats":"tjow"}]}`),
	}

	for _, d := range data {
		msg := &MyMsg{}
		if err := json.Unmarshal(d, msg); err != nil {
			panic(err)
		}
		fmt.Printf("msg: %s\n", msg)
	}
}
