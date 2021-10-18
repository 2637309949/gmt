package types

import (
	"encoding/json"
)

// {{toTitle .entity}} defined 文章
type {{toTitle .entity}} struct {
	// ID defined 主键
	ID int64 `xorm:"bigint64(20) notnull autoincr unique pk comment('主键') 'id'" json:"id" form:"id" xml:"id"`
}

// TableName defined TODO
func (m *{{toTitle .entity}}) TableName() string {
	return "{{.entity}}"
}

// Marshal defined TODO
func (r *{{toTitle .entity}}) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Unmarshal{{toTitle .entity}} defined TODO
func Unmarshal{{toTitle .entity}}(data []byte) ({{toTitle .entity}}, error) {
	var r {{toTitle .entity}}
	err := json.Unmarshal(data, &r)
	return r, err
}
