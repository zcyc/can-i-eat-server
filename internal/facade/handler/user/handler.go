package user_facade

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func handlerUserDetail(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
