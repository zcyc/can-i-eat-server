package food_to_food_tag_domain

import (
	"time"
)

type ListResp struct {
	Items   []*FoodToFoodTag `json:"items"`
	Total   int              `json:"total"`
	Current int              `json:"current"`
	Size    int              `json:"size"`
}

// FoodToFoodTag [...]
type FoodToFoodTag struct {
	Active     int8      `gorm:"column:active" json:"active"`
	Flag       int8      `gorm:"column:flag" json:"flag"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	ID         string    `gorm:"primaryKey;column:id" json:"-"` // 主键
	FoodID     string    `gorm:"column:food_id" json:"foodId"`
	FoodTagID  string    `gorm:"column:food_tag_id" json:"foodTagId"`
}
