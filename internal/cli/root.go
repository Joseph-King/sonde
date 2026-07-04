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

var RootCmd = &cobra.Command{
	Use:   "sonde",
	Short: "Sonde: AI-powered terminal and container log analyzer",
	Long:  `Sonde is a CLI tool that uses AI to analyze terminal and container logs.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if data is being piped into the root command
		if isStdinNotEmpty() {
			// Data is piped! E.g., "cat error.log | sonde"
			logs, err := readStdin()
			if err != nil {
				return fmt.Errorf("failed to read piped logs: %w", err)
			}

			if strings.TrimSpace(logs) == "" {
				return fmt.Errorf("piped input was empty")
			}

			fmt.Printf("\n--- Analyzing Piped Stream (Truncated) ---\n%s\n------------------------------------------\n", truncateLogs(logs, 10))

			// TODO: Pass logs straight to your internal/ai package here!
			fmt.Println("\n[Success] Sending piped data to AI...")
			return nil
		}

		// If no data was piped and no subcommands were called, just print the help menu
		return cmd.Help()
	},
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

// truncateLogs helper to show just the tail end of what was grabbed
func truncateLogs(logs string, maxLines int) string {
	lines := strings.Split(logs, "\n")
	if len(lines) <= maxLines {
		return logs
	}
	return "... [truncated] ...\n" + strings.Join(lines[len(lines)-maxLines:], "\n")
}

func Execute() {
	// Register subcommands
	RootCmd.AddCommand(LastCmd)
	RootCmd.AddCommand(VersionCmd)

	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
