package food_service

import (
	food_domain "can-i-eat/internal/domain/food"
)

type FoodService interface {
	ListForPage(size int64, page int64) (*food_domain.ListResp, error)
	FoodDetail(id int64) (*food_domain.Food, error)
	Create(food *food_domain.Food) (uint64, error)
	Update(food *food_domain.Food) error
	Delete(food *food_domain.Food) error
}
