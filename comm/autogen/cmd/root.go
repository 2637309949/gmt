package cmd

import (
	"github.com/spf13/cobra"
)

var Root = &cobra.Command{
	Use:               "autogen",
	Long:              `Code generation tool for go-micro`,
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

func init() {
	new.Flags().StringVarP(&source, "source", "s", "", "Specify the database address")
	new.Flags().StringVarP(&table, "table", "t", "entity", "Specify the table name")
	new.Flags().StringVarP(&proto, "proto", "p", "proto", "Specify the dir of proto")
	Root.AddCommand(new)
	Root.AddCommand(initd)
}
