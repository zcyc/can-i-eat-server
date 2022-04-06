package group_food_application

import (
	food_domain "can-i-eat/internal/domain/food"
	consumer_group_service "can-i-eat/internal/service/consumer_group"
	food_service "can-i-eat/internal/service/food"
	food_tag_service "can-i-eat/internal/service/food_tag"
	group_tag_service "can-i-eat/internal/service/group_tag"
	"errors"
	"github.com/labstack/echo/v4"
)

var Impl GroupFoodApplication = &groupFoodApplicationImpl{}

type groupFoodApplicationImpl struct {
}

func (g groupFoodApplicationImpl) ListFoodByGroup(c echo.Context) ([]*food_domain.Food, error) {
	groupID := c.QueryParam("group-id")
	if groupID == "" {
		return nil, errors.New("参数错误")
	}
	groupTagList, err := group_tag_service.Impl.ListByGroup(groupID)
	if err != nil {
		return nil, err
	}

	var tagIDList []string
	for i := range groupTagList {
		tagIDList = append(tagIDList, groupTagList[i].TagID)
	}

	foodTagList, err := food_tag_service.Impl.ListByTagList(tagIDList)
	if err != nil {
		return nil, err
	}

	var foodIDList []string
	for i := range foodTagList {
		foodIDList = append(foodIDList, foodTagList[i].FoodID)
	}

	foodList, err := food_service.Impl.ListByIDs(foodIDList)
	return foodList, nil
}

func (g groupFoodApplicationImpl) ListFoodByConsumer(c echo.Context) ([]*food_domain.Food, error) {
	consumerID := c.QueryParam("consumer-id")
	consumerGroupList, err := consumer_group_service.Impl.ListByConsumer(consumerID)
	if err != nil {
		return nil, err
	}

	var groupIDList []string
	for i := range consumerGroupList {
		groupIDList = append(groupIDList, consumerGroupList[i].GroupID)
	}

	groupTagList, err := group_tag_service.Impl.ListByGroupIDs(groupIDList)
	if err != nil {
		return nil, err
	}

	var tagIDList []string
	for i := range groupTagList {
		tagIDList = append(tagIDList, groupTagList[i].TagID)
	}

	foodTagList, err := food_tag_service.Impl.ListByTagList(tagIDList)
	if err != nil {
		return nil, err
	}

	var foodIDList []string
	for i := range foodTagList {
		foodIDList = append(foodIDList, foodTagList[i].FoodID)
	}

	foodList, err := food_service.Impl.ListByIDs(foodIDList)
	return foodList, nil
}
