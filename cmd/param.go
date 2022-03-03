package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/zdz1715/yaml-parser/pkg/file"
)

type paramOptions struct {
	Key string
}

var paramOptionsValue = paramOptions{}

var paramCmd = &cobra.Command{
	Use:   "param file",
	Short: "Extract parameters from a single yaml file",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 || len(args[0]) == 0 {
			return errors.New("Missing file path  ")
		}
		return file.ParseParam(args[0], paramOptionsValue.Key)
	},
}

func init() {
	rootCmd.AddCommand(paramCmd)

	paramCmd.Flags().StringVar(&paramOptionsValue.Key, "key", "", "Get the value of a key")
}
