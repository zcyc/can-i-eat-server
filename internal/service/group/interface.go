package group_service

import (
	group_domain "can-i-eat/internal/domain/group"
)

type GroupService interface {
	List(size int64, page int64) (*group_domain.ListResp, error)
	Detail(id int64) (*group_domain.Group, error)
	Create(t *group_domain.Group) (uint64, error)
	Update(t *group_domain.Group) error
	Delete(t *group_domain.Group) error
}
