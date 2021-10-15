package cmd

import (
	"log"
	"os"

	"github.com/2637309949/gmt/comm/gmt/template/dist"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/spf13/cobra"
)

var new = &cobra.Command{
	Use:   "new",
	Short: "New service",
	RunE: func(_ *cobra.Command, _ []string) error {
		if err := vfsutil.Walk(dist.Assets, "/boilerplate", func(p string, fi os.FileInfo, err error) error {
			if err != nil {
				log.Printf("can't stat file %s: %v\n", p, err)
				return nil
			}
			if fi.IsDir() {
				return nil
			}
			return nil
		}); err != nil {
			return err
		}
		return nil
	},
}
