package types

import (
	"encoding/json"
)

// {{toTitle .proto.Name}} defined 文章
type {{toTitle .proto.Name}} struct {
	{{- range .proto.Fields}}
	{{toTitle .Name}} {{.Type}} `xorm: "{{.XORM}}" json: "{{.Name}}" xml:"{{.Name}}"`
	{{- end}}
}

// TableName defined TODO
func (m *{{toTitle .proto.Name}}) TableName() string {
	return "{{.proto.Name}}"
}

// Marshal defined TODO
func (r *{{toTitle .proto.Name}}) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Unmarshal{{toTitle .proto.Name}} defined TODO
func Unmarshal{{toTitle .proto.Name}}(data []byte) ({{toTitle .proto.Name}}, error) {
	var r {{toTitle .proto.Name}}
	err := json.Unmarshal(data, &r)
	return r, err
}
