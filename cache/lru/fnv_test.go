package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash64a(t *testing.T) {
	t.Run("fnv", func(t *testing.T) {
		v := Hash64a("this is a sample content")
		assert.NotEqual(t, v, uint64(0))
	})
}

func BenchmarkHash64a(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Hash64a("this is a sample content")
	}
}
