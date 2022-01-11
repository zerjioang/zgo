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
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenSecret []byte

type JwtCustomClaims struct {
	jwt.StandardClaims
	Metadata map[string]interface{} `json:"metadata"`
}

// CreateToken creates a new JWT token signed with given secret
func CreateToken(secret TokenSecret, claimFiller func(map[string]interface{}) map[string]interface{}) (string, error) {
	meta := map[string]interface{}{}
	// Set custom claims
	claim := &JwtCustomClaims{}
	claim.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
	}
	meta = claimFiller(meta)
	claim.Metadata = meta
	// Create token with claims and metadata
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// Generate encoded token and send it as response.
	return token.SignedString([]byte(secret))
}

// CreateRecoveryToken creates a new JWT token signed with given secret to start recovery process
func CreateRecoveryToken(secret TokenSecret, userId string, email string) (string, error) {
	return CreateToken(secret, func(m map[string]interface{}) map[string]interface{} {
		m["email"] = email
		m["id"] = userId
		m["recover"] = true
		return m
	})
}

func ParseToken(str string, secret TokenSecret) (*jwt.Token, error) {
	return jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid authentication method")
		}
		return []byte(secret), nil
	})
}
