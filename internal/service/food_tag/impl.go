package food_tag_service

import (
	"can-i-eat/common/constant"
	id_util "can-i-eat/common/util/id"
	food_tag_domain "can-i-eat/internal/domain/food_tag"
	"can-i-eat/internal/infrastructure/model"
	mysql_infrastructure "can-i-eat/internal/infrastructure/mysql"
	food_tag_repo "can-i-eat/internal/repo/food_tag"
	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
)

var Impl FoodTagService = &foodTagServiceImpl{}

type foodTagServiceImpl struct {
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

func (f foodTagServiceImpl) ListForPage(size int64, page int64) (*food_tag_domain.ListResp, error) {
	resp := new(food_tag_domain.ListResp)
	foodTagMgr := model.FoodTagMgr(mysql_infrastructure.Get())
	foodTagPage := model.NewPage(size, page)
	result, err := foodTagMgr.SelectPage(foodTagPage, foodTagMgr.WithFlag(constant.Normal), foodTagMgr.WithActive(constant.Activated))
	if err != nil {
		return nil, err
	}
	foodTagList := make([]*food_tag_domain.FoodTag, 0)
	for _, foodTagRepo := range result.GetRecords().([]food_tag_repo.FoodTagDao) {
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

func (f foodTagServiceImpl) FoodDetail(id int64) (*food_tag_domain.FoodTag, error) {
	foodRepoList := make([]*food_tag_repo.FoodTagDao, 0)
	foodTagMgr := model.FoodTagMgr(mysql_infrastructure.Get())
	err := foodTagMgr.Where("id=?", id).Limit(1).Find(&foodRepoList).Error
	if err != nil {
		return nil, err
	}
	food := new(food_tag_domain.FoodTag)
	_ = copier.Copy(&food, foodRepoList[0])
	return food, nil
}

func (f foodTagServiceImpl) Create(foodTag *food_tag_domain.FoodTag) (uint64, error) {
	foodTagDao := new(food_tag_repo.FoodTagDao)
	_ = copier.Copy(foodTagDao, foodTag)
	foodTagDao.ID, _ = id_util.NextID()
	foodTagMgr := model.FoodTagMgr(mysql_infrastructure.Get())
	err := foodTagMgr.Omit("create_time", "update_time").Create(foodTagDao).Error
	if err != nil {
		return 0, err
	}
	log.Infof("create food success: %d", foodTagDao.ID)
	return foodTagDao.ID, nil
}