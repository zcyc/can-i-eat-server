package food_facade

import (
	"can-i-eat/common/constant"
	string_util "can-i-eat/common/util/string"
	food_domain "can-i-eat/internal/domain/food"
	consumer_tag_to_food_tag_service "can-i-eat/internal/service/consumer_tag_to_food_tag"
	food_service "can-i-eat/internal/service/food"
	food_to_food_tag_service "can-i-eat/internal/service/food_to_food_tag"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/utils"
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

func handlerListByFoodTagList(c echo.Context) error {
	foodTagIdList := new(food_domain.ListByFoodTagListReq)
	if err := c.Bind(foodTagIdList); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	foodToFoodTagLit, err := food_to_food_tag_service.Impl.ListByTagIDs(foodTagIdList.FoodTagIdList)
	if err != nil {
		return nil
	}

	var foodIDList []string
	for i := range foodToFoodTagLit {
		if utils.Contains(foodIDList, foodToFoodTagLit[i].FoodID) {
			continue
		}
		foodIDList = append(foodIDList, foodToFoodTagLit[i].FoodID)

	}

	foodList, err := food_service.Impl.ListByIDs(foodIDList)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, foodList)
}

func handlerListByFoodTagListAndConsumerTagId(c echo.Context) error {
	req := new(food_domain.ListByFoodTagListAndConsumerTagIdReq)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	foodToFoodTagLit, err := food_to_food_tag_service.Impl.ListByTagIDs(req.FoodTagIdList)
	if err != nil {
		return nil
	}

	var foodIDs []string
	for i := range foodToFoodTagLit {
		if utils.Contains(foodIDs, foodToFoodTagLit[i].FoodID) {
			continue
		}
		foodIDs = append(foodIDs, foodToFoodTagLit[i].FoodID)

	}

	foodList, err := food_service.Impl.ListByIDs(foodIDs)
	if err != nil {
		return err
	}

	// 如果是谨慎食用直接返回
	if req.EatMode == constant.EatModeWarning {
		return c.JSON(http.StatusOK, foodList)
	}

	// 如果是推荐食用需要过滤一下谨慎食用的标签
	foodTagList, err := food_to_food_tag_service.Impl.ListByFoodIDs(foodIDs)

	foodTagIDs := make([]string, 0)
	for i := range foodTagList {
		if utils.Contains(foodTagIDs, foodTagList[i].FoodTagID) {
			continue
		}
		foodTagIDs = append(foodTagIDs, foodTagList[i].FoodTagID)
	}

	consumerTagToFoodTagList, err := consumer_tag_to_food_tag_service.Impl.ListByFoodTagIDsAndConsumerTagIDAndEatMode(foodTagIDs, req.ConsumerTagId, constant.EatModeWarning)
	if err != nil {
		return err
	}

	warningFoodTagIDs := make([]string, 0)
	for _, consumerTagToFoodTag := range consumerTagToFoodTagList {
		if utils.Contains(warningFoodTagIDs, consumerTagToFoodTag.FoodTagID) {
			continue
		}
		warningFoodTagIDs = append(warningFoodTagIDs, consumerTagToFoodTag.FoodTagID)
	}

	warningFoodID := make([]string, 0)
	for _, foodTag := range foodTagList {
		if utils.Contains(warningFoodID, foodTag.FoodID) {
			continue
		}
		warningFoodID = append(warningFoodID, foodTag.FoodID)
	}

	res := removeWarningFood(foodList, warningFoodID)

	return c.JSON(http.StatusOK, res)
}

// 从 foodList 中删除 id 在 warningFoodID 中的食物
func removeWarningFood(foodList []*food_domain.Food, warningFoodID []string) []*food_domain.Food {
	res := make([]*food_domain.Food, 0)
	for i := range foodList {
		if utils.Contains(warningFoodID, foodList[i].ID) {
			continue
		}
		res = append(res, foodList[i])
	}
	return res
}
