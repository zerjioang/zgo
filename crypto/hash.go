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
	"crypto/sha256"
	"github.com/zerjioang/zgo/bytes"
)

// GenerateDigest Generates a request body sha-256 hash
// sha-256=e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
func GenerateDigest(reqBody []byte) (string, [32]byte) {
	hash := sha256.Sum256(reqBody)
	return bytes.BytesToHex(hash[:]), hash
}
