package cmd

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/2637309949/gmt/comm/gmt/template/dist"
	"github.com/2637309949/gmt/comm/gmt/util"
	"github.com/2637309949/gmt/comm/gmt/util/fs"
	"github.com/spf13/cobra"
)

var (
	source string
	table  string
	proto  string
)

var new = &cobra.Command{
	Use:   "new",
	Short: "New service",
	RunE: func(_ *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide service name")
		}
		name := args[0]
		if isExist(name) || isExist(proto+"/"+name) {
			return errors.New("directory already exists")
		}
		if err := fs.WalkFiles(dist.Assets, "/boilerplate", func(path string, info os.FileInfo, rs io.ReadSeeker, _ error) error {
			if info.IsDir() {
				return nil
			}
			fileBytes, _ := fs.ReadFile(dist.Assets, path)
			t := template.Must(template.New("boilerplate").Funcs(template.FuncMap{"toTitle": util.ToTitle, "camel2case": util.Camel2Case, "case2camel": util.Case2Camel}).Parse(string(fileBytes)))
			path = strings.ReplaceAll(path, "/boilerplate", name+"-service")
			root := filepath.Base(filepath.Dir(path))
			err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
			if err != nil {
				return nil
			}
			path = strings.ReplaceAll(path, "helloworld_handler.go", name+"_handler.go")
			path = strings.ReplaceAll(path, "helloworld_dbhandler.go", name+"_dbhandler.go")
			file, err := os.Create(path)
			if err != nil {
				return err
			}
			return t.Execute(file, map[string]interface{}{
				"root":   root,
				"name":   name,
				"entity": table,
			})
		}); err != nil {
			return err
		}
		if err := fs.WalkFiles(dist.Assets, "/proto", func(path string, info os.FileInfo, rs io.ReadSeeker, _ error) error {
			if info.IsDir() {
				return nil
			}
			fileBytes, _ := fs.ReadFile(dist.Assets, path)
			t := template.Must(template.New("proto").Funcs(template.FuncMap{"toTitle": util.ToTitle, "camel2case": util.Camel2Case, "case2camel": util.Case2Camel}).Parse(string(fileBytes)))
			path = strings.ReplaceAll(path, "/proto", proto+"/"+name)
			root := filepath.Base(filepath.Dir(path))
			err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
			if err != nil {
				return nil
			}
			path = strings.ReplaceAll(path, "helloworld.proto", name+".proto")
			path = strings.ReplaceAll(path, "article.go", table+".go")
			file, err := os.Create(path)
			if err != nil {
				return err
			}
			return t.Execute(file, map[string]interface{}{
				"root":   root,
				"name":   name,
				"entity": table,
			})
		}); err != nil {
			return err
		}
		return nil
	},
}
