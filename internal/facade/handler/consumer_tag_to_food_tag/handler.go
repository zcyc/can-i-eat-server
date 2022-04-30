package consumer_tag_to_food_tag_facade

import (
	string_util "can-i-eat/common/util/string"
	consumer_tag_to_food_tag_domain "can-i-eat/internal/domain/consumer_tag_to_food_tag"
	consumer_tag_to_food_tag_service "can-i-eat/internal/service/consumer_tag_to_food_tag"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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

	resp, err := consumer_tag_to_food_tag_service.Impl.List(size, page)
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
	consumerTagToFoodTag, _ := consumer_tag_to_food_tag_service.Impl.Detail(id)
	return c.JSON(http.StatusOK, consumerTagToFoodTag)
}

func handlerCreate(c echo.Context) error {
	consumerTagToFoodTag := new(consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag)
	if err := c.Bind(consumerTagToFoodTag); err != nil {
		return err
	}
	id, err := consumer_tag_to_food_tag_service.Impl.Create(consumerTagToFoodTag)
	if err != nil {
		return c.String(http.StatusOK, "创建失败")
	}
	return c.JSON(http.StatusOK, id)
}

func handlerUpdate(c echo.Context) error {
	consumerTagToFoodTag := new(consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag)
	if err := c.Bind(consumerTagToFoodTag); err != nil {
		return err
	}
	err := consumer_tag_to_food_tag_service.Impl.Update(consumerTagToFoodTag)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}

func handlerDelete(c echo.Context) error {
	consumerTagToFoodTag := new(consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag)
	if err := c.Bind(consumerTagToFoodTag); err != nil {
		return err
	}
	err := consumer_tag_to_food_tag_service.Impl.Delete(consumerTagToFoodTag)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}

func handlerListByConsumerTagAndEatMode(c echo.Context) error {
	consumerTagId := c.QueryParam("consumerTagId")
	currentEatModeId := c.QueryParam("currentEatModeId")
	log.Infof("handlerListByConsumerTagAndEatMode consumerTagId: %s, currentEatModeId: %s", consumerTagId, currentEatModeId)
	if consumerTagId == "" || currentEatModeId == "" {
		return errors.New("参数错误")
	}

	resp, err := consumer_tag_to_food_tag_service.Impl.ListByConsumerTagAndEatMode(consumerTagId, currentEatModeId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}
