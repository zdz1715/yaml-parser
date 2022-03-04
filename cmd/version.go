package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zdz1715/yaml-parser/pkg/version"
	"os"
)

type versionOptions struct {
	All bool
}

var versionOptionsValue = versionOptions{}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: `Print the version information.`,
	Long:  `Print the version information.`,
	Run: func(cmd *cobra.Command, args []string) {
		info := ""
		if versionOptionsValue.All {
			info = version.Get().String()
		} else {
			info = version.GetVersion()
		}
		_, _ = fmt.Fprintf(os.Stdout, "%s", info)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	versionCmd.Flags().BoolVarP(&versionOptionsValue.All, "all", "a", false, "Get all information\n\n")
}
