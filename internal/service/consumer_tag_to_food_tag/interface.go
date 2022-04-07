package consumer_tag_to_food_tag_service

import (
	consumer_tag_to_food_tag_domain "can-i-eat/internal/domain/consumer_tag_to_food_tag"
)

type ConsumerTagToFoodTagService interface {
	List(size int64, page int64) (*consumer_tag_to_food_tag_domain.ListResp, error)
	Detail(id string) (*consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag, error)
	Create(t *consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag) (string, error)
	Update(t *consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag) error
	Delete(t *consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag) error
	ListByConsumerTag(id string) ([]*consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag, error)
	ListByConsumerTagIDs(id []string) ([]*consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag, error)
}
