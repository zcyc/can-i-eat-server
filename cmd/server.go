package main

import (
	consumer_facade "can-i-eat/internal/facade/handler/consumer"
	consumer_tag_facade "can-i-eat/internal/facade/handler/consumer_tag"
	consumer_tag_to_food_tag_facade "can-i-eat/internal/facade/handler/consumer_tag_to_food_tag"
	consumer_to_consumer_tag_facade "can-i-eat/internal/facade/handler/consumer_to_consumer_tag"
	data_source_facade "can-i-eat/internal/facade/handler/data_source"
	eat_mode_facade "can-i-eat/internal/facade/handler/eat_mode"
	food_facade "can-i-eat/internal/facade/handler/food"
	food_tag_facade "can-i-eat/internal/facade/handler/food_tag"
	food_to_food_tag_facade "can-i-eat/internal/facade/handler/food_to_food_tag"
	mysql_infrastructure "can-i-eat/internal/infrastructure/mysql"
	redis_infrastructure "can-i-eat/internal/infrastructure/redis"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"os"
)

func main() {
	// 初始化 echo
	r := gin.Default()
	// 防跨域中间件
	r.Use(cors.Default())

	// 初始化路由
	initFacade(r)

	// 初始化基础设施
	initInfrastructure()

	// 启动 http 监听
	err := r.Run()
	if err != nil {
		log.Info(err)
		os.Exit(1)
	}
}

// initFacade 初始化路由
func initFacade(r *gin.Engine) {
	// 注册路由
	food_facade.RegisterHandlers(r)
	food_tag_facade.RegisterHandlers(r)
	food_to_food_tag_facade.RegisterHandlers(r)
	consumer_facade.RegisterHandlers(r)
	consumer_tag_facade.RegisterHandlers(r)
	consumer_to_consumer_tag_facade.RegisterHandlers(r)
	consumer_tag_to_food_tag_facade.RegisterHandlers(r)
	data_source_facade.RegisterHandlers(r)
	eat_mode_facade.RegisterHandlers(r)
}

// initInfrastructure 初始化基础设施
func initInfrastructure() {
	mysql_infrastructure.Init()
	redis_infrastructure.Init()
}
