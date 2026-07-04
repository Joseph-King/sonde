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
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// LastCmd represents the 'last' command
var LastCmd = &cobra.Command{
	Use:   "last",
	Short: "Analyze the most recent terminal output or piped logs",
	Long:  `Sonde attempts to capture your last terminal buffer, then sends it to your AI provider for analysis.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var logs string
		var err error

		logs, err = captureTerminalBuffer()

		if err != nil {
			return fmt.Errorf("could not automatically capture terminal: %w. Try piping your command: 'cmd | sonde'", err)
		}

		if logs == "" {
			fmt.Println("Nothing captured... exiting")
			return nil
		}

		if strings.TrimSpace(logs) == "" {
			return fmt.Errorf("captured logs are empty")
		}

		// Print a snippet of what we caught for verification
		fmt.Printf("\n--- Captured Logs (Truncated) ---\n%s\n---------------------------------\n", truncateLogs(logs, 10))

		// TODO: Pass 'logs' to your internal/ai package here in the next step!
		fmt.Println("\n[Success] Ready to send to AI provider...")
		return nil
	},
}

func captureTerminalBuffer() (string, error) {
	// TO DO add tmux/xclip/wl-clipboard if applicable
	shell := os.Getenv("SHELL")

	lastCmd, err := getLastCommand(shell)
	if err != nil {
		return "", err
	} else if lastCmd == "" {
		return "", nil
	}

	if strings.Contains(lastCmd, "sonde") {
		return "", fmt.Errorf("previous command includes sonde, terminating to prevent issues: %s", lastCmd)
	}

	output, err := runLastCommand(shell, lastCmd)
	if err != nil {
		return "", err
	}
	return output, nil
}

func runLastCommand(shell, cmd string) (string, error) {
	// Prompt the user to confirm
	reader := bufio.NewReader(os.Stdin)
	_, _ = io.WriteString(os.Stdout, fmt.Sprintf("Would you like to re-execute and capture the following command:\n%s\n? (y/N): ", cmd))
	response, err := reader.ReadString('\n')

	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	response = strings.TrimSpace(response)
	if response != "y" && response != "Y" {
		return "", nil
	}

	// Execute the command and capture its output
	fullCmd := exec.Command(shell, "-c", cmd)
	output, err := fullCmd.CombinedOutput()

	if err != nil {
		return "", fmt.Errorf("failed to execute command: %w", err)
	}

	return string(output), nil
}

func getLastBash() (string, error) {
	cmd := exec.Command("bash", "-c", "history", "-1")
	output, err := cmd.CombinedOutput()

	if err != nil {
		return "", fmt.Errorf("failed to capture bash history: %w", err)
	}

	lastCommand := strings.TrimSpace(string(output))
	if len(lastCommand) == 0 {
		return "", fmt.Errorf("no commands found in bash history")
	}

	if strings.HasPrefix(lastCommand, "bash -c") {
		lastCommand = strings.TrimPrefix(lastCommand, "bash -c ")
	}

	return lastCommand, nil
}

func getLastZsh() (string, error) {
	cmd := exec.Command("zsh", "-c", fmt.Sprintf("fc -l -n %d | sed '/^ *[^ ]/!d'", 1))
	output, err := cmd.CombinedOutput()

	if err != nil {
		return "", fmt.Errorf("failed to capture zsh history: %w", err)
	}

	lastCommand := strings.TrimSpace(string(output))
	if len(lastCommand) == 0 {
		return "", fmt.Errorf("no commands found in zsh history")
	}

	if strings.HasPrefix(lastCommand, "zsh -c") {
		lastCommand = strings.TrimPrefix(lastCommand, "zsh -c ")
	}

	return lastCommand, nil
}

func getLastFish() (string, error) {
	cmd := exec.Command("fish", "-c", "echo $history[2]")
	output, err := cmd.CombinedOutput()

	if err != nil {
		return "", fmt.Errorf("failed to capture fish history: %w", err)
	}

	lastCommand := strings.TrimSpace(string(output))
	if len(lastCommand) == 0 {
		return "", fmt.Errorf("no commands found in fish history")
	}

	if strings.HasPrefix(lastCommand, "fish -c") {
		lastCommand = strings.TrimPrefix(lastCommand, "fish -c ")
	}

	return lastCommand, nil
}

func getLastCommand(shell string) (string, error) {
	if strings.Contains(shell, "zsh") {
		return getLastZsh()
	} else if strings.Contains(shell, "bash") {
		return getLastBash()
	} else if strings.Contains(shell, "fish") {
		return getLastFish()
	}

	return "", fmt.Errorf("unsupported shell: %s", os.Getenv("SHELL"))
}
