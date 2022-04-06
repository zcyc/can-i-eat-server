package food_service

import (
	food_domain "can-i-eat/internal/domain/food"
)

type FoodService interface {
	List(size int64, page int64) (*food_domain.ListResp, error)
	Detail(id string) (*food_domain.Food, error)
	Create(t *food_domain.Food) (string, error)
	Update(t *food_domain.Food) error
	Delete(t *food_domain.Food) error
	ListByIDs(ids []string) ([]*food_domain.Food, error)
}
