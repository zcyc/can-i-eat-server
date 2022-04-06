package group_tag_facade

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo) {
	e.GET(GetContextPathV1("/list"), handlerList)
	e.GET(GetContextPathV1("/detail"), handlerDetail)
	e.GET(GetContextPathV1("/list-by-group"), handlerListByGroup)
	e.POST(GetContextPathV1("/create"), handlerCreate)
	e.POST(GetContextPathV1("/update"), handlerUpdate)
	e.POST(GetContextPathV1("/delete"), handlerDelete)
}

func GetContextPathV1(action string) string {
	return fmt.Sprintf("/v1/group-tag%s", action)
}
