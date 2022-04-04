package tag_service

import (
	tag_domain "can-i-eat/internal/domain/tag"
)

type TagService interface {
	ListForPage(size int64, page int64) (*tag_domain.ListResp, error)
	FoodDetail(id int64) (*tag_domain.Tag, error)
	Create(food *tag_domain.Tag) (uint64, error)
	Update(food *tag_domain.Tag) error
	Delete(food *tag_domain.Tag) error
}
