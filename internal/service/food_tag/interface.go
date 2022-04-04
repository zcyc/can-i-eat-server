package food_tag_service

import (
	food_tag_domain "can-i-eat/internal/domain/food_tag"
)

type FoodTagService interface {
	ListForPage(size int64, page int64) (*food_tag_domain.ListResp, error)
	FoodDetail(id int64) (*food_tag_domain.FoodTag, error)
	Create(food *food_tag_domain.FoodTag) (uint64, error)
	Update(food *food_tag_domain.FoodTag) error
	Delete(food *food_tag_domain.FoodTag) error
}
