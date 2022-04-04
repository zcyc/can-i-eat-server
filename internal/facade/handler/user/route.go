package user_facade

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo) {
	e.GET(GetContextPathV1("/detail"), handlerUserDetail)
}

func GetContextPathV1(action string) string {
	return fmt.Sprintf("/v1/user%s", action)
}
