package food_to_food_tag_service

import (
	food_tag_domain "can-i-eat/internal/domain/food_tag"
	food_to_food_tag_domain "can-i-eat/internal/domain/food_to_food_tag"
)

type FoodToFoodTagService interface {
	List(size int64, page int64) (*food_to_food_tag_domain.ListResp, error)
	Detail(id string) (*food_to_food_tag_domain.FoodToFoodTag, error)
	Create(t *food_to_food_tag_domain.FoodToFoodTag) (string, error)
	Bind(t []*food_tag_domain.FoodTag, foodToFoodTagMap map[string][]string) error
	Update(t *food_to_food_tag_domain.FoodToFoodTag) error
	Delete(t *food_to_food_tag_domain.FoodToFoodTag) error
	ListByTagList(ids []string) ([]*food_to_food_tag_domain.FoodToFoodTag, error)
}
