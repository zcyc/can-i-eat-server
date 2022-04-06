package food_service

import (
	"can-i-eat/common/constant"
	food_domain "can-i-eat/internal/domain/food"
	"can-i-eat/internal/infrastructure/model"
	mysql_infrastructure "can-i-eat/internal/infrastructure/mysql"
	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
	"github.com/mozillazg/go-pinyin"
	"strings"
)

var Impl FoodService = &foodServiceImpl{}

type foodServiceImpl struct {
}

func (f foodServiceImpl) ListByIDs(ids []string) ([]*food_domain.Food, error) {
	foodDaoList := make([]*model.Food, 0)
	foodMgr := model.FoodMgr(mysql_infrastructure.Get())
	err := foodMgr.Where("id in ?", ids).Find(&foodDaoList).Error
	if err != nil {
		return nil, err
	}
	foodList := make([]*food_domain.Food, 0)
	for _, foodRepo := range foodDaoList {
		food := new(food_domain.Food)
		_ = copier.Copy(&food, &foodRepo)
		foodList = append(foodList, food)
	}

	return foodList, nil
}

func (f foodServiceImpl) Delete(food *food_domain.Food) error {
	foodMgr := model.FoodMgr(mysql_infrastructure.Get())
	err := foodMgr.Update("flag", constant.Deleted).Where("id=?", food.ID).Error
	if err != nil {
		return err
	}
	log.Infof("delete food success: %s", food.ID)
	return nil
}

func (f foodServiceImpl) Update(food *food_domain.Food) error {
	foodMgr := model.FoodMgr(mysql_infrastructure.Get())
	err := foodMgr.Save(food).Error
	if err != nil {
		return err
	}
	log.Infof("update food success: %s", food.ID)
	return nil
}

func (f foodServiceImpl) List(size int64, page int64) (*food_domain.ListResp, error) {
	resp := new(food_domain.ListResp)
	foodMgr := model.FoodMgr(mysql_infrastructure.Get())
	foodPage := model.NewPage(size, page)
	result, err := foodMgr.SelectPage(foodPage, foodMgr.WithFlag(constant.Normal), foodMgr.WithActive(constant.Activated))
	if err != nil {
		return nil, err
	}
	foodList := make([]*food_domain.Food, 0)
	for _, foodRepo := range result.GetRecords().([]model.Food) {
		food := new(food_domain.Food)
		_ = copier.Copy(&food, &foodRepo)
		foodList = append(foodList, food)
	}
	resp.Items = foodList
	resp.Current = int(result.GetCurrent())
	resp.Size = int(result.GetSize())
	resp.Total = int(result.GetTotal())

	return resp, nil
}

func (f foodServiceImpl) Detail(id string) (*food_domain.Food, error) {
	foodRepoList := make([]*model.Food, 0)
	foodMgr := model.FoodMgr(mysql_infrastructure.Get())
	err := foodMgr.Where("id=?", id).Limit(1).Find(&foodRepoList).Error
	if err != nil {
		return nil, err
	}
	food := new(food_domain.Food)
	_ = copier.Copy(&food, foodRepoList[0])
	return food, nil
}

func (f foodServiceImpl) Create(t *food_domain.Food) (string, error) {
	foodDao := new(model.Food)
	_ = copier.Copy(foodDao, t)
	foodDao.ID = strings.Join(pinyin.LazyConvert(foodDao.Name, nil), "_")
	foodMgr := model.FoodMgr(mysql_infrastructure.Get())
	err := foodMgr.Omit("create_time", "update_time").Create(foodDao).Error
	if err != nil {
		return "", err
	}
	log.Infof("create food success: %s", foodDao.ID)
	return foodDao.ID, nil
}
