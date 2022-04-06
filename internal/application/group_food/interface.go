package group_food_application

import (
	food_domain "can-i-eat/internal/domain/food"
	"github.com/labstack/echo/v4"
)

type GroupFoodApplication interface {
	ListFoodByGroup(c echo.Context) ([]*food_domain.Food, error)
	ListFoodByConsumer(c echo.Context) ([]*food_domain.Food, error)
}
