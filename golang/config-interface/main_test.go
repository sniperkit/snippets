package main

import (
	"testing"
	"strings"
	"github.com/stretchr/testify/assert"
)

func TestJSONReader(t *testing.T) {
	obj := &Object{}
	obj.Configure(strings.NewReader(`{"a":666,"b":"is the number of the beast"}`))
	assert.Equal(t, 666, obj.A)
	assert.Equal(t, "is the number of the beast", obj.B)
}

func TestYAMLReader(t *testing.T) {
	obj := &Object{}
	obj.Configure(strings.NewReader("a: 666\nb: is the number of the beast"))
	assert.Equal(t, 666, obj.A)
	assert.Equal(t, "is the number of the beast", obj.B)
}

func TestTOMLReader(t *testing.T) {
	obj := &Object{}
	obj.Configure(strings.NewReader("a = 666\nb = \"is the number of the beast\""))
	assert.Equal(t, 666, obj.A)
	assert.Equal(t, "is the number of the beast", obj.B)
}
