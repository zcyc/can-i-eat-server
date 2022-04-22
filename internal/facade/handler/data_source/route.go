package data_source_facade

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo) {
	e.POST(GetContextPathV1("/upload/bh"), handlerUploadBh)
}

func GetContextPathV1(action string) string {
	return fmt.Sprintf("/v1/data-source%s", action)
}
