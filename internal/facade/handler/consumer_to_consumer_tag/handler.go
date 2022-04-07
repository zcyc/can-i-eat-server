package consumer_to_consumer_tag_facade

import (
	string_util "can-i-eat/common/util/string"
	consumer_to_consumer_tag_domain "can-i-eat/internal/domain/consumer_to_consumer_tag"
	consumer_to_consumer_tag_service "can-i-eat/internal/service/consumer_to_consumer_tag"
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

	resp, err := consumer_to_consumer_tag_service.Impl.List(size, page)
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
	consumerToConsumerTag, _ := consumer_to_consumer_tag_service.Impl.Detail(id)
	return c.JSON(http.StatusOK, consumerToConsumerTag)
}

func handlerCreate(c echo.Context) error {
	consumerToConsumerTag := new(consumer_to_consumer_tag_domain.ConsumerToConsumerTag)
	if err := c.Bind(consumerToConsumerTag); err != nil {
		return err
	}
	id, err := consumer_to_consumer_tag_service.Impl.Create(consumerToConsumerTag)
	if err != nil {
		return c.String(http.StatusOK, "创建失败")
	}
	return c.JSON(http.StatusOK, id)
}

func handlerUpdate(c echo.Context) error {
	consumerToConsumerTag := new(consumer_to_consumer_tag_domain.ConsumerToConsumerTag)
	if err := c.Bind(consumerToConsumerTag); err != nil {
		return err
	}
	err := consumer_to_consumer_tag_service.Impl.Update(consumerToConsumerTag)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}

func handlerDelete(c echo.Context) error {
	consumerToConsumerTag := new(consumer_to_consumer_tag_domain.ConsumerToConsumerTag)
	if err := c.Bind(consumerToConsumerTag); err != nil {
		return err
	}
	err := consumer_to_consumer_tag_service.Impl.Delete(consumerToConsumerTag)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}
