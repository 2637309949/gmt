package cmd

import "github.com/spf13/cobra"

var initd = &cobra.Command{
	Use:   "init",
	Short: "init a micro collection",
	RunE: func(_ *cobra.Command, _ []string) error {
		return nil
	},
}
