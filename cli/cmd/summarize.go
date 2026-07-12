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
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	ctx := &CommandContext{}

	SummarizeCmd := &cobra.Command{
		Use:     "summarize",
		Aliases: []string{"s"},
		Short:   "Summarize data from various sources",
		RunE: func(cmd *cobra.Command, args []string) error {
			dryRun, _ := cmd.Flags().GetBool("dry-run")

			logs, err := pipe()

			if err != nil {
				return err
			}

			shell := os.Getenv("SHELL")
			cm, err := getLastCommand(shell, 1)
			if err != nil {
				return err
			}

			split := strings.Split(cm, "|")
			if strings.Contains(split[0], "sonde") {
				return fmt.Errorf("piped command includes sonde, terminating to prevent issues: %s", split[0])
			}

			if logs != "" {
				data := "command: " + split[0] + "\n" + "output: " + logs
				if dryRun {
					var prompt = constructPrompt("pipe", data)
					fmt.Printf("[DRY-RUN]: %s\n", prompt)
				} else {
					var prompt = constructPrompt("pipe", data)
					fmt.Printf("[SUMMARY]: %s\n", strings.ToUpper(prompt))
				}
			}
			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			dryRun, _ := cmd.Flags().GetBool("dry-run")

			if ctx.Data != "" {
				if dryRun {
					var prompt = constructPrompt(ctx.Source, ctx.Data)
					fmt.Printf("[DRY-RUN]: %s\n", prompt)
				} else {
					var prompt = constructPrompt(ctx.Source, ctx.Data)
					fmt.Printf("[SUMMARY]: %s\n", strings.ToUpper(prompt))
				}
			}
			return nil
		},
	}

	SummarizeCmd.AddCommand(lastCmd(ctx))

	RootCmd.AddCommand(SummarizeCmd)
}

func constructPrompt(source string, data string) string {
	var initSummaryPrompt = "I have a user that requires a summary of the following action:\n"
	var restOfPrompt = "Source: " + source + "\n" + data
	return fmt.Sprintf("%s %s%s", initSondePrompt, initSummaryPrompt, restOfPrompt)
}
