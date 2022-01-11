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
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"github.com/zerjioang/zgo/bytes"
)

func Hmac(data, secret []byte) string {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, secret)
	// Write Data to it
	h.Write(data)
	// Get result and encode as hexadecimal string
	return bytes.BytesToHex(h.Sum(nil))
}

func HmacVerify(hash string, data, secret []byte) (bool, error) {
	// decode the hash
	givenHash, err := hex.DecodeString(hash)
	if err != nil {
		return false, err
	}
	// compute the hash again over given data
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, secret)
	// Write Data to it
	h.Write(data)
	// Get result and encode as hexadecimal string
	createdHash := h.Sum(nil)
	return subtle.ConstantTimeCompare(givenHash, createdHash) == 1, nil
}
