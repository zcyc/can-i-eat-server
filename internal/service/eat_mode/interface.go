package eat_mode_service

import eat_mode_domain "can-i-eat/internal/domain/eat_mode"

type EatModeService interface {
	List(size int64, page int64) (*eat_mode_domain.ListResp, error)
	Detail(id string) (*eat_mode_domain.EatMode, error)
	Create(t *eat_mode_domain.EatMode) (string, error)
	BatchCreate(t []*eat_mode_domain.EatMode) error
	Update(t *eat_mode_domain.EatMode) error
	Delete(t *eat_mode_domain.EatMode) error
	ListByIDs(ids []string) ([]*eat_mode_domain.EatMode, error)
}
