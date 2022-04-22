package consumer_tag_to_food_tag_application

import (
	food_domain "can-i-eat/internal/domain/food"
)

type ConsumerTagToFoodTagApplication interface {
	ListFoodByConsumerTag(consumerTagID string) ([]*food_domain.Food, error)
	ListFoodByConsumer(consumerID string) ([]*food_domain.Food, error)
}
