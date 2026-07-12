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
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var PipeCmd = &cobra.Command{
	Use:   "pipe",
	Short: "Print the build version and architecture information for sonde",
	RunE: func(cmd *cobra.Command, args []string) error {
		logs, err := readStdin()
		if err != nil {
			return fmt.Errorf("failed to read piped logs: %w", err)
		}

		if strings.TrimSpace(logs) == "" {
			return fmt.Errorf("piped input was empty")
		}

		return nil
	},
}

func pipe() (string, error) {
	if !isStdinNotEmpty() {
		return "", fmt.Errorf("piped input was empty")
	}

	logs, err := readStdin()
	if err != nil {
		return "", fmt.Errorf("failed to read piped logs: %w", err)
	}

	if strings.TrimSpace(logs) == "" {
		return "", fmt.Errorf("piped input was empty")
	}

	return logs, nil
}

func isStdinNotEmpty() bool {
	stat, _ := os.Stdin.Stat()
	return ((stat.Mode() & os.ModeCharDevice) == 0)
}

// readStdin reads all text from standard input
func readStdin() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	var output strings.Builder

	for {
		input, err := reader.ReadString('\n')

		if err == io.EOF {
			output.WriteString(input)
			break
		}

		if err != nil {
			return "", err
		}

		output.WriteString(input)
	}

	return output.String(), nil
}
