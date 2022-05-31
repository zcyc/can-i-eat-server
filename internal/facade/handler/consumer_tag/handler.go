package consumer_tag_facade

import (
	string_util "can-i-eat/common/util/string"
	consumer_tag_domain "can-i-eat/internal/domain/consumer_tag"
	consumer_tag_service "can-i-eat/internal/service/consumer_tag"
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

	resp, err := consumer_tag_service.Impl.List(size, page)
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
	consumerTag, err := consumer_tag_service.Impl.Detail(id)
	c.JSON(http.StatusOK, consumerTag)
}

func handlerCreate(c *gin.Context) {
	consumerTag := new(consumer_tag_domain.ConsumerTag)
	if err := c.Bind(consumerTag); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	id, err := consumer_tag_service.Impl.Create(consumerTag)
	if err != nil {
		c.String(http.StatusOK, "创建失败")
		return
	}
	c.JSON(http.StatusOK, id)
}

func handlerUpdate(c *gin.Context) {
	consumerTag := new(consumer_tag_domain.ConsumerTag)
	if err := c.Bind(consumerTag); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err := consumer_tag_service.Impl.Update(consumerTag)
	if err != nil {
		c.String(http.StatusOK, "更新失败")
		return
	}
	c.JSON(http.StatusOK, "更新成功")
}

func handlerDelete(c *gin.Context) {
	consumerTag := new(consumer_tag_domain.ConsumerTag)
	if err := c.Bind(consumerTag); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err := consumer_tag_service.Impl.Delete(consumerTag)
	if err != nil {
		c.String(http.StatusOK, "更新失败")
		return
	}
	c.JSON(http.StatusOK, "更新成功")
}
