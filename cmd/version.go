package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zdz1715/yaml-parser/pkg/version"
	"os"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: `Print the version information.`,
	Long:  `Print the version information.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, _ = fmt.Fprintf(os.Stdout, "%s", version.Get())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
