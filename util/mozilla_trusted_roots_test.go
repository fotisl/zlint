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
	"encoding/pem"
	"fmt"
	"testing"

	"github.com/zmap/zcrypto/x509"
)

var testCertificates = []string{
	`-----BEGIN CERTIFICATE-----
MIIFazCCA1OgAwIBAgIRAIIQz7DSQONZRGPgu2OCiwAwDQYJKoZIhvcNAQELBQAw
TzELMAkGA1UEBhMCVVMxKTAnBgNVBAoTIEludGVybmV0IFNlY3VyaXR5IFJlc2Vh
cmNoIEdyb3VwMRUwEwYDVQQDEwxJU1JHIFJvb3QgWDEwHhcNMTUwNjA0MTEwNDM4
WhcNMzUwNjA0MTEwNDM4WjBPMQswCQYDVQQGEwJVUzEpMCcGA1UEChMgSW50ZXJu
ZXQgU2VjdXJpdHkgUmVzZWFyY2ggR3JvdXAxFTATBgNVBAMTDElTUkcgUm9vdCBY
MTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBAK3oJHP0FDfzm54rVygc
h77ct984kIxuPOZXoHj3dcKi/vVqbvYATyjb3miGbESTtrFj/RQSa78f0uoxmyF+
0TM8ukj13Xnfs7j/EvEhmkvBioZxaUpmZmyPfjxwv60pIgbz5MDmgK7iS4+3mX6U
A5/TR5d8mUgjU+g4rk8Kb4Mu0UlXjIB0ttov0DiNewNwIRt18jA8+o+u3dpjq+sW
T8KOEUt+zwvo/7V3LvSye0rgTBIlDHCNAymg4VMk7BPZ7hm/ELNKjD+Jo2FR3qyH
B5T0Y3HsLuJvW5iB4YlcNHlsdu87kGJ55tukmi8mxdAQ4Q7e2RCOFvu396j3x+UC
B5iPNgiV5+I3lg02dZ77DnKxHZu8A/lJBdiB3QW0KtZB6awBdpUKD9jf1b0SHzUv
KBds0pjBqAlkd25HN7rOrFleaJ1/ctaJxQZBKT5ZPt0m9STJEadao0xAH0ahmbWn
OlFuhjuefXKnEgV4We0+UXgVCwOPjdAvBbI+e0ocS3MFEvzG6uBQE3xDk3SzynTn
jh8BCNAw1FtxNrQHusEwMFxIt4I7mKZ9YIqioymCzLq9gwQbooMDQaHWBfEbwrbw
qHyGO0aoSCqI3Haadr8faqU9GY/rOPNk3sgrDQoo//fb4hVC1CLQJ13hef4Y53CI
rU7m2Ys6xt0nUW7/vGT1M0NPAgMBAAGjQjBAMA4GA1UdDwEB/wQEAwIBBjAPBgNV
HRMBAf8EBTADAQH/MB0GA1UdDgQWBBR5tFnme7bl5AFzgAiIyBpY9umbbjANBgkq
hkiG9w0BAQsFAAOCAgEAVR9YqbyyqFDQDLHYGmkgJykIrGF1XIpu+ILlaS/V9lZL
ubhzEFnTIZd+50xx+7LSYK05qAvqFyFWhfFQDlnrzuBZ6brJFe+GnY+EgPbk6ZGQ
3BebYhtF8GaV0nxvwuo77x/Py9auJ/GpsMiu/X1+mvoiBOv/2X/qkSsisRcOj/KK
NFtY2PwByVS5uCbMiogziUwthDyC3+6WVwW6LLv3xLfHTjuCvjHIInNzktHCgKQ5
ORAzI4JMPJ+GslWYHb4phowim57iaztXOoJwTdwJx4nLCgdNbOhdjsnvzqvHu7Ur
TkXWStAmzOVyyghqpZXjFaH3pO3JLF+l+/+sKAIuvtd7u+Nxe5AW0wdeRlN8NwdC
jNPElpzVmbUq4JUagEiuTDkHzsxHpFKVK7q4+63SM1N95R1NbdWhscdCb+ZAJzVc
oyi3B43njTOQ5yOf+1CceWxG1bQVs5ZufpsMljq4Ui0/1lvh+wjChP4kqKOJ2qxq
4RgqsahDYVvTH9w7jXbyLeiNdd8XM2w9U/t7y0Ff/9yi0GE44Za4rF2LN9d11TPA
mRGunUHBcnWEvgJBQl9nJEiU0Zsnvgc/ubhPgXRR4Xq37Z0j4r7g1SgEEzwxA57d
emyPxgcYxn/eR44/KJ4EBs+lVDR3veyJm+kXQ99b21/+jh5Xos1AnX5iItreGCc=
-----END CERTIFICATE-----`,
	`-----BEGIN CERTIFICATE-----
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
-----END CERTIFICATE-----`,
}

func TestTrustedRoots(t *testing.T) {
	testCases := []struct {
		Name           string
		Certificate    string
		ExpectedResult bool
	}{
		{
			Name:           "Test trusted certificate",
			Certificate:    testCertificates[0],
			ExpectedResult: true,
		},
		{
			Name:           "Test untrusted certificate",
			Certificate:    testCertificates[1],
			ExpectedResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			fmt.Println(tc.Certificate)
			block, _ := pem.Decode([]byte(tc.Certificate))
			if block == nil {
				t.Error("cannot decode test certificate")
				return
			}
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				t.Error("cannot parse test certificate")
				return
			}
			result := IsInMozillaRootStore(cert)
			if result != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result)
			}
		})
	}
}
