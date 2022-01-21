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
	"fmt"
	"sync/atomic"
)

type Auint32 uint32

// Add increments the base value given delta
func (u *Auint32) Add(v uint32) {
	atomic.AddUint32((*uint32)(u), v)
}

// Set sets the uint32 value
func (u *Auint32) Set(v uint32) {
	atomic.StoreUint32((*uint32)(u), v)
}

// Get returns uint32 value
func (u *Auint32) Get() uint32 {
	return atomic.LoadUint32((*uint32)(u))
}

// Increment increments counter value and returns the new value
func (u *Auint32) Increment(v uint32) uint32 {
	atomic.AddUint32((*uint32)(u), v)
	return atomic.LoadUint32((*uint32)(u))
}

// Equals compares given value v with u value
func (u *Auint32) Equals(v uint32) bool {
	return u.Get() == v
}

func (u *Auint32) NotEquals(v uint32) bool {
	return !u.Equals(v)
}

// String implements Stringer interface
func (u *Auint32) String() string {
	return fmt.Sprintf("%d", u.Get())
}

// EqualsIn compares given value v with passed value list
// if value u is found in the list, returns true
// false otherwise
func (u *Auint32) EqualsIn(v ...uint32) bool {
	for i := 0; i < len(v); i++ {
		if u.Get() == v[i] {
			return true
		}
	}
	return false
}

// NotEqualsIn compares given value v with passed value list
// if value u is found in the list, returns false
// true otherwise
func (u *Auint32) NotEqualsIn(v ...uint32) bool {
	return !u.EqualsIn(v...)
}
