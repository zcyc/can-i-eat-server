package consumer_to_consumer_tag_facade

import (
	string_util "can-i-eat/common/util/string"
	consumer_to_consumer_tag_domain "can-i-eat/internal/domain/consumer_to_consumer_tag"
	consumer_to_consumer_tag_service "can-i-eat/internal/service/consumer_to_consumer_tag"
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

	resp, err := consumer_to_consumer_tag_service.Impl.List(size, page)
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
	consumerToConsumerTag, _ := consumer_to_consumer_tag_service.Impl.Detail(id)
	c.JSON(http.StatusOK, consumerToConsumerTag)
}

func handlerCreate(c *gin.Context) {
	consumerToConsumerTag := new(consumer_to_consumer_tag_domain.ConsumerToConsumerTag)
	if err := c.Bind(consumerToConsumerTag); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	id, err := consumer_to_consumer_tag_service.Impl.Create(consumerToConsumerTag)
	if err != nil {
		c.String(http.StatusOK, "创建失败")
		return
	}
	c.JSON(http.StatusOK, id)
}

func handlerUpdate(c *gin.Context) {
	consumerToConsumerTag := new(consumer_to_consumer_tag_domain.ConsumerToConsumerTag)
	if err := c.Bind(consumerToConsumerTag); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err := consumer_to_consumer_tag_service.Impl.Update(consumerToConsumerTag)
	if err != nil {
		c.String(http.StatusOK, "更新失败")
		return
	}
	c.JSON(http.StatusOK, "更新成功")
}

func handlerDelete(c *gin.Context) {
	consumerToConsumerTag := new(consumer_to_consumer_tag_domain.ConsumerToConsumerTag)
	if err := c.Bind(consumerToConsumerTag); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err := consumer_to_consumer_tag_service.Impl.Delete(consumerToConsumerTag)
	if err != nil {
		c.String(http.StatusOK, "更新失败")
		return
	}
	c.JSON(http.StatusOK, "更新成功")
}
