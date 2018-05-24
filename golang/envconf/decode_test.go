package envconf

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeString(t *testing.T) {
	type Env struct {
		User string
	}
	const tvExpString = "_xor-gate_"
	tvs := []struct {
		Name   string
		Format string
	}{
		{"exact", "User=%s"},
		{"case", "USER=%s"},
		{"keyUnderscores", "U_SE_R=%s"},
		{"unquoteValue", `USER="%s"`},
	}

	for _, tv := range tvs {
		t.Run(tv.Name, func(t *testing.T) {
			var env Env
			require.Nil(t, Unmarshal([]byte(fmt.Sprintf(tv.Format, tvExpString)), &env))
			assert.Equal(t, tvExpString, env.User)
		})
	}
}

func ExampleEnviron() {
	type Env struct {
		User string
	}

	var env Env

	osEnviron := os.Environ()
	Environ(osEnviron, &env)

	fmt.Println(env)
}

func ExampleUnmarshal() {
	type Env struct {
		User          string
		Path          []string    `envconf:"sep=':'"`
		SSHConnection []string    `envconf:"sep=' '"`
		MyJSONStruct  interface{} `envconf:"json"`
	}

	const data = `USER=xor-gate
PATH=/bin:/usr/bin:/usr/local/bin
SSH_CONNECTION=192.168.1.205 63983 192.168.1.201 22
MY_JSON_STRUCT={"hello":"world"}
`

	var env Env

	err := Unmarshal([]byte(data), &env)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(env)

	// Output: {xor-gate [/bin /usr/bin /usr/local/bin] [192.168.1.205 63983 192.168.1.201 22] map[hello:world]}
}
