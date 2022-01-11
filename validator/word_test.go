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

package validators

import (
	"github.com/zerjioang/zgo/assert"
	"testing"
)

func TestIsLetter(t *testing.T) {
	t.Run("is-letter", func(t *testing.T) {
		t.Run("false", func(t *testing.T) {
			assert.False(t, IsLetter("hyperledger_indy../$"))
		})
		t.Run("true", func(t *testing.T) {
			assert.True(t, IsLetter("charsetAZaz09"))
		})
	})
}

func BenchmarkIsLetter(b *testing.B) {
	b.Run("is-letter", func(b *testing.B) {
		b.Run("false", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = IsLetter("hyperledger_indy../$")
			}
		})
		b.Run("true", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = IsLetter("charsetAZaz09")
			}
		})
	})
}
