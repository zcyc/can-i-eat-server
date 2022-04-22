package food_tag_service

import (
	food_tag_domain "can-i-eat/internal/domain/food_tag"
)

type FoodTagService interface {
	List(size int64, page int64) (*food_tag_domain.ListResp, error)
	Detail(id int64) (*food_tag_domain.FoodTag, error)
	Create(t *food_tag_domain.FoodTag) (string, error)
	BatchCreate(t []*food_tag_domain.FoodTag) error
	Update(t *food_tag_domain.FoodTag) error
	Delete(t *food_tag_domain.FoodTag) error
}
