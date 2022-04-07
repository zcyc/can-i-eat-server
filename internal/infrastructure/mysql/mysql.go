package mysql_infrastructure

import (
	"can-i-eat/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlClient *gorm.DB

func Init() {
	// 解析配置文件
	configMap := config.Init("./config/mysql/config")
	user := configMap["user"]
	password := configMap["password"]
	host := configMap["host"]
	port := configMap["port"]
	database := configMap["database"]

	// 连接 mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	//_ = db.AutoMigrate(&model.Food{}, &model.Tag{}, &model.FoodToFoodTag{}, &model.Consumer{}, &model.ConsumerTag{}, &model.ConsumerToConsumerTag{})

	mysqlClient = db
}

// Get 获取数据库连接
func Get() *gorm.DB {
	return mysqlClient
}
