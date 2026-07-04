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

package ai

import "context"

// Provider defines the capabilities any LLM backend must have to work with Sonde.
type Provider interface {
	// Summary takes a raw wall of logs and returns a human-readable analysis.
	Summary(ctx context.Context, logs string) (string, error)

	// StreamSummary does the same, but yields tokens in real-time to the terminal.
	StreamSummary(ctx context.Context, logs string, callback func(string)) error
}
