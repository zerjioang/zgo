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

package nettools

import (
	"fmt"
	"github.com/zerjioang/zgo/assert"
	"testing"
)

func TestGetOutboundIP(t *testing.T) {
	t.Run("local", func(t *testing.T) {
		ip := GetOutboundIP()
		assert.NotNil(t, ip)
		fmt.Println(ip)
	})
}

// BenchmarkGetOutboundIP/outbound-ip-12         	616442737	         1.912 ns/op	 523.01 MB/s	       0 B/op	       0 allocs/op
// BenchmarkGetOutboundIP/outbound-ip-12         	1000000000	         0.2971 ns/op	3365.78 MB/s	       0 B/op	       0 allocs/op
func BenchmarkGetOutboundIP(b *testing.B) {
	b.Run("outbound-ip", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = GetOutboundIP()
		}
	})
}
