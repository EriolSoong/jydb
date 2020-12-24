package field

import (
	"time"
)

type Meta struct {
	IsPrimaryKey   bool   // 是否为主键
	IsForeignKey   bool   // 是否为外键
	IsUnique       bool   // 是否为唯一约束
	IsDeprecated   bool   // 是否已废弃, 字段重命名操作后，旧字段名标记为废弃
	IsDeleted      bool   // 是否被删除
	Type           string // 类型
	DeprecatedName string // 废弃之前的名字
}

type Field struct {
	Meta       *Meta
	TbUUID     string
	Name       string
	UUID       string
	Value      string
	Remark     string
	ModifyTime time.Time
	History    map[time.Time]*Field
}
