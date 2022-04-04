package consumer_group_repo

import "time"

// ConsumerGroupDao [...]
type ConsumerGroupDao struct {
	Active     int8      `gorm:"column:active" json:"active"`          // 表示数据是否处于可用状态， active = 1 可用，active=0不可用，操作可逆转
	Flag       int8      `gorm:"column:flag" json:"flag"`              // 是否删除 1 删除 0 未删除
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 记录写入时间
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 记录更新时间
	ID         uint64    `gorm:"primaryKey;column:id" json:"-"`        // 主键
	ConsumerID int64     `gorm:"column:consumer_id" json:"consumerId"` // 名称
	GroupID    int64     `gorm:"column:group_id" json:"groupId"`       // 别名
}

// TableName get sql table name.获取数据库表名
func (m *ConsumerGroupDao) TableName() string {
	return "consumer_group"
}

// ConsumerGroupColumns get sql column name.获取数据库列名
var ConsumerGroupColumns = struct {
	Active     string
	Flag       string
	CreateTime string
	UpdateTime string
	ID         string
	ConsumerID string
	GroupID    string
}{
	Active:     "active",
	Flag:       "flag",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	ID:         "id",
	ConsumerID: "consumer_id",
	GroupID:    "group_id",
}
