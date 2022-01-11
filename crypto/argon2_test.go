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

package crypto

import (
	"fmt"
	"log"
	"testing"
)

func TestArgonHashing(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		// Establish the parameters to use for Argon2.
		p := &params{
			memory:      64 * 1024,
			iterations:  3,
			parallelism: 2,
			saltLength:  16,
			keyLength:   32,
		}

		// Pass the plaintext password and parameters to our generateFromPassword
		// helper function.
		hash, salt, err := GenerateFromPassword("password123", p)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(hash)
		hashStr := HashToString(p, hash, salt)
		fmt.Println(hashStr)
		match, err := ComparePasswordAndHash("password123", hashStr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(match)
	})
}
