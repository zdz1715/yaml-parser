package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/zdz1715/yml-parser/pkg/yml"
)

var splitCmd = &cobra.Command{
	Use:   "split file...",
	Short: "Dividing a YML file into multiple",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := validate(args); err != nil {
			return err
		}
		return yml.SplitByFiles(args)
	},
}

func validate(filepath []string) error {
	if len(filepath) == 0 {
		return errors.New("Missing file path ")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(splitCmd)
}
