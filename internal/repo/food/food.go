package food_repo

import (
	"time"
)

// FoodDao [...]
type FoodDao struct {
	Active     int8      `gorm:"column:active" json:"active"`
	Flag       int8      `gorm:"column:flag" json:"flag"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	ID         uint64    `gorm:"primaryKey;column:id" json:"-"` // 主键
	CategoryID uint64    `gorm:"column:category_id" json:"categoryId"`
	Name       string    `gorm:"column:name" json:"name"`
	Alias      string    `gorm:"column:alias" json:"alias"`
}

// TableName get sql table name.获取数据库表名
func (m *FoodDao) TableName() string {
	return "food"
}

// FoodColumns get sql column name.获取数据库列名
var FoodColumns = struct {
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
