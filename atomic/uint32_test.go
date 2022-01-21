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

package atomic

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtomic32(t *testing.T) {
	t.Run("noop", func(t *testing.T) {
		t.Log("boilerplate test to avoid FAIL in travis CI due to missing _test file")
	})
	t.Run("increment", func(t *testing.T) {
		var counter Auint32
		var wg = sync.WaitGroup{}
		wg.Add(100)
		for i := 0; i < 100; i++ {
			go func() {
				counter.Increment(1)
				wg.Done()
			}()
		}
		wg.Wait()
		t.Log(counter)
		assert.Equal(t, counter.Get(), uint32(100))
	})
	t.Run("set", func(t *testing.T) {
		var counter Auint32
		var wg = sync.WaitGroup{}
		wg.Add(100)
		for i := 0; i < 100; i++ {
			go func(v int) {
				if uint32(v) > counter.Get() {
					counter.Set(uint32(v))
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
		t.Log(counter)
		assert.Equal(t, counter.Get(), uint32(99))
	})
	t.Run("add", func(t *testing.T) {
		var counter Auint32
		var wg = sync.WaitGroup{}
		wg.Add(100)
		for i := 0; i < 100; i++ {
			go func(v int) {
				counter.Add(2)
				wg.Done()
			}(i)
		}
		wg.Wait()
		t.Log(counter)
		assert.Equal(t, counter.Get(), uint32(200))
	})
	t.Run("string", func(t *testing.T) {
		var counter Auint32
		counter.Set(uint32(50))
		assert.Equal(t, counter.String(), "50")
	})
	t.Run("equals", func(t *testing.T) {
		var counter Auint32
		counter.Set(uint32(50))
		assert.True(t, counter.Equals(50))
	})
	t.Run("not-equals", func(t *testing.T) {
		var counter Auint32
		counter.Set(uint32(50))
		assert.True(t, counter.NotEquals(60))
	})
	t.Run("equals-in", func(t *testing.T) {
		var counter Auint32
		counter.Set(uint32(50))
		assert.True(t, counter.EqualsIn(50, 30))
		counter.Set(uint32(30))
		assert.True(t, counter.EqualsIn(50, 30))
		counter.Set(uint32(40))
		assert.False(t, counter.EqualsIn(50, 30))
	})
	t.Run("not-equals-in", func(t *testing.T) {
		var counter Auint32
		counter.Set(uint32(40))
		assert.True(t, counter.NotEqualsIn(60, 50))
		counter.Set(uint32(50))
		assert.False(t, counter.NotEqualsIn(60, 50))
		counter.Set(uint32(60))
		assert.False(t, counter.NotEqualsIn(60, 50))
	})
}
