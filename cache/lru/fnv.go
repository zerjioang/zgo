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

package lru

const (
	offset32 = 0x811c9dc5
	prime32  = 0x1000193
	offset64 = 0xcbf29ce484222325
	prime64  = 0x100000001b3
)

// Hash64a takes a string and
// returns a 64 bit FNV-1a.
func Hash64a(s string) uint64 {
	var h uint64 = offset64
	for _, c := range s {
		h ^= uint64(c)
		h *= prime64
	}
	return h
}

// Hash64aSlice takes a []byte and
// returns a 64 bit FNV-1a.
func Hash64aSlice(s []byte) uint64 {
	var h uint64 = offset64
	for _, c := range s {
		h ^= uint64(c)
		h *= prime64
	}
	return h
}
