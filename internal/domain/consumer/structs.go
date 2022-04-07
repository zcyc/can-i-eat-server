package consumer_domain

import (
	"time"
)

type ListResp struct {
	Items   []*Consumer `json:"items"`
	Total   int         `json:"total"`
	Current int         `json:"current"`
	Size    int         `json:"size"`
}

// Consumer [...]
type Consumer struct {
	Active     int8      `gorm:"column:active" json:"active"`
	Flag       int8      `gorm:"column:flag" json:"flag"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	ID         string    `gorm:"primaryKey;column:id" json:"-"` // 主键
	Name       string    `gorm:"column:name" json:"name"`
	Account    string    `gorm:"column:account" json:"account"`
	Password   string    `gorm:"column:password" json:"password"`
}
