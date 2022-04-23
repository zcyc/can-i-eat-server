package consumer_to_consumer_tag_domain

import (
	"time"
)

type ListResp struct {
	Items   []*ConsumerToConsumerTag `json:"items"`
	Total   int                      `json:"total"`
	Current int                      `json:"current"`
	Size    int                      `json:"size"`
}

// ConsumerToConsumerTag [...]
type ConsumerToConsumerTag struct {
	Active        int8      `gorm:"column:active" json:"active"`
	Flag          int8      `gorm:"column:flag" json:"flag"`
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"updateTime"`
	ID            string    `gorm:"primaryKey;column:id" json:"id"` // 主键
	ConsumerID    string    `gorm:"column:consumer_id" json:"consumerId"`
	ConsumerTagID string    `gorm:"column:consumer_tag_id" json:"consumerTagId"`
}
