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

import (
	cache "github.com/hashicorp/golang-lru"
)

type Cache struct {
	c *cache.Cache
}

func NewLRUCache(size uint) *Cache {
	ch, err := cache.New(int(size))
	if err != nil {
		panic(err)
	}
	c := Cache{
		c: ch,
	}
	return &c
}

// Add adds a value to the cache.  Returns true if an eviction occurred.
func (c *Cache) Add(key, value interface{}) (evicted bool) {
	return c.c.Add(key, value)
}

// Get looks up a key's value from the cache.
func (c *Cache) Get(key interface{}) (value interface{}, ok bool) {
	return c.c.Get(key)
}

// Delete removes the provided key from the cache.
func (c *Cache) Delete(key interface{}) {
	c.c.Remove(key)
}
