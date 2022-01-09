package r2_test

import (
	"github.com/zerjioang/zgo/assert"
	"github.com/zerjioang/zgo/r2"
	"reflect"
	"testing"
)

type A struct {
	B int    `json:"b"`
	C string `json:"c"`
	D *int   `json:"d"`
}

func TestKind(t *testing.T) {
	// Run these statements twice to test the type caching
	for i := 0; i < 2; i++ {
		assert.Equal(t, r2.TypeOf(true).Kind(), reflect.Bool)
		assert.Equal(t, r2.TypeOf(1).Kind(), reflect.Int)
		assert.Equal(t, r2.TypeOf("Hello").Kind(), reflect.String)
		assert.Equal(t, r2.TypeOf(struct{}{}).Kind(), reflect.Struct)
	}
}

func TestSet(t *testing.T) {
	i := 0
	typ := r2.TypeOf(i)
	assert.Equal(t, i, 0)
	typ.Set(&i, 1)
	assert.Equal(t, i, 1)
}

func TestSetField(t *testing.T) {
	a := A{}
	typ := r2.TypeOf(a).(r2.StructType)
	assert.Equal(t, a.B, 0)
	typ.SetField(&a, "B", 1)
	assert.Equal(t, a.B, 1)
}

func TestSetJSONField(t *testing.T) {
	a := A{}
	typ := r2.TypeOf(a).(r2.StructType)
	assert.Equal(t, a.B, 0)
	typ.SetFieldByJSONTag(&a, "b", 1)
	assert.Equal(t, a.B, 1)
}

func BenchmarkSetField(b *testing.B) {
	a := A{}
	typ := r2.TypeOf(a).(r2.StructType)
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		typ.SetField(&a, "B", 1)
	}
}

func BenchmarkSetFieldByJSONTag(b *testing.B) {
	a := A{}
	typ := r2.TypeOf(a).(r2.StructType)
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		typ.SetFieldByJSONTag(&a, "b", 1)
	}
}
