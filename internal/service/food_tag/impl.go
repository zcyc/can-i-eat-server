package food_tag_service

import (
	"can-i-eat/common/constant"
	food_tag_domain "can-i-eat/internal/domain/food_tag"
	"can-i-eat/internal/infrastructure/model"
	mysql_infrastructure "can-i-eat/internal/infrastructure/mysql"
	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
)

var Impl FoodTagService = &foodTagServiceImpl{}

type foodTagServiceImpl struct {
}

func (f foodTagServiceImpl) ListByTagList(ids []string) ([]*food_tag_domain.FoodTag, error) {
	foodTagDaoList := make([]*model.FoodTag, 0)
	foodTagMgr := model.FoodTagMgr(mysql_infrastructure.Get())
	err := foodTagMgr.Where("tag_id in ?", ids).Find(&foodTagDaoList).Error
	if err != nil {
		return nil, err
	}
	foodTagList := make([]*food_tag_domain.FoodTag, 0)
	for _, foodTagRepo := range foodTagDaoList {
		foodTag := new(food_tag_domain.FoodTag)
		_ = copier.Copy(&foodTag, &foodTagRepo)
		foodTagList = append(foodTagList, foodTag)
	}

	return foodTagList, nil
}

func (f foodTagServiceImpl) Delete(foodTag *food_tag_domain.FoodTag) error {
	foodTagMgr := model.FoodTagMgr(mysql_infrastructure.Get())
	err := foodTagMgr.Update("flag", constant.Deleted).Where("id=?", foodTag.ID).Error
	if err != nil {
		return err
	}
	log.Infof("delete food success: %d", foodTag.ID)
	return nil
}

func (f foodTagServiceImpl) Update(foodTag *food_tag_domain.FoodTag) error {
	foodTagMgr := model.FoodTagMgr(mysql_infrastructure.Get())
	err := foodTagMgr.Save(foodTag).Error
	if err != nil {
		return err
	}
	log.Infof("update food success: %d", foodTag.ID)
	return nil
}

func (f foodTagServiceImpl) List(size int64, page int64) (*food_tag_domain.ListResp, error) {
	resp := new(food_tag_domain.ListResp)
	foodTagMgr := model.FoodTagMgr(mysql_infrastructure.Get())
	foodTagPage := model.NewPage(size, page)
	result, err := foodTagMgr.SelectPage(foodTagPage, foodTagMgr.WithFlag(constant.Normal), foodTagMgr.WithActive(constant.Activated))
	if err != nil {
		return nil, err
	}
	foodTagList := make([]*food_tag_domain.FoodTag, 0)
	for _, foodTagRepo := range result.GetRecords().([]model.FoodTag) {
		foodTag := new(food_tag_domain.FoodTag)
		_ = copier.Copy(&foodTag, &foodTagRepo)
		foodTagList = append(foodTagList, foodTag)
	}
	resp.Items = foodTagList
	resp.Current = int(result.GetCurrent())
	resp.Size = int(result.GetSize())
	resp.Total = int(result.GetTotal())

	return resp, nil
}

func (f foodTagServiceImpl) Detail(id string) (*food_tag_domain.FoodTag, error) {
	foodRepoList := make([]*model.FoodTag, 0)
	foodTagMgr := model.FoodTagMgr(mysql_infrastructure.Get())
	err := foodTagMgr.Where("id=?", id).Limit(1).Find(&foodRepoList).Error
	if err != nil {
		return nil, err
	}
	food := new(food_tag_domain.FoodTag)
	_ = copier.Copy(&food, foodRepoList[0])
	return food, nil
}

func (f foodTagServiceImpl) Create(t *food_tag_domain.FoodTag) (string, error) {
	foodTagDao := new(model.FoodTag)
	_ = copier.Copy(foodTagDao, t)
	foodTagDao.ID = foodTagDao.FoodID + "_" + foodTagDao.TagID
	foodTagMgr := model.FoodTagMgr(mysql_infrastructure.Get())
	err := foodTagMgr.Omit("create_time", "update_time").Create(foodTagDao).Error
	if err != nil {
		return "", err
	}
	log.Infof("create food success: %d", foodTagDao.ID)
	return foodTagDao.ID, nil
}
