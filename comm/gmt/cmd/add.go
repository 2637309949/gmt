package cmd

import "github.com/spf13/cobra"

var add = &cobra.Command{
	Use:   "add",
	Short: "Add handler",
	RunE: func(_ *cobra.Command, _ []string) error {
		return nil
	},
}
