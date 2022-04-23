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
	ID         string    `gorm:"primaryKey;column:id" json:"id"` // 主键
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

// ConsumerTag [...]
type ConsumerTag struct {
	Active     int8      `gorm:"column:active" json:"active"`
	Flag       int8      `gorm:"column:flag" json:"flag"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	ID         string    `gorm:"primaryKey;column:id" json:"id"` // 主键
	Name       string    `gorm:"column:name" json:"name"`
}

// TableName get sql table name.获取数据库表名
func (m *ConsumerTag) TableName() string {
	return "consumer_tag"
}

// ConsumerTagColumns get sql column name.获取数据库列名
var ConsumerTagColumns = struct {
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

// TableName get sql table name.获取数据库表名
func (m *ConsumerTagToFoodTag) TableName() string {
	return "consumer_tag_to_food_tag"
}

// ConsumerTagToFoodTagColumns get sql column name.获取数据库列名
var ConsumerTagToFoodTagColumns = struct {
	Active        string
	Flag          string
	CreateTime    string
	UpdateTime    string
	ID            string
	ConsumerTagID string
	FoodTagID     string
	EatMode       string
}{
	Active:        "active",
	Flag:          "flag",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
	ID:            "id",
	ConsumerTagID: "consumer_tag_id",
	FoodTagID:     "food_tag_id",
	EatMode:       "eat_mode",
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

// TableName get sql table name.获取数据库表名
func (m *ConsumerToConsumerTag) TableName() string {
	return "consumer_to_consumer_tag"
}

// ConsumerToConsumerTagColumns get sql column name.获取数据库列名
var ConsumerToConsumerTagColumns = struct {
	Active        string
	Flag          string
	CreateTime    string
	UpdateTime    string
	ID            string
	ConsumerID    string
	ConsumerTagID string
}{
	Active:        "active",
	Flag:          "flag",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
	ID:            "id",
	ConsumerID:    "consumer_id",
	ConsumerTagID: "consumer_tag_id",
}

// Food [...]
type Food struct {
	Active     int8      `gorm:"column:active" json:"active"`
	Flag       int8      `gorm:"column:flag" json:"flag"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	ID         string    `gorm:"primaryKey;column:id" json:"id"` // 主键
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
	ID         string    `gorm:"primaryKey;column:id" json:"id"` // 主键
	Name       string    `gorm:"column:name" json:"name"`
	ParentID   string    `gorm:"column:parent_id" json:"parentId"`
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

// FoodToFoodTag [...]
type FoodToFoodTag struct {
	Active     int8      `gorm:"column:active" json:"active"`
	Flag       int8      `gorm:"column:flag" json:"flag"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	ID         string    `gorm:"primaryKey;column:id" json:"id"` // 主键
	FoodID     string    `gorm:"column:food_id" json:"foodId"`
	FoodTagID  string    `gorm:"column:food_tag_id" json:"foodTagId"`
}

// TableName get sql table name.获取数据库表名
func (m *FoodToFoodTag) TableName() string {
	return "food_to_food_tag"
}

// FoodToFoodTagColumns get sql column name.获取数据库列名
var FoodToFoodTagColumns = struct {
	Active     string
	Flag       string
	CreateTime string
	UpdateTime string
	ID         string
	FoodID     string
	FoodTagID  string
}{
	Active:     "active",
	Flag:       "flag",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	ID:         "id",
	FoodID:     "food_id",
	FoodTagID:  "food_tag_id",
}
