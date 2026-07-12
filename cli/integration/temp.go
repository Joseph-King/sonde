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

package integration

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func getLastBash(howFar int) (string, error) {
	cmd := exec.Command("bash", "-c", "history", fmt.Sprintf("-%d", howFar))
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

func getLastZsh(howFar int) (string, error) {
	cmd := exec.Command("zsh", "-c", fmt.Sprintf("fc -l -n %d | sed '/^ *[^ ]/!d'", howFar))
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

func getLastFish(howFar int) (string, error) {
	cmd := exec.Command("fish", "-c", fmt.Sprintf("echo $history[%d]", howFar))
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

func getLastCommand(shell string, howFar int) (string, error) {
	if strings.Contains(shell, "zsh") {
		return getLastZsh(howFar)
	} else if strings.Contains(shell, "bash") {
		return getLastBash(howFar)
	} else if strings.Contains(shell, "fish") {
		return getLastFish(howFar)
	}

	return "", fmt.Errorf("unsupported shell: %s", os.Getenv("SHELL"))
}
