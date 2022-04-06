package category_facade

import (
	string_util "can-i-eat/common/util/string"
	category_domain "can-i-eat/internal/domain/category"
	category_service "can-i-eat/internal/service/category"
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

	resp, err := category_service.Impl.List(size, page)
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
	category, err := category_service.Impl.Detail(id)
	return c.JSON(http.StatusOK, category)
}

func handlerCreate(c echo.Context) error {
	category := new(category_domain.Category)
	if err := c.Bind(category); err != nil {
		return err
	}
	id, err := category_service.Impl.Create(category)
	if err != nil {
		return c.String(http.StatusOK, "创建失败")
	}
	return c.JSON(http.StatusOK, id)
}

func handlerUpdate(c echo.Context) error {
	category := new(category_domain.Category)
	if err := c.Bind(category); err != nil {
		return err
	}
	err := category_service.Impl.Update(category)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}

func handlerDelete(c echo.Context) error {
	category := new(category_domain.Category)
	if err := c.Bind(category); err != nil {
		return err
	}
	err := category_service.Impl.Delete(category)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}
