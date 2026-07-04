package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// These variables are populated at build time using -ldflags
var (
	Version   = "dev"
	CommitSHA = "unknown"
	BuildDate = "unknown"
)

var VersionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Print the build version and architecture information for Fathom",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Fathom Analyzer %s\n", Version)
		fmt.Printf("Commit: %s\n", CommitSHA)
		fmt.Printf("Build Date: %s\n", BuildDate)
	},
}
