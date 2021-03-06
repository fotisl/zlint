/*
 * ZLint Copyright 2018 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package util

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/zmap/zcrypto/x509"
)

// IsInMozillaRootStore checks whether the SPKI of a certificate is trusted
// by Mozilla.
func IsInMozillaRootStore(cert *x509.Certificate) bool {
	spki := sha256.Sum256(cert.RawSubjectPublicKeyInfo)
	encSPKI := hex.EncodeToString(spki[:])
	for _, s := range mozillaTrustedSPKIs {
		if s == encSPKI {
			return true
		}
	}

	return false
}
