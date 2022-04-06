package consumer_group_facade

import (
	string_util "can-i-eat/common/util/string"
	consumer_group_domain "can-i-eat/internal/domain/consumer_group"
	consumer_group_service "can-i-eat/internal/service/consumer_group"
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

	resp, err := consumer_group_service.Impl.List(size, page)
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
	consumerGroup, _ := consumer_group_service.Impl.Detail(id)
	return c.JSON(http.StatusOK, consumerGroup)
}

func handlerCreate(c echo.Context) error {
	consumerGroup := new(consumer_group_domain.ConsumerGroup)
	if err := c.Bind(consumerGroup); err != nil {
		return err
	}
	id, err := consumer_group_service.Impl.Create(consumerGroup)
	if err != nil {
		return c.String(http.StatusOK, "创建失败")
	}
	return c.JSON(http.StatusOK, id)
}

func handlerUpdate(c echo.Context) error {
	consumerGroup := new(consumer_group_domain.ConsumerGroup)
	if err := c.Bind(consumerGroup); err != nil {
		return err
	}
	err := consumer_group_service.Impl.Update(consumerGroup)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}

func handlerDelete(c echo.Context) error {
	consumerGroup := new(consumer_group_domain.ConsumerGroup)
	if err := c.Bind(consumerGroup); err != nil {
		return err
	}
	err := consumer_group_service.Impl.Delete(consumerGroup)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}
