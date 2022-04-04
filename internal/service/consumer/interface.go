package consumer_service

import (
	consumer_domain "can-i-eat/internal/domain/consumer"
)

type ConsumerService interface {
	ListForPage(size int64, page int64) (*consumer_domain.ListResp, error)
	FoodDetail(id int64) (*consumer_domain.Consumer, error)
	Create(food *consumer_domain.Consumer) (uint64, error)
	Update(food *consumer_domain.Consumer) error
	Delete(food *consumer_domain.Consumer) error
}
