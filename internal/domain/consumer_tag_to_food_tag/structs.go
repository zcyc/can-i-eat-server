package consumer_tag_to_food_tag_domain

import (
	"time"
)

type ListResp struct {
	Items   []*ConsumerTagToFoodTag `json:"items"`
	Total   int                     `json:"total"`
	Current int                     `json:"current"`
	Size    int                     `json:"size"`
}

// ConsumerTagToFoodTag [...]
type ConsumerTagToFoodTag struct {
	Active        int8      `gorm:"column:active" json:"active"`
	Flag          int8      `gorm:"column:flag" json:"flag"`
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"updateTime"`
	ID            string    `gorm:"primaryKey;column:id" json:"id"` // 主键
	ConsumerTagID string    `gorm:"column:consumer_tag_id" json:"consumerTagId"`
	FoodTagID     string    `gorm:"column:food_tag_id" json:"foodTagId"`
	EatMode       string    `gorm:"column:eat_mode" json:"eatMode"`
}
