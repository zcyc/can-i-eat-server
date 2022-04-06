package consumer_group_service

import consumer_group_domain "can-i-eat/internal/domain/consumer_group"

type ConsumerGroupService interface {
	List(size int64, page int64) (*consumer_group_domain.ListResp, error)
	Detail(id int64) (*consumer_group_domain.ConsumerGroup, error)
	Create(consumerGroup *consumer_group_domain.ConsumerGroup) (uint64, error)
	Update(consumerGroup *consumer_group_domain.ConsumerGroup) error
	Delete(consumerGroup *consumer_group_domain.ConsumerGroup) error
}
