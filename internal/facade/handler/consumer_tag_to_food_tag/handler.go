package consumer_tag_to_food_tag_facade

import (
	string_util "can-i-eat/common/util/string"
	consumer_tag_to_food_tag_domain "can-i-eat/internal/domain/consumer_tag_to_food_tag"
	consumer_tag_to_food_tag_service "can-i-eat/internal/service/consumer_tag_to_food_tag"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
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

	resp, err := consumer_tag_to_food_tag_service.Impl.List(size, page)
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
	consumerTagToFoodTag, _ := consumer_tag_to_food_tag_service.Impl.Detail(id)
	c.JSON(http.StatusOK, consumerTagToFoodTag)
}

func handlerCreate(c *gin.Context) {
	consumerTagToFoodTag := new(consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag)
	if err := c.Bind(consumerTagToFoodTag); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	id, err := consumer_tag_to_food_tag_service.Impl.Create(consumerTagToFoodTag)
	if err != nil {
		c.String(http.StatusOK, "创建失败")
		return
	}
	c.JSON(http.StatusOK, id)
}

func handlerUpdate(c *gin.Context) {
	consumerTagToFoodTag := new(consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag)
	if err := c.Bind(consumerTagToFoodTag); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err := consumer_tag_to_food_tag_service.Impl.Update(consumerTagToFoodTag)
	if err != nil {
		c.String(http.StatusOK, "更新失败")
		return
	}
	c.JSON(http.StatusOK, "更新成功")
}

func handlerDelete(c *gin.Context) {
	consumerTagToFoodTag := new(consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag)
	if err := c.Bind(consumerTagToFoodTag); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err := consumer_tag_to_food_tag_service.Impl.Delete(consumerTagToFoodTag)
	if err != nil {
		c.String(http.StatusOK, "更新失败")
		return
	}
	c.JSON(http.StatusOK, "更新成功")
}

func handlerListByConsumerTagAndEatMode(c *gin.Context) {
	consumerTagId := c.Query("consumerTagId")
	currentEatModeId := c.Query("currentEatModeId")
	log.Infof("handlerListByConsumerTagAndEatMode consumerTagId: %s, currentEatModeId: %s", consumerTagId, currentEatModeId)
	if consumerTagId == "" || currentEatModeId == "" {
		c.String(http.StatusBadRequest, "参数错误")
		return
	}

	resp, err := consumer_tag_to_food_tag_service.Impl.ListByConsumerTagAndEatMode(consumerTagId, currentEatModeId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}
