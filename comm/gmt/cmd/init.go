package cmd

import (
	"errors"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/cobra"
)

func removeAll(path ...string) error {
	for i := range path {
		os.RemoveAll(path[i])
	}
	return nil
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

var initd = &cobra.Command{
	Use:   "init",
	Short: "init a micro collection",
	RunE: func(_ *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide the collection name")
		}
		if isExist(args[0]) {
			return errors.New("directory already exists")
		}
		cmd := exec.Command("git", "clone", "--depth=1", "https://hub.fastgit.org/2637309949/gmt.git", args[0])
		_, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		removeAll(path.Join(args[0], ".git"), path.Join(args[0], "docs"), path.Join(args[0], "README.md"))
		return nil
	},
}
