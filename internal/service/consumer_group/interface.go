package consumer_group_service

import consumer_group_domain "can-i-eat/internal/domain/consumer_group"

type ConsumerGroupService interface {
	List(size int64, page int64) (*consumer_group_domain.ListResp, error)
	Detail(id string) (*consumer_group_domain.ConsumerGroup, error)
	Create(t *consumer_group_domain.ConsumerGroup) (string, error)
	Update(t *consumer_group_domain.ConsumerGroup) error
	Delete(t *consumer_group_domain.ConsumerGroup) error
	ListByConsumer(id string) ([]*consumer_group_domain.ConsumerGroup, error)
}
