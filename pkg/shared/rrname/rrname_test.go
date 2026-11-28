/*
Copyright 2026 Jordi Gil.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package rrname_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jordigilh/kubernaut/pkg/shared/rrname"
)

var _ = Describe("New", func() {
	It("UT-RRNAME-001: uses the first 12 fingerprint chars and a 12-hex suffix", func() {
		name := rrname.New("a1b2c3d4e5f6a7b8c9d0e1f2")
		Expect(name).To(MatchRegexp(`^rr-a1b2c3d4e5f6-[0-9a-f]{12}$`))
	})

	It("UT-RRNAME-002: keeps a fingerprint shorter than 12 chars as is", func() {
		Expect(rrname.New("abc123")).To(MatchRegexp(`^rr-abc123-[0-9a-f]{12}$`))
	})

	It("UT-RRNAME-003: gives different names for the same fingerprint", func() {
		seen := map[string]struct{}{}
		for range 1000 {
			seen[rrname.New("a1b2c3d4e5f6a7b8")] = struct{}{}
		}
		Expect(seen).To(HaveLen(1000))
	})
})
