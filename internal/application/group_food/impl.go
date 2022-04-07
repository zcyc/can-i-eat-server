package consumer_tag_to_food_tag_application

import (
	food_domain "can-i-eat/internal/domain/food"
	consumer_tag_to_food_tag_service "can-i-eat/internal/service/consumer_tag_to_food_tag"
	consumer_to_consumer_tag_service "can-i-eat/internal/service/consumer_to_consumer_tag"
	food_service "can-i-eat/internal/service/food"
	food_to_food_tag_service "can-i-eat/internal/service/food_to_food_tag"
	"errors"
	"github.com/labstack/echo/v4"
)

var Impl ConsumerTagToFoodTagApplication = &consumerTagToFoodTagApplicationImpl{}

type consumerTagToFoodTagApplicationImpl struct {
}

func (g consumerTagToFoodTagApplicationImpl) ListFoodByConsumerTag(c echo.Context) ([]*food_domain.Food, error) {
	consumerTagID := c.QueryParam("consumer-tag-id")
	if consumerTagID == "" {
		return nil, errors.New("参数错误")
	}

	consumerTagToFoodTagList, err := consumer_tag_to_food_tag_service.Impl.ListByConsumerTag(consumerTagID)
	if err != nil {
		return nil, err
	}

	var tagIDList []string
	for i := range consumerTagToFoodTagList {
		tagIDList = append(tagIDList, consumerTagToFoodTagList[i].FoodTagID)
	}

	foodToFoodTagList, err := food_to_food_tag_service.Impl.ListByTagList(tagIDList)
	if err != nil {
		return nil, err
	}

	var foodIDList []string
	for i := range foodToFoodTagList {
		foodIDList = append(foodIDList, foodToFoodTagList[i].FoodID)
	}

	foodList, err := food_service.Impl.ListByIDs(foodIDList)
	return foodList, nil
}

func (g consumerTagToFoodTagApplicationImpl) ListFoodByConsumer(c echo.Context) ([]*food_domain.Food, error) {
	consumerID := c.QueryParam("consumer-id")
	consumerToConsumerTagList, err := consumer_to_consumer_tag_service.Impl.ListByConsumer(consumerID)
	if err != nil {
		return nil, err
	}

	var consumerTagIDList []string
	for i := range consumerToConsumerTagList {
		consumerTagIDList = append(consumerTagIDList, consumerToConsumerTagList[i].ConsumerTagID)
	}

	consumerTagToFoodTagList, err := consumer_tag_to_food_tag_service.Impl.ListByConsumerTagIDs(consumerTagIDList)
	if err != nil {
		return nil, err
	}

	var tagIDList []string
	for i := range consumerTagToFoodTagList {
		tagIDList = append(tagIDList, consumerTagToFoodTagList[i].FoodTagID)
	}

	foodToFoodTagList, err := food_to_food_tag_service.Impl.ListByTagList(tagIDList)
	if err != nil {
		return nil, err
	}

	var foodIDList []string
	for i := range foodToFoodTagList {
		foodIDList = append(foodIDList, foodToFoodTagList[i].FoodID)
	}

	foodList, err := food_service.Impl.ListByIDs(foodIDList)
	return foodList, nil
}
