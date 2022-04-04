package consumer_group_service

import consumer_group_domain "can-i-eat/internal/domain/consumer_group"

type ConsumerGroupService interface {
	ListForPage(size int64, page int64) (*consumer_group_domain.ListResp, error)
	FoodDetail(id int64) (*consumer_group_domain.ConsumerGroup, error)
	Create(food *consumer_group_domain.ConsumerGroup) (uint64, error)
	Update(food *consumer_group_domain.ConsumerGroup) error
	Delete(food *consumer_group_domain.ConsumerGroup) error
}
