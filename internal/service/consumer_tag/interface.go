package consumer_tag_service

import consumer_tag_domain "can-i-eat/internal/domain/consumer_tag"

type ConsumerTagService interface {
	List(size int64, page int64) (*consumer_tag_domain.ListResp, error)
	Detail(id int64) (*consumer_tag_domain.ConsumerTag, error)
	Create(t *consumer_tag_domain.ConsumerTag) (string, error)
	Update(t *consumer_tag_domain.ConsumerTag) error
	Delete(t *consumer_tag_domain.ConsumerTag) error
	ListByIDs(id []int64) ([]*consumer_tag_domain.ConsumerTag, error)
}
