package consumer_to_consumer_tag_service

import consumer_to_consumer_tag_domain "can-i-eat/internal/domain/consumer_to_consumer_tag"

type ConsumerToConsumerTagService interface {
	List(size int64, page int64) (*consumer_to_consumer_tag_domain.ListResp, error)
	Detail(id string) (*consumer_to_consumer_tag_domain.ConsumerToConsumerTag, error)
	Create(t *consumer_to_consumer_tag_domain.ConsumerToConsumerTag) (string, error)
	Update(t *consumer_to_consumer_tag_domain.ConsumerToConsumerTag) error
	Delete(t *consumer_to_consumer_tag_domain.ConsumerToConsumerTag) error
	ListByConsumer(id string) ([]*consumer_to_consumer_tag_domain.ConsumerToConsumerTag, error)
}
