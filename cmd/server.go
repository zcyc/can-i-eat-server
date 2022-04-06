package main

import (
	consumer_facade "can-i-eat/internal/facade/handler/consumer"
	consumer_group_facade "can-i-eat/internal/facade/handler/consumer_group"
	food_facade "can-i-eat/internal/facade/handler/food"
	food_tag_facade "can-i-eat/internal/facade/handler/food_tag"
	group_facade "can-i-eat/internal/facade/handler/group"
	tag_facade "can-i-eat/internal/facade/handler/tag"
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
	tag_facade.RegisterHandlers(e)
	food_tag_facade.RegisterHandlers(e)
	consumer_facade.RegisterHandlers(e)
	group_facade.RegisterHandlers(e)
	consumer_group_facade.RegisterHandlers(e)
}

// initInfrastructure 初始化基础设施
func initInfrastructure() {
	mysql_infrastructure.Init()
	redis_infrastructure.Init()
}
