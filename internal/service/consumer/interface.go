package consumer_service

import (
	consumer_domain "can-i-eat/internal/domain/consumer"
)

type ConsumerService interface {
	List(size int64, page int64) (*consumer_domain.ListResp, error)
	Detail(id int64) (*consumer_domain.Consumer, error)
	Create(t *consumer_domain.Consumer) (uint64, error)
	Update(t *consumer_domain.Consumer) error
	Delete(t *consumer_domain.Consumer) error
}
