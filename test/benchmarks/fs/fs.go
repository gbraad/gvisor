// Copyright 2020 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package fs holds benchmarks around filesystem performance.
package fs

import (
	"os"
	"testing"

	"gvisor.dev/gvisor/test/benchmarks/harness"
)

var h harness.Harness

// TestMain is the main method for package fs.
func TestMain(m *testing.M) {
	h.Init()
	os.Exit(m.Run())
}
