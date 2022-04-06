package group_facade

import (
	string_util "can-i-eat/common/util/string"
	group_domain "can-i-eat/internal/domain/group"
	group_service "can-i-eat/internal/service/group"
	"github.com/labstack/echo/v4"
	"net/http"
)

func handlerList(c echo.Context) error {
	sizeStr := c.QueryParam("page")
	pageStr := c.QueryParam("size")

	size, err := string_util.StringToInt64(sizeStr)
	if err != nil {
		return err
	}
	page, err := string_util.StringToInt64(pageStr)
	if err != nil {
		return err
	}

	resp, err := group_service.Impl.List(size, page)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func handlerDetail(c echo.Context) error {
	idStr := c.QueryParam("id")
	id, err := string_util.StringToInt64(idStr)
	if err != nil {
		return err
	}
	group, err := group_service.Impl.Detail(id)
	return c.JSON(http.StatusOK, group)
}

func handlerCreate(c echo.Context) error {
	group := new(group_domain.Group)
	if err := c.Bind(group); err != nil {
		return err
	}
	id, err := group_service.Impl.Create(group)
	if err != nil {
		return c.String(http.StatusOK, "创建失败")
	}
	return c.JSON(http.StatusOK, id)
}

func handlerUpdate(c echo.Context) error {
	group := new(group_domain.Group)
	if err := c.Bind(group); err != nil {
		return err
	}
	err := group_service.Impl.Update(group)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}

func handlerDelete(c echo.Context) error {
	group := new(group_domain.Group)
	if err := c.Bind(group); err != nil {
		return err
	}
	err := group_service.Impl.Delete(group)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}
