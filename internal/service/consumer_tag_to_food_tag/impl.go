package consumer_tag_to_food_tag_service

import (
	"can-i-eat/common/constant"
	consumer_tag_to_food_tag_domain "can-i-eat/internal/domain/consumer_tag_to_food_tag"
	"can-i-eat/internal/infrastructure/model"
	mysql_infrastructure "can-i-eat/internal/infrastructure/mysql"
	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
)

var Impl ConsumerTagToFoodTagService = &consumerTagToFoodTagServiceImpl{}

type consumerTagToFoodTagServiceImpl struct {
}

func (f consumerTagToFoodTagServiceImpl) ListByConsumerTagIDs(ids []string) ([]*consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag, error) {
	consumerTagToFoodTagDaoList := make([]*model.ConsumerTagToFoodTag, 0)
	consumerTagToFoodTagMgr := model.ConsumerTagToFoodTagMgr(mysql_infrastructure.Get())
	err := consumerTagToFoodTagMgr.Where("consumerTag_id in ?", ids).Find(&consumerTagToFoodTagDaoList).Error
	if err != nil {
		return nil, err
	}
	consumerTagToFoodTagList := make([]*consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag, 0)
	for _, consumerTagToFoodTagRepo := range consumerTagToFoodTagDaoList {
		consumerTagToFoodTag := new(consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag)
		_ = copier.Copy(&consumerTagToFoodTag, &consumerTagToFoodTagRepo)
		consumerTagToFoodTagList = append(consumerTagToFoodTagList, consumerTagToFoodTag)
	}

	return consumerTagToFoodTagList, nil
}

func (f consumerTagToFoodTagServiceImpl) ListByConsumerTag(id string) ([]*consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag, error) {
	consumerTagToFoodTagDaoList := make([]*model.ConsumerTagToFoodTag, 0)
	consumerTagToFoodTagMgr := model.ConsumerTagToFoodTagMgr(mysql_infrastructure.Get())
	err := consumerTagToFoodTagMgr.Where("consumerTag_id = ?", id).Find(&consumerTagToFoodTagDaoList).Error
	if err != nil {
		return nil, err
	}
	consumerTagToFoodTagList := make([]*consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag, 0)
	for _, consumerTagToFoodTagRepo := range consumerTagToFoodTagDaoList {
		consumerTagToFoodTag := new(consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag)
		_ = copier.Copy(&consumerTagToFoodTag, &consumerTagToFoodTagRepo)
		consumerTagToFoodTagList = append(consumerTagToFoodTagList, consumerTagToFoodTag)
	}

	return consumerTagToFoodTagList, nil
}

func (f consumerTagToFoodTagServiceImpl) Delete(consumerTagToFoodTag *consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag) error {
	consumerTagToFoodTagMgr := model.ConsumerTagToFoodTagMgr(mysql_infrastructure.Get())
	err := consumerTagToFoodTagMgr.Update("flag", constant.Deleted).Where("id=?", consumerTagToFoodTag.ID).Error
	if err != nil {
		return err
	}
	log.Infof("delete consumerTagToFoodTag success: %s", consumerTagToFoodTag.ID)
	return nil
}

func (f consumerTagToFoodTagServiceImpl) Update(consumerTagToFoodTag *consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag) error {
	consumerTagToFoodTagMgr := model.ConsumerTagToFoodTagMgr(mysql_infrastructure.Get())
	err := consumerTagToFoodTagMgr.Save(consumerTagToFoodTag).Error
	if err != nil {
		return err
	}
	log.Infof("update consumerTagToFoodTag success: %s", consumerTagToFoodTag.ID)
	return nil
}

func (f consumerTagToFoodTagServiceImpl) List(size int64, page int64) (*consumer_tag_to_food_tag_domain.ListResp, error) {
	resp := new(consumer_tag_to_food_tag_domain.ListResp)
	consumerTagToFoodTagMgr := model.ConsumerTagToFoodTagMgr(mysql_infrastructure.Get())
	consumerTagPage := model.NewPage(size, page)
	result, err := consumerTagToFoodTagMgr.SelectPage(consumerTagPage, consumerTagToFoodTagMgr.WithFlag(constant.Normal), consumerTagToFoodTagMgr.WithActive(constant.Activated))
	if err != nil {
		return nil, err
	}
	consumerTagToFoodTagList := make([]*consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag, 0)
	for _, consumerTagToFoodTagRepo := range result.GetRecords().([]model.ConsumerTagToFoodTag) {
		consumerTagToFoodTag := new(consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag)
		_ = copier.Copy(&consumerTagToFoodTag, &consumerTagToFoodTagRepo)
		consumerTagToFoodTagList = append(consumerTagToFoodTagList, consumerTagToFoodTag)
	}
	resp.Items = consumerTagToFoodTagList
	resp.Current = int(result.GetCurrent())
	resp.Size = int(result.GetSize())
	resp.Total = int(result.GetTotal())

	return resp, nil
}

func (f consumerTagToFoodTagServiceImpl) Detail(id string) (*consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag, error) {
	consumerTagToFoodTagDaoList := make([]*model.ConsumerTagToFoodTag, 0)
	consumerTagToFoodTagMgr := model.ConsumerTagToFoodTagMgr(mysql_infrastructure.Get())
	err := consumerTagToFoodTagMgr.Where("id=?", id).Limit(1).Find(&consumerTagToFoodTagDaoList).Error
	if err != nil {
		return nil, err
	}
	consumerTagToFoodTag := new(consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag)
	_ = copier.Copy(&consumerTagToFoodTag, consumerTagToFoodTagDaoList[0])
	return consumerTagToFoodTag, nil
}

func (f consumerTagToFoodTagServiceImpl) Create(t *consumer_tag_to_food_tag_domain.ConsumerTagToFoodTag) (string, error) {
	consumerTagToFoodTagDao := new(model.ConsumerTagToFoodTag)
	_ = copier.Copy(consumerTagToFoodTagDao, t)
	consumerTagToFoodTagDao.ID = consumerTagToFoodTagDao.ConsumerTagID + "_" + consumerTagToFoodTagDao.FoodTagID
	consumerTagToFoodTagMgr := model.ConsumerTagToFoodTagMgr(mysql_infrastructure.Get())
	err := consumerTagToFoodTagMgr.Omit("create_time", "update_time").Create(consumerTagToFoodTagDao).Error
	if err != nil {
		return "", err
	}
	log.Infof("create consumerTagToFoodTag success: %s", consumerTagToFoodTagDao.ID)
	return consumerTagToFoodTagDao.ID, nil
}