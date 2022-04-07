package consumer_tag_facade

import (
	string_util "can-i-eat/common/util/string"
	consumer_tag_domain "can-i-eat/internal/domain/consumer_tag"
	consumer_tag_service "can-i-eat/internal/service/consumer_tag"
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

	resp, err := consumer_tag_service.Impl.List(size, page)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func handlerDetail(c echo.Context) error {
	idStr := c.QueryParam("id")
	id, err := string_util.StringToInt64(idStr)
	if err != nil {
		return err
	}
	consumerTag, err := consumer_tag_service.Impl.Detail(id)
	return c.JSON(http.StatusOK, consumerTag)
}

func handlerCreate(c echo.Context) error {
	consumerTag := new(consumer_tag_domain.ConsumerTag)
	if err := c.Bind(consumerTag); err != nil {
		return err
	}
	id, err := consumer_tag_service.Impl.Create(consumerTag)
	if err != nil {
		return c.String(http.StatusOK, "创建失败")
	}
	return c.JSON(http.StatusOK, id)
}

func handlerUpdate(c echo.Context) error {
	consumerTag := new(consumer_tag_domain.ConsumerTag)
	if err := c.Bind(consumerTag); err != nil {
		return err
	}
	err := consumer_tag_service.Impl.Update(consumerTag)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}

func handlerDelete(c echo.Context) error {
	consumerTag := new(consumer_tag_domain.ConsumerTag)
	if err := c.Bind(consumerTag); err != nil {
		return err
	}
	err := consumer_tag_service.Impl.Delete(consumerTag)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}
