package food_service

import (
	food_domain "can-i-eat/internal/domain/food"
	"can-i-eat/internal/infrastructure/model"
	mysql_infrastructure "can-i-eat/internal/infrastructure/mysql"
	"can-i-eat/internal/repo/food"
	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
)

var Impl FoodService = &foodImpl{}

type foodImpl struct {
}

func (f foodImpl) ListForPage(size int64, page int64) (*food_domain.ListResp, error) {
	resp := new(food_domain.ListResp)
	foodMgr := model.FoodMgr(mysql_infrastructure.Get())
	foodPage := model.NewPage(size, page)
	result, err := foodMgr.SelectPage(foodPage)

	foodList := make([]*food_domain.Food, 0)
	for _, foodRepo := range result.GetRecords().([]food_repo.FoodRepo) {
		food := new(food_domain.Food)
		_ = copier.Copy(&food, &foodRepo)
		foodList = append(foodList, food)
	}
	resp.Items = foodList
	resp.Current = int(result.GetCurrent())
	resp.Size = int(result.GetSize())
	resp.Total = int(result.GetTotal())

	return resp, err
}

func (f foodImpl) FoodDetail(id int64) (*food_domain.Food, error) {
	foodRepoList := make([]*food_repo.FoodRepo, 0)
	foodMgr := model.FoodMgr(mysql_infrastructure.Get())
	err := foodMgr.Where("id=?", id).Limit(1).Find(&foodRepoList).Error
	if err != nil {
		return nil, err
	}
	food := new(food_domain.Food)
	_ = copier.Copy(&food, foodRepoList[0])
	return food, nil
}

func (f foodImpl) Create(food *food_domain.Food) (uint64, error) {
	foodMgr := model.FoodMgr(mysql_infrastructure.Get())
	err := foodMgr.Omit("create_time", "update_time").Create(food).Error
	if err != nil {
		return 0, err
	}
	log.Infof("create food success: %d", food.ID)
	return food.ID, nil
}