package food_tag_service

import (
	"can-i-eat/common/constant"
	food_tag_domain "can-i-eat/internal/domain/food_tag"
	"can-i-eat/internal/infrastructure/model"
	mysql_infrastructure "can-i-eat/internal/infrastructure/mysql"
	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
	"github.com/mozillazg/go-pinyin"
	"gorm.io/gorm/clause"
	"strings"
)

var Impl FoodTagService = &foodTagServiceImpl{}

type foodTagServiceImpl struct {
}

func (f foodTagServiceImpl) BatchCreate(t []*food_tag_domain.FoodTag) error {
	// domain 对象转 repo 对象
	foodTagDaoList := make([]*model.FoodTag, 0)
	for _, foodTag := range t {
		foodTagDao := new(model.FoodTag)
		_ = copier.Copy(foodTagDao, foodTag)
		foodTagDaoList = append(foodTagDaoList, foodTagDao)
	}

	// 执行批量
	foodTagMgr := model.FoodTagMgr(mysql_infrastructure.Get())
	err := foodTagMgr.Omit("create_time", "update_time").Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&foodTagDaoList, 100).Error
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (f foodTagServiceImpl) Delete(foodTag *food_tag_domain.FoodTag) error {
	foodTagMgr := model.FoodTagMgr(mysql_infrastructure.Get())
	err := foodTagMgr.Update("flag", constant.DataDeleted).Where("id=?", foodTag.ID).Error
	if err != nil {
		return err
	}
	log.Infof("delete food_to_food_tag success: %s", foodTag.ID)
	return nil
}

func (f foodTagServiceImpl) Update(foodTag *food_tag_domain.FoodTag) error {
	foodTagMgr := model.FoodTagMgr(mysql_infrastructure.Get())
	err := foodTagMgr.Save(foodTag).Error
	if err != nil {
		return err
	}
	log.Infof("update food_to_food_tag success: %s", foodTag.ID)
	return nil
}

func (f foodTagServiceImpl) List(size int64, page int64) (*food_tag_domain.ListResp, error) {
	resp := new(food_tag_domain.ListResp)
	foodTagMgr := model.FoodTagMgr(mysql_infrastructure.Get())
	foodTagPage := model.NewPage(size, page)
	result, err := foodTagMgr.SelectPage(foodTagPage, foodTagMgr.WithFlag(constant.DataNormal), foodTagMgr.WithActive(constant.DataActivated))
	if err != nil {
		return nil, err
	}
	foodTagList := make([]*food_tag_domain.FoodTag, 0)
	for _, tagRepo := range result.GetRecords().([]model.FoodTag) {
		tag := new(food_tag_domain.FoodTag)
		_ = copier.Copy(&tag, &tagRepo)
		foodTagList = append(foodTagList, tag)
	}
	resp.Items = foodTagList
	resp.Current = int(result.GetCurrent())
	resp.Size = int(result.GetSize())
	resp.Total = int(result.GetTotal())

	return resp, nil
}

func (f foodTagServiceImpl) Detail(id int64) (*food_tag_domain.FoodTag, error) {
	tagDaoList := make([]*model.FoodTag, 0)
	foodTagMgr := model.FoodTagMgr(mysql_infrastructure.Get())
	err := foodTagMgr.Where("id=?", id).Limit(1).Find(&tagDaoList).Error
	if err != nil {
		return nil, err
	}
	tag := new(food_tag_domain.FoodTag)
	_ = copier.Copy(&tag, tagDaoList[0])
	return tag, nil
}

func (f foodTagServiceImpl) Create(t *food_tag_domain.FoodTag) (string, error) {
	foodTagDao := new(model.FoodTag)
	_ = copier.Copy(foodTagDao, t)
	foodTagDao.ID = strings.Join(pinyin.LazyConvert(foodTagDao.Name, nil), "_")
	foodTagMgr := model.FoodTagMgr(mysql_infrastructure.Get())
	err := foodTagMgr.Omit("create_time", "update_time").Create(foodTagDao).Error
	if err != nil {
		return "", err
	}
	log.Infof("create food_to_food_tag success: %s", foodTagDao.ID)
	return foodTagDao.ID, nil
}
