package consumer_tag_to_food_tag_application

import (
	food_domain "can-i-eat/internal/domain/food"
	"github.com/labstack/echo/v4"
)

type ConsumerTagToFoodTagApplication interface {
	ListFoodByConsumerTag(c echo.Context) ([]*food_domain.Food, error)
	ListFoodByConsumer(c echo.Context) ([]*food_domain.Food, error)
}
