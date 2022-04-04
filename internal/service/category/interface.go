package category_service

import category_domain "can-i-eat/internal/domain/category"

type CategoryService interface {
	ListForPage(size int64, page int64) (*category_domain.ListResp, error)
	FoodDetail(id int64) (*category_domain.Category, error)
	Create(food *category_domain.Category) (uint64, error)
	Update(food *category_domain.Category) error
	Delete(food *category_domain.Category) error
}
