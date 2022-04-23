package food_facade

import (
	string_util "can-i-eat/common/util/string"
	consumer_tag_to_food_tag_application "can-i-eat/internal/application/group_food"
	food_domain "can-i-eat/internal/domain/food"
	food_service "can-i-eat/internal/service/food"
	food_to_food_tag_service "can-i-eat/internal/service/food_to_food_tag"
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

	resp, err := food_service.Impl.List(size, page)
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
	food, _ := food_service.Impl.Detail(id)
	return c.JSON(http.StatusOK, food)
}

func handlerCreate(c echo.Context) error {
	food := new(food_domain.Food)
	if err := c.Bind(food); err != nil {
		return err
	}
	id, err := food_service.Impl.Create(food)
	if err != nil {
		return c.String(http.StatusOK, "创建失败")
	}
	return c.JSON(http.StatusOK, id)
}

func handlerUpdate(c echo.Context) error {
	food := new(food_domain.Food)
	if err := c.Bind(food); err != nil {
		return err
	}
	err := food_service.Impl.Update(food)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}

func handlerDelete(c echo.Context) error {
	food := new(food_domain.Food)
	if err := c.Bind(food); err != nil {
		return err
	}
	err := food_service.Impl.Delete(food)
	if err != nil {
		return c.String(http.StatusOK, "更新失败")
	}
	return c.JSON(http.StatusOK, "更新成功")
}

func handlerListByConsumerTag(c echo.Context) error {
	consumerTagID := c.QueryParam("consumer-tag-id")
	list, err := consumer_tag_to_food_tag_application.Impl.ListFoodByConsumerTag(consumerTagID)
	if err != nil {
		return nil
	}
	return c.JSON(http.StatusOK, list)
}

func handlerListByFoodTagList(c echo.Context) error {
	foodTagIdList := new(food_domain.ListByFoodTagListReq)
	if err := c.Bind(foodTagIdList); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.Info()

	foodToFoodTagLit, err := food_to_food_tag_service.Impl.ListByTagList(foodTagIdList.FoodTagIdList)
	if err != nil {
		return nil
	}

	var foodIDList []string
	for i := range foodToFoodTagLit {
		if !isInList(foodIDList, foodToFoodTagLit[i].FoodID) {
			foodIDList = append(foodIDList, foodToFoodTagLit[i].FoodID)
		}
	}

	foodList, err := food_service.Impl.ListByIDs(foodIDList)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, foodList)
}

func handlerListByConsumer(c echo.Context) error {
	consumerID := c.QueryParam("consumer-id")
	list, err := consumer_tag_to_food_tag_application.Impl.ListFoodByConsumer(consumerID)
	if err != nil {
		return nil
	}
	return c.JSON(http.StatusOK, list)
}

// 判断 string 是不是在 []string 中
func isInList(list []string, str string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
