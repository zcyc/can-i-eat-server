package main

import (
	consumer_facade "can-i-eat/internal/facade/handler/consumer"
	consumer_tag_facade "can-i-eat/internal/facade/handler/consumer_tag"
	consumer_tag_to_food_tag_facade "can-i-eat/internal/facade/handler/consumer_tag_to_food_tag"
	consumer_to_consumer_tag_facade "can-i-eat/internal/facade/handler/consumer_to_consumer_tag"
	food_facade "can-i-eat/internal/facade/handler/food"
	food_tag_facade "can-i-eat/internal/facade/handler/food_tag"
	food_to_food_tag_facade "can-i-eat/internal/facade/handler/food_to_food_tag"
	mysql_infrastructure "can-i-eat/internal/infrastructure/mysql"
	redis_infrastructure "can-i-eat/internal/infrastructure/redis"
	"github.com/labstack/echo/v4"
)

func main() {
	// 初始化 echo
	e := echo.New()

	// 初始化路由
	initFacade(e)

	// 初始化基础设施
	initInfrastructure()

	// 启动 http 监听
	e.Logger.Fatal(e.Start(":1323"))
}

// initFacade 初始化路由
func initFacade(e *echo.Echo) {
	food_facade.RegisterHandlers(e)
	food_tag_facade.RegisterHandlers(e)
	food_to_food_tag_facade.RegisterHandlers(e)
	consumer_facade.RegisterHandlers(e)
	consumer_tag_facade.RegisterHandlers(e)
	consumer_to_consumer_tag_facade.RegisterHandlers(e)
	consumer_tag_to_food_tag_facade.RegisterHandlers(e)
}

// initInfrastructure 初始化基础设施
func initInfrastructure() {
	mysql_infrastructure.Init()
	redis_infrastructure.Init()
}
