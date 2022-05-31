package food_facade

import (
	"can-i-eat/common/constant"
	string_util "can-i-eat/common/util/string"
	food_domain "can-i-eat/internal/domain/food"
	consumer_tag_to_food_tag_service "can-i-eat/internal/service/consumer_tag_to_food_tag"
	food_service "can-i-eat/internal/service/food"
	food_to_food_tag_service "can-i-eat/internal/service/food_to_food_tag"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/utils"
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

	resp, err := food_service.Impl.List(size, page)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func handlerDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.String(http.StatusOK, "参数错误")
		return
	}
	food, _ := food_service.Impl.Detail(id)
	c.JSON(http.StatusOK, food)
}

func handlerCreate(c *gin.Context) {
	food := new(food_domain.Food)
	if err := c.Bind(food); err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	id, err := food_service.Impl.Create(food)
	if err != nil {
		c.String(http.StatusOK, "创建失败")
		return
	}
	c.JSON(http.StatusOK, id)
}

func handlerUpdate(c *gin.Context) {
	food := new(food_domain.Food)
	if err := c.Bind(food); err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	err := food_service.Impl.Update(food)
	if err != nil {
		c.String(http.StatusOK, "更新失败")
		return
	}
	c.JSON(http.StatusOK, "更新成功")
}

func handlerDelete(c *gin.Context) {
	food := new(food_domain.Food)
	if err := c.Bind(food); err != nil {
		c.String(http.StatusOK, "更新失败")
		return
	}
	err := food_service.Impl.Delete(food)
	if err != nil {
		c.String(http.StatusOK, "更新失败")
		return
	}
	c.String(http.StatusOK, "更新成功")
}

func handlerListByFoodTagList(c *gin.Context) {
	foodTagIdList := new(food_domain.ListByFoodTagListReq)
	if err := c.Bind(foodTagIdList); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	foodToFoodTagLit, err := food_to_food_tag_service.Impl.ListByTagIDs(foodTagIdList.FoodTagIdList)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
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
		c.String(http.StatusOK, err.Error())
		return
	}

	c.JSON(http.StatusOK, foodList)
}

func handlerListByFoodTagListAndConsumerTagId(c *gin.Context) {
	req := new(food_domain.ListByFoodTagListAndConsumerTagIdReq)
	if err := c.Bind(req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	foodToFoodTagLit, err := food_to_food_tag_service.Impl.ListByTagIDs(req.FoodTagIdList)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
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
		c.String(http.StatusOK, err.Error())
		return
	}

	// 如果是谨慎食用直接返回
	if req.EatMode == constant.EatModeWarning {
		c.JSON(http.StatusOK, foodList)
		return
	}

	// 如果是推荐食用需要过滤一下谨慎食用的标签
	foodTagList, err := food_to_food_tag_service.Impl.ListByFoodIDs(foodIDs)

	// 取出食品标签id
	foodTagIDs := make([]string, 0)
	for i := range foodTagList {
		if utils.Contains(foodTagIDs, foodTagList[i].FoodTagID) {
			continue
		}
		foodTagIDs = append(foodTagIDs, foodTagList[i].FoodTagID)
	}

	// 获取当前用户标签禁止食用的食品标签
	consumerTagToFoodTagList, err := consumer_tag_to_food_tag_service.Impl.ListByFoodTagIDsAndConsumerTagIDAndEatMode(foodTagIDs, req.ConsumerTagId, constant.EatModeWarning)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}

	// 当前用户标签禁止食用的食品标签 ids
	warningFoodTagIDs := make([]string, 0)
	for _, consumerTagToFoodTag := range consumerTagToFoodTagList {
		if utils.Contains(warningFoodTagIDs, consumerTagToFoodTag.FoodTagID) {
			continue
		}
		warningFoodTagIDs = append(warningFoodTagIDs, consumerTagToFoodTag.FoodTagID)
	}

	// 整理食物到禁止食用的标签的关系
	foodToFoodTagListMap := make(map[string][]string, 0)
	for _, foodTag := range foodTagList {
		// 判断是否在禁止食用标签中
		if utils.Contains(warningFoodTagIDs, foodTag.FoodTagID) {
			if utils.Contains(foodToFoodTagListMap[foodTag.FoodID], foodTag.FoodTagID) {
				continue
			}
			foodToFoodTagListMap[foodTag.FoodID] = append(foodToFoodTagListMap[foodTag.FoodID], foodTag.FoodTagID)
		}
	}

	// 整理禁止食用的食物
	warningFoodID := make([]string, 0)
	for foodID, foodTags := range foodToFoodTagListMap {
		if len(foodTags) == 0 {
			continue
		}
		warningFoodID = append(warningFoodID, foodID)
	}

	// 删除禁止食用的食物
	res := removeWarningFood(foodList, warningFoodID)

	c.JSON(http.StatusOK, res)
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
