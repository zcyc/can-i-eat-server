package main

import (
	food_facade "can-i-eat/internal/facade/handler/food"
	mysql_infrastructure "can-i-eat/internal/infrastructure/mysql"
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

// 初始化 facade
func initFacade(e *echo.Echo) {
	// 所有 facade 都在这里注册
	food_facade.RegisterHandlers(e)
}

func initInfrastructure() {
	// 所有基础设施都在这里注册
	mysql_infrastructure.Init()
}
