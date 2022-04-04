package food_tag_domain

import (
	"time"
)

type ListResp struct {
	Items   []*FoodTag `json:"items"`
	Total   int        `json:"total"`
	Current int        `json:"current"`
	Size    int        `json:"size"`
}

// FoodTag [...]
type FoodTag struct {
	Active     int8      `gorm:"column:active" json:"active"`          // 表示数据是否处于可用状态， active = 1 可用，active=0不可用，操作可逆转
	Flag       int8      `gorm:"column:flag" json:"flag"`              // 是否删除 1 删除 0 未删除
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 记录写入时间
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 记录更新时间
	ID         uint64    `gorm:"primaryKey;column:id" json:"-"`        // 主键
	FoodID     int64     `gorm:"column:food_id" json:"foodId"`         // 名称
	TagID      int64     `gorm:"column:tag_id" json:"tagId"`
	Type       string    `gorm:"column:type" json:"type"`
}
