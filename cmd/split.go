package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/zdz1715/yaml-parser/pkg/file"
)

var splitCmd = &cobra.Command{
	Use:   "split file...",
	Short: "Dividing a YAML file into multiple",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := validateSplitArgs(args); err != nil {
			return err
		}
		return file.SplitByFiles(args)
	},
}

func validateSplitArgs(filepath []string) error {
	if len(filepath) == 0 {
		return errors.New("Missing file path ")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(splitCmd)
}
