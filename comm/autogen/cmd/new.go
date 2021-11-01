package cmd

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/2637309949/gmt/comm/autogen/template/dist"
	"github.com/2637309949/gmt/comm/autogen/util"
	"github.com/2637309949/gmt/comm/autogen/util/fs"
	"github.com/spf13/cobra"
	"github.com/xormplus/xorm"
	"github.com/xormplus/xorm/schemas"
)

var (
	source string
	table  string
	proto  string
)

type Field struct {
	Name    string
	Type    string
	Options string
	XORM    string
}

func (f *Field) Index(i int) int {
	return i + 1
}

type Proto struct {
	Name   string
	Fields []*Field
}

func toProtoFieldTypeNameBySql(f schemas.SQLType) string {
	switch f.Name {
	case "VARCHAR", "TEXT", "LONGTEXT", "CHAR", "MEDIUMTEXT", "TINYTEXT":
		return "string"
	case "DATETIME", "TIMESTAMP", "ENUM", "INT", "SMALLINT", "BIGINT", "TINYINT":
		return "int64"
	case "DECIMAL":
		return "string"
	case "BOOLEAN":
		return "bool"
	case "FLOAT", "DOUBLE":
		return "float64"
	case "MEDIUMBLOB", "BLOB":
		return "string"
	}
	return "string"
}

func table2Proto() *Proto {
	t := Proto{
		Name:   "entity",
		Fields: []*Field{{Name: "id", Type: "uint64", XORM: "bigint64(20) notnull autoincr unique pk comment('主键') 'id'", Options: fmt.Sprintf(`[json_name="%v"]`, "id")}},
	}
	if table == "" || source == "" {
		return &t
	}
	engine, err := xorm.NewEngine("mysql", source)
	if err != nil {
		log.Println(err)
		return &t
	}
	if err := engine.Ping(); err != nil {
		log.Println(err)
		return &t
	}
	metas, _ := engine.DBMetas()
	for _, tb := range metas {
		if tb.Name == table {
			t.Name = tb.Name
			for _, col := range tb.Columns() {
				field := Field{Name: col.Name, Type: toProtoFieldTypeNameBySql(col.SQLType), Options: fmt.Sprintf(`[json_name="%v"]`, col.Name)}
				t.Fields = append(t.Fields, &field)
				switch col.SQLType.Name {
				case "VARCHAR", "TEXT":
					field.XORM = fmt.Sprintf("%v", strings.ToLower(col.SQLType.Name))
					field.XORM = fmt.Sprintf("%v(%v)", field.XORM, col.SQLType.DefaultLength)
				case "DATETIME":
					field.XORM = fmt.Sprintf("%v", strings.ToLower(col.SQLType.Name))
				case "ENUM", "INT", "BIGINT", "TINYINT", "SMALLINT", "MEDIUMINT", "INTEGER":
					if col.SQLType.Name == "ENUM" {
						field.XORM = fmt.Sprintf("int(%v)", 10)
					} else {
						field.XORM = fmt.Sprintf("%v", strings.ToLower(col.SQLType.Name))
						if col.SQLType.DefaultLength != 0 {
							field.XORM = fmt.Sprintf("%v(%v)", field.XORM, col.SQLType.DefaultLength)
						}
					}
				case "BOOLEAN":
					field.XORM = fmt.Sprintf("%v", strings.ToLower(col.SQLType.Name))
				case "MEDIUMBLOB", "BLOB":
					field.XORM = fmt.Sprintf("%v", strings.ToLower(col.SQLType.Name))
					if col.SQLType.DefaultLength != 0 {
						field.XORM = fmt.Sprintf("%v(%v)", field.XORM, col.SQLType.DefaultLength)
					}
				case "FLOAT", "DOUBLE":
					field.XORM = fmt.Sprintf("%v", strings.ToLower(col.SQLType.Name))
					if col.SQLType.DefaultLength != 0 {
						if col.SQLType.DefaultLength2 != 0 {
							field.XORM = fmt.Sprintf("%v(%v,%v)", field.XORM, col.SQLType.DefaultLength, col.SQLType.DefaultLength2)
						} else {
							field.XORM = fmt.Sprintf("%v(%v)", field.XORM, col.SQLType.DefaultLength)
						}
					}
				case "DECIMAL":
					field.XORM = fmt.Sprintf("%v", strings.ToLower(col.SQLType.Name))
					field.XORM = fmt.Sprintf("%v(%v,%v)", field.XORM, col.SQLType.DefaultLength, col.SQLType.DefaultLength2)
				}
				if col.IsPrimaryKey {
					field.XORM = fmt.Sprintf("%v %v", field.XORM, "pk")
				}
				if !col.Nullable {
					field.XORM = fmt.Sprintf("%v %v", field.XORM, "notnull")
				}
				if col.IsAutoIncrement {
					field.XORM = fmt.Sprintf("%v %v", field.XORM, "autoincr")
				}
				if col.Default != "" {
					field.XORM = fmt.Sprintf(`%v default(%v)`, field.XORM, col.Default)
				}
				field.XORM = fmt.Sprintf("%v", field.XORM)
			}
		}
	}
	return &t
}

func buildBoilerplate(name string) func(path string, info os.FileInfo, rs io.ReadSeeker, _ error) error {
	return func(path string, info os.FileInfo, rs io.ReadSeeker, _ error) error {
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
			"root":  root,
			"name":  name,
			"proto": table2Proto(),
		})
	}
}

func buildProto(name string) func(path string, info os.FileInfo, rs io.ReadSeeker, _ error) error {
	return func(path string, info os.FileInfo, rs io.ReadSeeker, _ error) error {
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
			"root":  root,
			"name":  name,
			"proto": table2Proto(),
		})
	}
}

var new = &cobra.Command{
	Use:   "new",
	Short: "New service",
	RunE: func(_ *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide service name")
		}
		name := args[0]

		if isExist(name) {
			return fmt.Errorf("directory already exists %v", name)
		}
		fs.WalkFiles(dist.Assets, "/boilerplate", buildBoilerplate(name))
		if isExist(proto + "/" + name) {
			return fmt.Errorf("directory already exists %v", proto+"/"+name)
		}
		fs.WalkFiles(dist.Assets, "/proto", buildProto(name))
		return nil
	},
}
