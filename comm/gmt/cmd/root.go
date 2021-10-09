package cmd

import (
	"github.com/spf13/cobra"
)

var Root = &cobra.Command{
	Use:               "gmt",
	Long:              `Code generation tool for go-micro`,
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

func init() {
	Root.AddCommand(new)
	Root.AddCommand(add)
}
