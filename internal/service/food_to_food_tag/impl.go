package food_to_food_tag_service

import (
	"can-i-eat/common/constant"
	food_tag_domain "can-i-eat/internal/domain/food_tag"
	food_to_food_tag_domain "can-i-eat/internal/domain/food_to_food_tag"
	"can-i-eat/internal/infrastructure/model"
	mysql_infrastructure "can-i-eat/internal/infrastructure/mysql"
	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm/clause"
)

var Impl FoodToFoodTagService = &foodToFoodTagServiceImpl{}

type foodToFoodTagServiceImpl struct {
}

func (f foodToFoodTagServiceImpl) Bind(FoodTagList []*food_tag_domain.FoodTag, foodToFoodTagMap map[string][]string) error {
	foodToFoodTagDaoList := make([]*model.FoodToFoodTag, 0)
	for _, foodTag := range FoodTagList {
		for foodID, FoodTagIDList := range foodToFoodTagMap {
			for i := range FoodTagIDList {
				if foodTag.ID == FoodTagIDList[i] {
					foodToFoodTagDao := new(model.FoodToFoodTag)
					foodToFoodTagDao.Active = constant.Activated
					foodToFoodTagDao.Flag = constant.Normal
					foodToFoodTagDao.ID = foodID + "_" + FoodTagIDList[i]
					foodToFoodTagDao.FoodTagID = FoodTagIDList[i]
					foodToFoodTagDao.FoodID = foodID
					foodToFoodTagDaoList = append(foodToFoodTagDaoList, foodToFoodTagDao)
				}
			}
		}
	}

	// 执行批量
	foodToFoodTagMgr := model.FoodToFoodTagMgr(mysql_infrastructure.Get())
	err := foodToFoodTagMgr.Omit("create_time", "update_time").Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&foodToFoodTagDaoList, 100).Error
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (f foodToFoodTagServiceImpl) ListByTagList(ids []string) ([]*food_to_food_tag_domain.FoodToFoodTag, error) {
	foodToFoodTagDaoList := make([]*model.FoodToFoodTag, 0)
	foodToFoodTagMgr := model.FoodToFoodTagMgr(mysql_infrastructure.Get())
	err := foodToFoodTagMgr.Where("food_tag_id in ?", ids).Find(&foodToFoodTagDaoList).Error
	if err != nil {
		return nil, err
	}
	foodToFoodTagList := make([]*food_to_food_tag_domain.FoodToFoodTag, 0)
	for _, foodToFoodTagRepo := range foodToFoodTagDaoList {
		foodToFoodTag := new(food_to_food_tag_domain.FoodToFoodTag)
		_ = copier.Copy(&foodToFoodTag, &foodToFoodTagRepo)
		foodToFoodTagList = append(foodToFoodTagList, foodToFoodTag)
	}

	return foodToFoodTagList, nil
}

func (f foodToFoodTagServiceImpl) Delete(foodToFoodTag *food_to_food_tag_domain.FoodToFoodTag) error {
	foodToFoodTagMgr := model.FoodToFoodTagMgr(mysql_infrastructure.Get())
	err := foodToFoodTagMgr.Update("flag", constant.Deleted).Where("id=?", foodToFoodTag.ID).Error
	if err != nil {
		return err
	}
	log.Infof("delete food success: %s", foodToFoodTag.ID)
	return nil
}

func (f foodToFoodTagServiceImpl) Update(foodToFoodTag *food_to_food_tag_domain.FoodToFoodTag) error {
	foodToFoodTagMgr := model.FoodToFoodTagMgr(mysql_infrastructure.Get())
	err := foodToFoodTagMgr.Save(foodToFoodTag).Error
	if err != nil {
		return err
	}
	log.Infof("update food success: %s", foodToFoodTag.ID)
	return nil
}

func (f foodToFoodTagServiceImpl) List(size int64, page int64) (*food_to_food_tag_domain.ListResp, error) {
	resp := new(food_to_food_tag_domain.ListResp)
	foodToFoodTagMgr := model.FoodToFoodTagMgr(mysql_infrastructure.Get())
	foodToFoodTagPage := model.NewPage(size, page)
	result, err := foodToFoodTagMgr.SelectPage(foodToFoodTagPage, foodToFoodTagMgr.WithFlag(constant.Normal), foodToFoodTagMgr.WithActive(constant.Activated))
	if err != nil {
		return nil, err
	}
	foodToFoodTagList := make([]*food_to_food_tag_domain.FoodToFoodTag, 0)
	for _, foodToFoodTagRepo := range result.GetRecords().([]model.FoodToFoodTag) {
		foodToFoodTag := new(food_to_food_tag_domain.FoodToFoodTag)
		_ = copier.Copy(&foodToFoodTag, &foodToFoodTagRepo)
		foodToFoodTagList = append(foodToFoodTagList, foodToFoodTag)
	}
	resp.Items = foodToFoodTagList
	resp.Current = int(result.GetCurrent())
	resp.Size = int(result.GetSize())
	resp.Total = int(result.GetTotal())

	return resp, nil
}

func (f foodToFoodTagServiceImpl) Detail(id string) (*food_to_food_tag_domain.FoodToFoodTag, error) {
	foodRepoList := make([]*model.FoodToFoodTag, 0)
	foodToFoodTagMgr := model.FoodToFoodTagMgr(mysql_infrastructure.Get())
	err := foodToFoodTagMgr.Where("id=?", id).Limit(1).Find(&foodRepoList).Error
	if err != nil {
		return nil, err
	}
	food := new(food_to_food_tag_domain.FoodToFoodTag)
	_ = copier.Copy(&food, foodRepoList[0])
	return food, nil
}

func (f foodToFoodTagServiceImpl) Create(t *food_to_food_tag_domain.FoodToFoodTag) (string, error) {
	foodToFoodTagDao := new(model.FoodToFoodTag)
	_ = copier.Copy(foodToFoodTagDao, t)
	foodToFoodTagDao.ID = foodToFoodTagDao.FoodID + "_" + foodToFoodTagDao.FoodTagID
	foodToFoodTagMgr := model.FoodToFoodTagMgr(mysql_infrastructure.Get())
	err := foodToFoodTagMgr.Omit("create_time", "update_time").Create(foodToFoodTagDao).Error
	if err != nil {
		return "", err
	}
	log.Infof("create food success: %s", foodToFoodTagDao.ID)
	return foodToFoodTagDao.ID, nil
}
