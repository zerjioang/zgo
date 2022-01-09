package unsafe_test

import (
	"github.com/zerjioang/zgo/assert"
	"github.com/zerjioang/zgo/unsafe"
	"testing"
)

func TestBytesToString(t *testing.T) {
	out := "hello"
	in := []byte(out)
	assert.DeepEqual(t, out, unsafe.BytesToString(in))
}

func TestStringToBytes(t *testing.T) {
	in := "hello"
	out := []byte(in)
	assert.DeepEqual(t, out, unsafe.StringToBytes(in))
}
