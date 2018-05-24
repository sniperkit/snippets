package main

import (
	"fmt"
	"encoding/json"
)

func ExampleJSONMapString_MarshalJSON() {
	// Values are raw JSON strings
	myvals := map[string]string {
		"val1" : `{"boem":"knal"}`,
		// We just show one key, value because go will have a random order in the print output
		// so we can not check the output because it changes between test runs
	}
	
	js, err := json.Marshal(JSONMapString(myvals))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(js))

	// Output: {"val1":{"boem":"knal"}}
}
