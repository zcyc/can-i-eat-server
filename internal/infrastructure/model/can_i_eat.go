package model

import (
	"time"
)

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

// TableName get sql table name.获取数据库表名
func (m *Consumer) TableName() string {
	return "consumer"
}

// ConsumerColumns get sql column name.获取数据库列名
var ConsumerColumns = struct {
	Active     string
	Flag       string
	CreateTime string
	UpdateTime string
	ID         string
	Name       string
	Account    string
	Password   string
}{
	Active:     "active",
	Flag:       "flag",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	ID:         "id",
	Name:       "name",
	Account:    "account",
	Password:   "password",
}

// ConsumerGroup [...]
type ConsumerGroup struct {
	Active     int8      `gorm:"column:active" json:"active"`
	Flag       int8      `gorm:"column:flag" json:"flag"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	ID         string    `gorm:"primaryKey;column:id" json:"-"` // 主键
	ConsumerID string    `gorm:"column:consumer_id" json:"consumerId"`
	GroupID    string    `gorm:"column:group_id" json:"groupId"`
}

// TableName get sql table name.获取数据库表名
func (m *ConsumerGroup) TableName() string {
	return "consumer_group"
}

// ConsumerGroupColumns get sql column name.获取数据库列名
var ConsumerGroupColumns = struct {
	Active     string
	Flag       string
	CreateTime string
	UpdateTime string
	ID         string
	ConsumerID string
	GroupID    string
}{
	Active:     "active",
	Flag:       "flag",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	ID:         "id",
	ConsumerID: "consumer_id",
	GroupID:    "group_id",
}

// Food [...]
type Food struct {
	Active     int8      `gorm:"column:active" json:"active"`
	Flag       int8      `gorm:"column:flag" json:"flag"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	ID         string    `gorm:"primaryKey;column:id" json:"-"` // 主键
	Name       string    `gorm:"column:name" json:"name"`
	Alias      string    `gorm:"column:alias" json:"alias"`
}

// TableName get sql table name.获取数据库表名
func (m *Food) TableName() string {
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
}{
	Active:     "active",
	Flag:       "flag",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	ID:         "id",
	Name:       "name",
	Alias:      "alias",
}

// FoodTag [...]
type FoodTag struct {
	Active     int8      `gorm:"column:active" json:"active"`
	Flag       int8      `gorm:"column:flag" json:"flag"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	ID         string    `gorm:"primaryKey;column:id" json:"-"` // 主键
	FoodID     string    `gorm:"column:food_id" json:"foodId"`
	TagID      string    `gorm:"column:tag_id" json:"tagId"`
}

// TableName get sql table name.获取数据库表名
func (m *FoodTag) TableName() string {
	return "food_tag"
}

// FoodTagColumns get sql column name.获取数据库列名
var FoodTagColumns = struct {
	Active     string
	Flag       string
	CreateTime string
	UpdateTime string
	ID         string
	FoodID     string
	TagID      string
}{
	Active:     "active",
	Flag:       "flag",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	ID:         "id",
	FoodID:     "food_id",
	TagID:      "tag_id",
}

// Group [...]
type Group struct {
	Active     int8      `gorm:"column:active" json:"active"`
	Flag       int8      `gorm:"column:flag" json:"flag"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	ID         string    `gorm:"primaryKey;column:id" json:"-"` // 主键
	Name       string    `gorm:"column:name" json:"name"`
}

// TableName get sql table name.获取数据库表名
func (m *Group) TableName() string {
	return "group"
}

// GroupColumns get sql column name.获取数据库列名
var GroupColumns = struct {
	Active     string
	Flag       string
	CreateTime string
	UpdateTime string
	ID         string
	Name       string
}{
	Active:     "active",
	Flag:       "flag",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	ID:         "id",
	Name:       "name",
}

// GroupTag [...]
type GroupTag struct {
	Active     int8      `gorm:"column:active" json:"active"`
	Flag       int8      `gorm:"column:flag" json:"flag"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	ID         string    `gorm:"primaryKey;column:id" json:"-"` // 主键
	GroupID    string    `gorm:"column:group_id" json:"groupId"`
	TagID      string    `gorm:"column:tag_id" json:"tagId"`
	Type       string    `gorm:"column:type" json:"type"`
}

// TableName get sql table name.获取数据库表名
func (m *GroupTag) TableName() string {
	return "group_tag"
}

// GroupTagColumns get sql column name.获取数据库列名
var GroupTagColumns = struct {
	Active     string
	Flag       string
	CreateTime string
	UpdateTime string
	ID         string
	GroupID    string
	TagID      string
	Type       string
}{
	Active:     "active",
	Flag:       "flag",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	ID:         "id",
	GroupID:    "group_id",
	TagID:      "tag_id",
	Type:       "type",
}

// Tag [...]
type Tag struct {
	Active     int8      `gorm:"column:active" json:"active"`
	Flag       int8      `gorm:"column:flag" json:"flag"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	ID         string    `gorm:"primaryKey;column:id" json:"-"` // 主键
	Name       string    `gorm:"column:name" json:"name"`
	ParentID   string    `gorm:"column:parent_id" json:"parentId"`
}

// TableName get sql table name.获取数据库表名
func (m *Tag) TableName() string {
	return "tag"
}

// TagColumns get sql column name.获取数据库列名
var TagColumns = struct {
	Active     string
	Flag       string
	CreateTime string
	UpdateTime string
	ID         string
	Name       string
	ParentID   string
}{
	Active:     "active",
	Flag:       "flag",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	ID:         "id",
	Name:       "name",
	ParentID:   "parent_id",
}
