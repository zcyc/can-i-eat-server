package consumer_group_domain

import (
	"time"
)

type ListResp struct {
	Items   []*ConsumerGroup `json:"items"`
	Total   int              `json:"total"`
	Current int              `json:"current"`
	Size    int              `json:"size"`
}

// ConsumerGroup [...]
type ConsumerGroup struct {
	Active     int8      `gorm:"column:active;default:1" json:"active"`
	Flag       int8      `gorm:"column:flag;default:0" json:"flag"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	ID         string    `gorm:"primaryKey;column:id" json:"id"` // 主键
	ConsumerID string    `gorm:"column:consumer_id" json:"consumerId"`
	GroupID    string    `gorm:"column:group_id" json:"groupId"`
}
