package tag_service

import (
	tag_domain "can-i-eat/internal/domain/tag"
)

type TagService interface {
	List(size int64, page int64) (*tag_domain.ListResp, error)
	Detail(id int64) (*tag_domain.Tag, error)
	Create(t *tag_domain.Tag) (string, error)
	Update(t *tag_domain.Tag) error
	Delete(t *tag_domain.Tag) error
}
