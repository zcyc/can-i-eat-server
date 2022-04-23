package eat_mode_facade

import (
	string_util "can-i-eat/common/util/string"
	eat_mode_domain "can-i-eat/internal/domain/eat_mode"
	eat_mode_service "can-i-eat/internal/service/eat_mode"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func handlerList(c echo.Context) error {
	pageStr := c.QueryParam("page")
	page, err := string_util.StringToInt64(pageStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	sizeStr := c.QueryParam("size")
	size, err := string_util.StringToInt64(sizeStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	resp, err := eat_mode_service.Impl.List(size, page)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func handlerDetail(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return errors.New("参数错误")
	}
	eatMode, _ := eat_mode_service.Impl.Detail(id)
	return c.JSON(http.StatusOK, eatMode)
}

func handlerCreate(c echo.Context) error {
	eatMode := new(eat_mode_domain.EatMode)
	if err := c.Bind(eatMode); err != nil {
		return err
	}
	id, err := eat_mode_service.Impl.Create(eatMode)
	if err != nil {
		return c.String(http.StatusOK, "创建失败")
	}
	return c.JSON(http.StatusOK, id)
}

func handlerUpdate(c echo.Context) error {
	eatMode := new(eat_mode_domain.EatMode)
	if err := c.Bind(eatMode); err != nil {
		return err
	}
	err := eat_mode_service.Impl.Update(eatMode)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}

func handlerDelete(c echo.Context) error {
	eatMode := new(eat_mode_domain.EatMode)
	if err := c.Bind(eatMode); err != nil {
		return err
	}
	err := eat_mode_service.Impl.Delete(eatMode)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}
