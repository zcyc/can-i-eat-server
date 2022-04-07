package food_to_food_tag_facade

import (
	string_util "can-i-eat/common/util/string"
	food_to_food_tag_domain "can-i-eat/internal/domain/food_to_food_tag"
	food_to_food_tag_service "can-i-eat/internal/service/food_to_food_tag"
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

	resp, err := food_to_food_tag_service.Impl.List(size, page)
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
	foodToFoodTag, _ := food_to_food_tag_service.Impl.Detail(id)
	return c.JSON(http.StatusOK, foodToFoodTag)
}

func handlerCreate(c echo.Context) error {
	foodToFoodTag := new(food_to_food_tag_domain.FoodToFoodTag)
	if err := c.Bind(foodToFoodTag); err != nil {
		return err
	}
	id, err := food_to_food_tag_service.Impl.Create(foodToFoodTag)
	if err != nil {
		return c.String(http.StatusOK, "创建失败")
	}
	return c.JSON(http.StatusOK, id)
}

func handlerUpdate(c echo.Context) error {
	foodToFoodTag := new(food_to_food_tag_domain.FoodToFoodTag)
	if err := c.Bind(foodToFoodTag); err != nil {
		return err
	}
	err := food_to_food_tag_service.Impl.Update(foodToFoodTag)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}

func handlerDelete(c echo.Context) error {
	foodToFoodTag := new(food_to_food_tag_domain.FoodToFoodTag)
	if err := c.Bind(foodToFoodTag); err != nil {
		return err
	}
	err := food_to_food_tag_service.Impl.Delete(foodToFoodTag)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}
