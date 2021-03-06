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

package host

import (
	"github.com/zerjioang/zgo/assert"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("hostname", func(t *testing.T) {
		name := Name()
		assert.NotNil(t, name)
	})
	t.Run("hostname-reload", func(t *testing.T) {
		err := Reload()
		assert.NoError(t, err)
	})
}

func BenchmarkName(b *testing.B) {
	b.Run("std-hostname", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		b.SetBytes(1)
		for i := 0; i < b.N; i++ {
			_, _ = os.Hostname()
		}
	})
	b.Run("custom-hostname", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		b.SetBytes(1)
		for i := 0; i < b.N; i++ {
			_ = Name()
		}
	})
}
