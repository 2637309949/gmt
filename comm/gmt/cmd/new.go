package cmd

import "github.com/spf13/cobra"

var new = &cobra.Command{
	Use:   "new",
	Short: "New service",
	RunE: func(_ *cobra.Command, _ []string) error {
		return nil
	},
}
