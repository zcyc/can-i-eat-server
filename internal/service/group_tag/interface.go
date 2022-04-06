package group_tag_service

import (
	group_tag_domain "can-i-eat/internal/domain/group_tag"
)

type GroupTagService interface {
	List(size int64, page int64) (*group_tag_domain.ListResp, error)
	Detail(id int64) (*group_tag_domain.GroupTag, error)
	Create(t *group_tag_domain.GroupTag) (uint64, error)
	Update(t *group_tag_domain.GroupTag) error
	Delete(t *group_tag_domain.GroupTag) error
}
