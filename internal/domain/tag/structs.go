package tag_domain

import (
	"time"
)

type ListResp struct {
	Items   []*Tag `json:"items"`
	Total   int    `json:"total"`
	Current int    `json:"current"`
	Size    int    `json:"size"`
}

// Tag [...]
type Tag struct {
	Active     int8      `gorm:"column:active" json:"active"`
	Flag       int8      `gorm:"column:flag" json:"flag"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	ID         string    `gorm:"primaryKey;column:id" json:"id"` // 主键
	Name       string    `gorm:"column:name" json:"name"`
	ParentID   string    `gorm:"column:parent_id" json:"parentId"`
}
