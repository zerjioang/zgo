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
	"crypto/ecdsa"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt"
)

// SignToken Generates a signature using P-256 ECDSA over passed arbitrary data
// and return the signature as JWT ES-256
func SignToken(sha256hash [32]byte, privkey *ecdsa.PrivateKey, retentionDays int, protectedHeaders interface{}) (string, error) {
	//logger.Logger.Debug("signing digest with server private key")
	tt := time.Now()
	payload := jwt.MapClaims{
		"content-hash": base64.StdEncoding.EncodeToString(sha256hash[:]),
		"hash-alg":     "sha-256",
		"iss":          "api",                                  // issuer
		"iat":          tt,                                     // issued at
		"exp":          tt.AddDate(0, 0, retentionDays).Unix(), // expiration
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, payload)
	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(privkey)
}
