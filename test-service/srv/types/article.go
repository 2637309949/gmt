package types

import (
	"encoding/json"
	"time"
)

// Entity defined 文章
type Entity struct {
	// ID defined 主键
	ID int64 `xorm:"bigint64(20) notnull autoincr unique pk comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// Creater defined 创建人
	Creater int64 `xorm:"bigint64(20) notnull comment('创建人') 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateTime defined 创建时间
	CreateTime time.Time `xorm:"datetime comment('创建时间') 'create_time'" json:"create_time" form:"create_time" xml:"create_time"`
	// Updater defined 最后更新人
	Updater int64 `xorm:"bigint64(20) notnull comment('最后更新人') 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateTime defined 最后更新时间
	UpdateTime time.Time `xorm:"datetime comment('最后更新时间') 'update_time'" json:"update_time" form:"update_time" xml:"update_time"`
	// IsDelete defined 删除标记
	IsDelete int64 `xorm:"notnull comment('删除标记') 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// Remark defined 备注
	Remark string `xorm:"varchar(200) comment('备注') 'remark'" json:"remark" form:"remark" xml:"remark"`
}

// TableName defined TODO
func (m *Entity) TableName() string {
	return "entity"
}

// Marshal defined TODO
func (r *Entity) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// UnmarshalEntity defined TODO
func UnmarshalEntity(data []byte) (Entity, error) {
	var r Entity
	err := json.Unmarshal(data, &r)
	return r, err
}
