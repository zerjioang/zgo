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

package timer

import (
	"github.com/zerjioang/zgo/assert"
	"testing"
)

func TestTimer(t *testing.T) {
	t.Run("now", func(t *testing.T) {
		assert.True(t, Now() > 0)
	})
	t.Run("time", func(t *testing.T) {
		assert.True(t, Time().Unix() > 0)
	})
}

func BenchmarkTimer(b *testing.B) {
	b.Run("now", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Now()
		}
	})
	b.Run("time", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Time()
		}
	})
}
