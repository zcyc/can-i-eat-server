package group_tag_facade

import (
	string_util "can-i-eat/common/util/string"
	group_tag_domain "can-i-eat/internal/domain/group_tag"
	group_tag_service "can-i-eat/internal/service/group_tag"
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

	resp, err := group_tag_service.Impl.List(size, page)
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
	groupTag, _ := group_tag_service.Impl.Detail(id)
	return c.JSON(http.StatusOK, groupTag)
}

func handlerCreate(c echo.Context) error {
	groupTag := new(group_tag_domain.GroupTag)
	if err := c.Bind(groupTag); err != nil {
		return err
	}
	id, err := group_tag_service.Impl.Create(groupTag)
	if err != nil {
		return c.String(http.StatusOK, "创建失败")
	}
	return c.JSON(http.StatusOK, id)
}

func handlerUpdate(c echo.Context) error {
	groupTag := new(group_tag_domain.GroupTag)
	if err := c.Bind(groupTag); err != nil {
		return err
	}
	err := group_tag_service.Impl.Update(groupTag)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}

func handlerDelete(c echo.Context) error {
	groupTag := new(group_tag_domain.GroupTag)
	if err := c.Bind(groupTag); err != nil {
		return err
	}
	err := group_tag_service.Impl.Delete(groupTag)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}

func handlerListByGroup(c echo.Context) error {
	groupId := c.QueryParam("group_id")
	if groupId == "" {
		return errors.New("参数错误")
	}
	resp, err := group_tag_service.Impl.ListByGroup(groupId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}
