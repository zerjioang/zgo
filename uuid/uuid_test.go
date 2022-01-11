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

package uuid

import (
	"fmt"
	"github.com/zerjioang/zgo/assert"
	"testing"
)

func TestUUID(t *testing.T) {
	t.Run("set-id", func(t *testing.T) {
		u := New()
		assert.NotNil(t, u)
	})
	t.Run("test-collisions", func(t *testing.T) {
		var holder = map[string]int{}
		for i := 0; i < 5000000; i++ {
			k := New()
			holder[k] = holder[k] + 1
		}
		for k, v := range holder {
			if v != 1 {
				fmt.Println(k)
				assert.Fail(t, "collision")
			}
		}
	})
}

func BenchmarkUUID(b *testing.B) {
	b.Run("new-uuid", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = New()
		}
	})
	b.Run("new-uuid-half", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Half()
		}
	})
}
