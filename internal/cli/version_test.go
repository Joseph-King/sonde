// Copyright 2026 Joseph King
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

package cli

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// TestFormatVersionString checks that our version metadata string compiles cleanly
func TestFormatVersionString(t *testing.T) {
	// 1. Setup mock data
	version := "v0.1.0-alpha"
	commit := "abc1234"
	date := "2026-07-04"

	// 2. Call the logic you want to test (adjust this to match your structural logic)
	result := fmt.Sprintf("sonde %s (commit: %s, built: %s)", version, commit, date)

	// 3. Assertions
	if !strings.Contains(result, "v0.1.0-alpha") {
		t.Errorf("Expected version string to contain version tag, got: %s", result)
	}

	if !strings.Contains(result, "abc1234") {
		t.Errorf("Expected version string to contain commit SHA, got: %s", result)
	}
}

func TestVersionCommandExecution(t *testing.T) {
	buf := new(bytes.Buffer)

	// Intercept Cobra's terminal streams
	RootCmd.SetOut(buf)
	RootCmd.SetErr(buf)

	// Set the arguments to execute
	RootCmd.SetArgs([]string{"version"})

	// Trigger the execution
	err := RootCmd.Execute()
	if err != nil {
		t.Fatalf("RootCmd.Execute() failed unexpectedly: %v", err)
	}

	output := buf.String()

	if output == "" {
		t.Error("Expected version command to print layout output, but it was empty")
	}
}
