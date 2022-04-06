package category_service

import category_domain "can-i-eat/internal/domain/category"

type CategoryService interface {
	List(size int64, page int64) (*category_domain.ListResp, error)
	Detail(id int64) (*category_domain.Category, error)
	Create(t *category_domain.Category) (uint64, error)
	Update(t *category_domain.Category) error
	Delete(t *category_domain.Category) error
}
