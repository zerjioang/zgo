package strconv_test

import (
	"github.com/zerjioang/zgo/assert"
	"github.com/zerjioang/zgo/strconv"
	strconvstd "strconv"
	"testing"
)

func TestDecToInt(t *testing.T) {
	assert.Equal(t, int64(0), strconv.DecToInt([]byte("")))
	assert.Equal(t, int64(0), strconv.DecToInt([]byte("0")))
	assert.Equal(t, int64(1), strconv.DecToInt([]byte("1")))
	assert.Equal(t, int64(10), strconv.DecToInt([]byte("10")))
	assert.Equal(t, int64(100), strconv.DecToInt([]byte("100")))
	assert.Equal(t, int64(123456789), strconv.DecToInt([]byte("123456789")))
	assert.Equal(t, int64(0), strconv.DecToInt([]byte("ZZZ")))
}

func TestHexToInt(t *testing.T) {
	assert.Equal(t, int64(0x0), strconv.HexToInt([]byte("")))
	assert.Equal(t, int64(0x0), strconv.HexToInt([]byte("0")))
	assert.Equal(t, int64(0x1), strconv.HexToInt([]byte("1")))
	assert.Equal(t, int64(0xA), strconv.HexToInt([]byte("A")))
	assert.Equal(t, int64(0x10), strconv.HexToInt([]byte("10")))
	assert.Equal(t, int64(0xAFFE), strconv.HexToInt([]byte("AFFE")))
	assert.Equal(t, int64(0xAFFE), strconv.HexToInt([]byte("Affe")))
	assert.Equal(t, int64(0xCAFE), strconv.HexToInt([]byte("CAFE")))
	assert.Equal(t, int64(0xCAFE), strconv.HexToInt([]byte("Cafe")))
	assert.Equal(t, int64(0xBADFACE), strconv.HexToInt([]byte("BADFACE")))
	assert.Equal(t, int64(0xBADFACE), strconv.HexToInt([]byte("BadFace")))
	assert.Equal(t, int64(0), strconv.HexToInt([]byte("ZZZ")))
}

func BenchmarkDecToInt(b *testing.B) {
	example := []byte("123456789")
	b.ResetTimer()
	b.ReportAllocs()
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		strconv.DecToInt(example)
	}
}

func BenchmarkStrconvDecToInt(b *testing.B) {
	example := "123456789"
	b.ResetTimer()
	b.ReportAllocs()
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		_, _ = strconvstd.ParseInt(example, 10, 64)
	}
}

func BenchmarkHexToInt(b *testing.B) {
	example := []byte("CAFE")
	b.ResetTimer()
	b.ReportAllocs()
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		strconv.HexToInt(example)
	}
}

func BenchmarkStrconvHexToInt(b *testing.B) {
	example := "CAFE"
	b.ResetTimer()
	b.ReportAllocs()
	b.SetBytes(1)
	for i := 0; i < b.N; i++ {
		_, _ = strconvstd.ParseInt(example, 16, 64)
	}
}
