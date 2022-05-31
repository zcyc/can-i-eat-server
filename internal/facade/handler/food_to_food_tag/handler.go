package food_to_food_tag_facade

import (
	string_util "can-i-eat/common/util/string"
	food_to_food_tag_domain "can-i-eat/internal/domain/food_to_food_tag"
	food_to_food_tag_service "can-i-eat/internal/service/food_to_food_tag"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handlerList(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := string_util.StringToInt64(pageStr)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	sizeStr := c.Query("size")
	size, err := string_util.StringToInt64(sizeStr)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := food_to_food_tag_service.Impl.List(size, page)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func handlerDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.String(http.StatusBadRequest, "参数错误")
		return
	}
	foodToFoodTag, _ := food_to_food_tag_service.Impl.Detail(id)
	c.JSON(http.StatusOK, foodToFoodTag)
}

func handlerCreate(c *gin.Context) {
	foodToFoodTag := new(food_to_food_tag_domain.FoodToFoodTag)
	if err := c.Bind(foodToFoodTag); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	id, err := food_to_food_tag_service.Impl.Create(foodToFoodTag)
	if err != nil {
		c.String(http.StatusOK, "创建失败")
		return
	}
	c.JSON(http.StatusOK, id)
}

func handlerUpdate(c *gin.Context) {
	foodToFoodTag := new(food_to_food_tag_domain.FoodToFoodTag)
	if err := c.Bind(foodToFoodTag); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err := food_to_food_tag_service.Impl.Update(foodToFoodTag)
	if err != nil {
		c.String(http.StatusOK, "更新失败")
		return
	}
	c.JSON(http.StatusOK, "更新成功")
}

func handlerDelete(c *gin.Context) {
	foodToFoodTag := new(food_to_food_tag_domain.FoodToFoodTag)
	if err := c.Bind(foodToFoodTag); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err := food_to_food_tag_service.Impl.Delete(foodToFoodTag)
	if err != nil {
		c.String(http.StatusOK, "更新失败")
		return
	}
	c.JSON(http.StatusOK, "更新成功")
}
