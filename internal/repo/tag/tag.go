package tag_repo

import "time"

// TagDao [...]
type TagDao struct {
	Active     int8      `gorm:"column:active" json:"active"`          // 表示数据是否处于可用状态， active = 1 可用，active=0不可用，操作可逆转
	Flag       int8      `gorm:"column:flag" json:"flag"`              // 是否删除 1 删除 0 未删除
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 记录写入时间
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 记录更新时间
	ID         uint64    `gorm:"primaryKey;column:id" json:"-"`        // 主键
	Name       string    `gorm:"column:name" json:"name"`              // 名称
	Alias      string    `gorm:"column:alias" json:"alias"`            // 别名
	Category   string    `gorm:"column:category" json:"category"`      // 分类
}

// TableName get sql table name.获取数据库表名
func (m *TagDao) TableName() string {
	return "tag"
}

// TagColumns get sql column name.获取数据库列名
var TagColumns = struct {
	Active     string
	Flag       string
	CreateTime string
	UpdateTime string
	ID         string
	Name       string
	Alias      string
	Category   string
}{
	Active:     "active",
	Flag:       "flag",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	ID:         "id",
	Name:       "name",
	Alias:      "alias",
	Category:   "category",
}
