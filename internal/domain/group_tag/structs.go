package group_tag_domain

import (
	"time"
)

type ListResp struct {
	Items   []*GroupTag `json:"items"`
	Total   int         `json:"total"`
	Current int         `json:"current"`
	Size    int         `json:"size"`
}

// GroupTag [...]
type GroupTag struct {
	Active     int8      `gorm:"column:active" json:"active"`
	Flag       int8      `gorm:"column:flag" json:"flag"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	ID         uint64    `gorm:"primaryKey;column:id" json:"-"` // 主键
	GroupID    int64     `gorm:"column:group_id" json:"groupId"`
	TagID      int64     `gorm:"column:tag_id" json:"tagId"`
	Type       string    `gorm:"column:type" json:"type"`
}
