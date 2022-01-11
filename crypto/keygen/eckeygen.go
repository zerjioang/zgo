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
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"math/big"
	"time"
)

var (
	errNotInCurve = errors.New("generated curve points are not valid")
	errNilKey     = errors.New("cannot export a nil key")
)

// GenerateECKeys Generates strong P-256 key pair for development or production envs
func GenerateECKeys(generateFiles bool) (*ecdsa.PrivateKey, error) {
	// generates a random P-256 ECDSA private key. secp256r1
	k, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	if !elliptic.P256().IsOnCurve(k.X, k.Y) {
		return nil, errNotInCurve
	}
	pKraw, err := exportPrivate(k)
	if err != nil {
		return nil, err
	}
	pRaw, err := exportPublic(&k.PublicKey)
	if err != nil {
		return nil, err
	}
	// now generate a certificate too
	certpemRaw, certErr := generateCert(k)
	// convert pem bytes to base64
	certpemRaw64 := encodeBase64(certpemRaw)
	if certErr != nil {
		return nil, certErr
	}
	if generateFiles {
		// generated_private.pem == server_ec_priv
		// generated_pub.pem == server_ec_pub
		_ = ioutil.WriteFile("server_ec_pub", pRaw, 0644)
		_ = ioutil.WriteFile("server_ec_priv", pKraw, 0644)
		_ = ioutil.WriteFile("server_ec_cert_ca", certpemRaw64, 0644)
	}
	return k, nil
}

// exportPrivate Converts Private ECDSA Key to PEM string
func exportPrivate(privkey *ecdsa.PrivateKey) ([]byte, error) {
	if privkey == nil {
		return nil, errNilKey
	}
	privkeyBytes, err := x509.MarshalECPrivateKey(privkey)
	if err != nil {
		return nil, err
	}
	pemBlock := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: privkeyBytes,
		},
	)
	return encodeBase64(pemBlock), nil
}

// exportPublic Converts Public ECDSA Key to PEM string
func exportPublic(pubkey *ecdsa.PublicKey) ([]byte, error) {
	if pubkey == nil {
		return nil, errNilKey
	}
	pubkeyBytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return nil, err
	}
	pemBlock := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: pubkeyBytes,
		},
	)
	return encodeBase64(pemBlock), nil
}

// encodeBase64 Encode byte array to base64 array
func encodeBase64(src []byte) []byte {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(buf, src)
	return buf
}

/*
Encode byte array to hex array
*/
func encodeHex(src []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}

/*
Generates a PEM encoded certificate from previously generated keys
*/
func generateCert(priv *ecdsa.PrivateKey) ([]byte, error) {
	rng, err := newSerialNumber()
	if err != nil {
		return nil, err
	}
	template := x509.Certificate{
		Version:      3,
		SerialNumber: rng,
		Subject: pkix.Name{
			Country:            []string{"US"},
			Organization:       []string{"Company, INC."},
			OrganizationalUnit: []string{"Certificates"},
			Locality:           []string{"San Francisco"},
			Province:           []string{""},
			StreetAddress:      []string{"Golden Gate Bridge"},
			PostalCode:         []string{"94016"},
			SerialNumber:       hex.EncodeToString(rng.Bytes()),
			CommonName:         "DP3T",
		},
		EmailAddresses:     []string{"support@domain.tld"},
		NotBefore:          time.Now(),
		NotAfter:           time.Now().AddDate(10, 0, 0),
		IsCA:               true,
		SignatureAlgorithm: x509.ECDSAWithSHA256,
		SubjectKeyId:       bigIntHash(priv.D),
		AuthorityKeyId:     bigIntHash(priv.D),
		//KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		//ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		// PublicKey: &priv.PublicKey,
	}

	/*
	   hosts := strings.Split(*host, ",")
	   for _, h := range hosts {
	   	if ip := net.ParseIP(h); ip != nil {
	   		template.IPAddresses = append(template.IPAddresses, ip)
	   	} else {
	   		template.DNSNames = append(template.DNSNames, h)
	   	}
	   }
	   if *isCA {
	   	template.IsCA = true
	   	template.KeyUsage |= x509.KeyUsageCertSign
	   }
	*/

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return nil, err
	}
	out := &bytes.Buffer{}
	_ = pem.Encode(out, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	return out.Bytes(), nil
}

/*
newSerialNumber returns a new random serial number suitable
for use in a certificate.
*/
func newSerialNumber() (*big.Int, error) {
	// A serial number can be up to 20 octets in size.
	// https://tools.ietf.org/html/rfc5280#section-4.1.2.2
	n, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 8*20))
	if err != nil {
		return nil, err
	}
	return n, nil
}

/*
Generates a sha-256 hash signature of given big integer
*/
func bigIntHash(n *big.Int) []byte {
	h := sha256.Sum256(n.Bytes())
	return h[:]
}

func PrivateToPublicHex(priv *ecdsa.PrivateKey) string {
	pubkeyBytes, err := x509.MarshalPKIXPublicKey(&priv.PublicKey)
	if err == nil {
		return hex.EncodeToString(pubkeyBytes)
	}
	return ""
}
