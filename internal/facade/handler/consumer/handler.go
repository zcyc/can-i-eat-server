package consumer_facade

import (
	string_util "can-i-eat/common/util/string"
	food_domain "can-i-eat/internal/domain/food"
	food_service "can-i-eat/internal/service/food"
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

	resp, err := food_service.Impl.List(size, page)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func handlerDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.String(http.StatusOK, "参数错误")
		return
	}
	food, _ := food_service.Impl.Detail(id)
	c.JSON(http.StatusOK, food)
}

func handlerCreate(c *gin.Context) {
	food := new(food_domain.Food)
	if err := c.Bind(food); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	id, err := food_service.Impl.Create(food)
	if err != nil {
		c.String(http.StatusOK, "创建失败")
		return
	}
	c.JSON(http.StatusOK, id)
}

func handlerUpdate(c *gin.Context) {
	food := new(food_domain.Food)
	if err := c.Bind(food); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err := food_service.Impl.Update(food)
	if err != nil {
		c.String(http.StatusOK, "更新失败")
		return
	}
	c.JSON(http.StatusOK, "更新成功")
}

func handlerDelete(c *gin.Context) {
	food := new(food_domain.Food)
	if err := c.Bind(food); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err := food_service.Impl.Delete(food)
	if err != nil {
		c.String(http.StatusOK, "更新失败")
		return
	}
	c.JSON(http.StatusOK, "更新成功")
}
