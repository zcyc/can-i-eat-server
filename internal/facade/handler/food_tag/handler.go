package food_tag_facade

import (
	string_util "can-i-eat/common/util/string"
	food_tag_domain "can-i-eat/internal/domain/food_tag"
	food_tag_service "can-i-eat/internal/service/food_tag"
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

	resp, err := food_tag_service.Impl.ListForPage(size, page)
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
	foodTag, err := food_tag_service.Impl.FoodDetail(id)
	return c.JSON(http.StatusOK, foodTag)
}

func handlerCreate(c echo.Context) error {
	foodTag := new(food_tag_domain.FoodTag)
	if err := c.Bind(foodTag); err != nil {
		return err
	}
	id, err := food_tag_service.Impl.Create(foodTag)
	if err != nil {
		return c.String(http.StatusOK, "创建失败")
	}
	return c.JSON(http.StatusOK, id)
}

func handlerUpdate(c echo.Context) error {
	foodTag := new(food_tag_domain.FoodTag)
	if err := c.Bind(foodTag); err != nil {
		return err
	}
	err := food_tag_service.Impl.Update(foodTag)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}

func handlerDelete(c echo.Context) error {
	foodTag := new(food_tag_domain.FoodTag)
	if err := c.Bind(foodTag); err != nil {
		return err
	}
	err := food_tag_service.Impl.Delete(foodTag)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}
