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

// Package rrname generates RemediationRequest names per DD-AUDIT-CORRELATION-002.
// Every component that creates a RemediationRequest (Gateway, API Frontend)
// uses New so the name format lives in one place.
package rrname

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

const (
	// FingerprintPrefixLen is how many fingerprint chars go into the name.
	FingerprintPrefixLen = 12
	// SuffixLen is how many random hex chars (dashes removed) go into the name.
	SuffixLen = 12
)

// New returns rr-{fingerprint[:12]}-{uuid-hex[:12]}. The fingerprint prefix
// gives human-readable correlation for debugging. The suffix is the first 12
// hex chars of a random UUID with dashes removed (48 bits), so two names with
// the same prefix collide with probability ~2^-48 per pair.
func New(fingerprint string) string {
	prefix := fingerprint
	if len(prefix) > FingerprintPrefixLen {
		prefix = prefix[:FingerprintPrefixLen]
	}
	suffix := strings.ReplaceAll(uuid.New().String(), "-", "")[:SuffixLen]
	return fmt.Sprintf("rr-%s-%s", prefix, suffix)
}
