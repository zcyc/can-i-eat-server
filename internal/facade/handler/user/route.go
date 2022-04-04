package user_facade

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo) {
	e.GET(GetContextPathV1("/list"), handlerUserList)
	e.GET(GetContextPathV1("/detail"), handlerUserDetail)
	e.POST(GetContextPathV1("/create"), handlerUserCreate)
}

func GetContextPathV1(action string) string {
	return fmt.Sprintf("/v1/food%s", action)
}
