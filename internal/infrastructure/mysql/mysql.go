package mysql_infrastructure

import (
	"can-i-eat/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlDB *gorm.DB

func Init() {
	configMap := config.Init("./config/mysql/config")
	host := configMap["host"]
	port := configMap["port"]
	fmt.Println("host=", host, " port=", port)
	// 创建数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", configMap["user"], configMap["password"], configMap["host"], configMap["port"], configMap["dbname"])
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 赋值
	mysqlDB = db

	// 迁移 schema
	//db.AutoMigrate(&userDomain.User{})
}

// 获取数据库连接
func Get() *gorm.DB {
	return mysqlDB
}
