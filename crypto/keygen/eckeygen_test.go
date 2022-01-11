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

package keygen

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/zerjioang/zgo/assert"
	"testing"
)

func TestGenerateECKeys(t *testing.T) {
	t.Run("generate-no-files", func(t *testing.T) {
		k, err := GenerateECKeys(false)
		assert.NoError(t, err)
		assert.NotNil(t, k)
	})
	t.Run("generate-files", func(t *testing.T) {
		k, err := GenerateECKeys(true)
		assert.NoError(t, err)
		assert.NotNil(t, k)
	})
	t.Run("try-export-nil-private-key", func(t *testing.T) {
		raw, err := exportPrivate(nil)
		assert.Nil(t, raw)
		assert.Equal(t, err, errNilKey)
	})
	t.Run("try-export-nil-public-key", func(t *testing.T) {
		raw, err := exportPublic(nil)
		assert.Nil(t, raw)
		assert.Equal(t, err, errNilKey)
	})
	t.Run("try-export-bad-private-key", func(t *testing.T) {
		raw, err := exportPrivate(&ecdsa.PrivateKey{})
		assert.Nil(t, raw)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), "x509: unknown elliptic curve")
	})
	t.Run("generate-cert", func(t *testing.T) {
		k, err := GenerateECKeys(false)
		assert.NoError(t, err)
		assert.NotNil(t, k)
		cert, err := generateCert(k)
		t.Log(cert)
		assert.NoError(t, err)
		assert.NotNil(t, cert)
	})
	t.Run("generate-cert-get-public-key", func(t *testing.T) {
		k, err := GenerateECKeys(false)
		assert.NoError(t, err)
		assert.NotNil(t, k)
		cert, err := generateCert(k)
		t.Log(string(cert))
		assert.NoError(t, err)
		assert.NotNil(t, cert)
		// decode pem and get public key
		block, _ := pem.Decode(cert)
		var certitem *x509.Certificate
		certitem, parseErr := x509.ParseCertificate(block.Bytes)
		assert.NoError(t, parseErr)
		assert.NotNil(t, certitem)
		pubk := certitem.PublicKey.(*ecdsa.PublicKey)
		assert.NotNil(t, certitem)
		t.Log(pubk.X)
		t.Log(pubk.Y)
		assert.Equal(t, pubk.X, k.PublicKey.X)
		assert.Equal(t, pubk.Y, k.PublicKey.Y)
	})
}
