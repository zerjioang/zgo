//
// Copyright zerjioang. 2021 All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package bytes

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/zerjioang/zgo/assert"
	"testing"
)

func TestEndianess(t *testing.T) {
	t.Run("reload-get-endianess", func(t *testing.T) {
		checkEndian()
		e := Endianess()
		assert.True(t, e == binary.BigEndian || e == binary.LittleEndian)
	})
	t.Run("get-endianess", func(t *testing.T) {
		e := Endianess()
		assert.True(t, e == binary.BigEndian || e == binary.LittleEndian)
	})
}

func TestBytesToHex(t *testing.T) {
	t.Run("bytes-to-hex", func(t *testing.T) {
		str := BytesToHex([]byte("foo-bar"))
		assert.Equal(t, str, "666f6f2d626172")
	})
	t.Run("bytes-to-hex-nil", func(t *testing.T) {
		str := BytesToHex(nil)
		assert.Equal(t, str, "")
	})
	t.Run("bytes-to-hex-char", func(t *testing.T) {
		str := BytesToHex([]byte("5"))
		assert.Equal(t, str, "35")
	})
}

func TestUint32ToBytes(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		raw, err := uint32ToBytes(256454565)
		assert.NoError(t, err)
		assert.Equal(t, raw, []byte{0xa5, 0x2f, 0x49, 0xf})
	})
	t.Run("custom", func(t *testing.T) {
		raw := Uint32ToBytes(256454565)
		assert.Equal(t, raw, [4]byte{0xa5, 0x2f, 0x49, 0xf})
	})
}

func TestOtherUtils(t *testing.T) {
	t.Run("removeDoubleQuotes", func(t *testing.T) {
		str := RemoveDoubleQuotes([]byte(`hello "foo-bar"`))
		assert.Equal(t, string(str), "hello 'foo-bar'")
	})
}

func BenchmarkUtils(b *testing.B) {
	b.Run("default", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = uint32ToBytes(256454565)
		}
	})
	b.Run("custom", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Uint32ToBytes(256454565)
		}
	})
	b.Run("bytes-to-hex-zero", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = BytesToHex(nil)
		}
	})
	b.Run("bytes-to-hex-char", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		msg := []byte("6")
		for i := 0; i < b.N; i++ {
			_ = BytesToHex(msg)
		}
	})
	b.Run("bytes-to-hex-raw", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		msg := []byte("foo-bar")
		for i := 0; i < b.N; i++ {
			_ = BytesToHex(msg)
		}
	})
	b.Run("bytes-to-hex-stdlib", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		msg := []byte("foo-bar")
		for i := 0; i < b.N; i++ {
			_ = hex.EncodeToString(msg)
		}
	})
	b.Run("removeDoubleQuotes", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = RemoveDoubleQuotes([]byte(`hello "foo-bar"`))
		}
	})
}
