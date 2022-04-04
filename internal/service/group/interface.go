package group_service

import (
	group_domain "can-i-eat/internal/domain/group"
)

type GroupService interface {
	ListForPage(size int64, page int64) (*group_domain.ListResp, error)
	FoodDetail(id int64) (*group_domain.Group, error)
	Create(food *group_domain.Group) (uint64, error)
	Update(food *group_domain.Group) error
	Delete(food *group_domain.Group) error
}
