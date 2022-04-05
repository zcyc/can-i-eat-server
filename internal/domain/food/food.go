package food_domain

import (
	"time"
)

type ListResp struct {
	Items   []*Food `json:"items"`
	Total   int     `json:"total"`
	Current int     `json:"current"`
	Size    int     `json:"size"`
}

// Food [...]
type Food struct {
	Active     int8      `gorm:"column:active" json:"active"`
	Flag       int8      `gorm:"column:flag" json:"flag"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	ID         uint64    `gorm:"primaryKey;column:id" json:"-"` // 主键
	CategoryID uint64    `gorm:"column:category_id" json:"categoryId"`
	Name       string    `gorm:"column:name" json:"name"`
	Alias      string    `gorm:"column:alias" json:"alias"`
}
