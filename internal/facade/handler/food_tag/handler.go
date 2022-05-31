package food_tag_facade

import (
	string_util "can-i-eat/common/util/string"
	food_tag_domain "can-i-eat/internal/domain/food_tag"
	food_tag_service "can-i-eat/internal/service/food_tag"
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

	resp, err := food_tag_service.Impl.List(size, page)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func handlerDetail(c *gin.Context) {
	idStr := c.Query("id")
	id, err := string_util.StringToInt64(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	tag, err := food_tag_service.Impl.Detail(id)
	c.JSON(http.StatusOK, tag)
}

func handlerCreate(c *gin.Context) {
	tag := new(food_tag_domain.FoodTag)
	if err := c.Bind(tag); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	id, err := food_tag_service.Impl.Create(tag)
	if err != nil {
		c.String(http.StatusOK, "创建失败")
		return
	}
	c.JSON(http.StatusOK, id)
}

func handlerUpdate(c *gin.Context) {
	tag := new(food_tag_domain.FoodTag)
	if err := c.Bind(tag); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err := food_tag_service.Impl.Update(tag)
	if err != nil {
		c.String(http.StatusOK, "更新失败")
		return
	}
	c.JSON(http.StatusOK, "更新成功")
}

func handlerDelete(c *gin.Context) {
	tag := new(food_tag_domain.FoodTag)
	if err := c.Bind(tag); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err := food_tag_service.Impl.Delete(tag)
	if err != nil {
		c.String(http.StatusOK, "更新失败")
		return
	}
	c.JSON(http.StatusOK, "更新成功")
}
