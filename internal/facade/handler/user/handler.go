package user_facade

import (
	"can-i-eat/common/util"
	food_domain "can-i-eat/internal/domain/food"
	food_service "can-i-eat/internal/service/food"
	"github.com/labstack/echo/v4"
	"net/http"
)

func handlerUserList(c echo.Context) error {
	sizeStr := c.QueryParam("page")
	pageStr := c.QueryParam("size")

	size, err := util.StringToInt64(sizeStr)
	if err != nil {
		return err
	}
	page, err := util.StringToInt64(pageStr)
	if err != nil {
		return err
	}

	resp, err := food_service.Impl.ListForPage(size, page)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func handlerUserDetail(c echo.Context) error {
	idStr := c.QueryParam("id")
	id, err := util.StringToInt64(idStr)
	if err != nil {
		return err
	}
	food, err := food_service.Impl.FoodDetail(id)
	return c.JSON(http.StatusOK, food)
}

func handlerUserCreate(c echo.Context) error {
	food := new(food_domain.Food)
	if err := c.Bind(food); err != nil {
		return err
	}
	id, err := food_service.Impl.Create(food)
	if err != nil {
		return c.String(http.StatusOK, "创建失败")
	}
	return c.JSON(http.StatusOK, id)
}
