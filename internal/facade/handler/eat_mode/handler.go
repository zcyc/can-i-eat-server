package eat_mode_facade

import (
	string_util "can-i-eat/common/util/string"
	eat_mode_domain "can-i-eat/internal/domain/eat_mode"
	eat_mode_service "can-i-eat/internal/service/eat_mode"
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

	resp, err := eat_mode_service.Impl.List(size, page)
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
	eatMode, _ := eat_mode_service.Impl.Detail(id)
	c.JSON(http.StatusOK, eatMode)
}

func handlerCreate(c *gin.Context) {
	eatMode := new(eat_mode_domain.EatMode)
	if err := c.Bind(eatMode); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	id, err := eat_mode_service.Impl.Create(eatMode)
	if err != nil {
		c.String(http.StatusOK, "创建失败")
		return
	}
	c.JSON(http.StatusOK, id)
}

func handlerUpdate(c *gin.Context) {
	eatMode := new(eat_mode_domain.EatMode)
	if err := c.Bind(eatMode); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err := eat_mode_service.Impl.Update(eatMode)
	if err != nil {
		c.String(http.StatusOK, "更新失败")
		return
	}
	c.JSON(http.StatusOK, "更新成功")
}

func handlerDelete(c *gin.Context) {
	eatMode := new(eat_mode_domain.EatMode)
	if err := c.Bind(eatMode); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err := eat_mode_service.Impl.Delete(eatMode)
	if err != nil {
		c.String(http.StatusOK, "更新失败")
		return
	}
	c.JSON(http.StatusOK, "更新成功")
}
