package consumer_to_consumer_tag_service

import (
	"can-i-eat/common/constant"
	consumer_to_consumer_tag_domain "can-i-eat/internal/domain/consumer_to_consumer_tag"
	"can-i-eat/internal/infrastructure/model"
	mysql_infrastructure "can-i-eat/internal/infrastructure/mysql"
	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
)

var Impl ConsumerToConsumerTagService = &consumerToConsumerTagServiceImpl{}

type consumerToConsumerTagServiceImpl struct {
}

func (f consumerToConsumerTagServiceImpl) ListByConsumer(id string) ([]*consumer_to_consumer_tag_domain.ConsumerToConsumerTag, error) {
	consumerToConsumerTagDaoList := make([]*model.ConsumerToConsumerTag, 0)
	consumerToConsumerTagMgr := model.ConsumerToConsumerTagMgr(mysql_infrastructure.Get())
	err := consumerToConsumerTagMgr.Where("consumer_id = ?", id).Find(&consumerToConsumerTagDaoList).Error
	if err != nil {
		return nil, err
	}
	consumerToConsumerTagList := make([]*consumer_to_consumer_tag_domain.ConsumerToConsumerTag, 0)
	for i := range consumerToConsumerTagDaoList {
		consumerToConsumerTag := new(consumer_to_consumer_tag_domain.ConsumerToConsumerTag)
		_ = copier.Copy(consumerToConsumerTag, consumerToConsumerTagDaoList[i])
		consumerToConsumerTagList = append(consumerToConsumerTagList, consumerToConsumerTag)
	}
	log.Infof("ListByConsumer consumerToConsumerTag success: %s", len(consumerToConsumerTagList))
	return consumerToConsumerTagList, nil
}

func (f consumerToConsumerTagServiceImpl) Delete(consumerToConsumerTag *consumer_to_consumer_tag_domain.ConsumerToConsumerTag) error {
	consumerToConsumerTagMgr := model.ConsumerToConsumerTagMgr(mysql_infrastructure.Get())
	err := consumerToConsumerTagMgr.Update("flag", constant.DataDeleted).Where("id=?", consumerToConsumerTag.ID).Error
	if err != nil {
		return err
	}
	log.Infof("delete consumerToConsumerTag success: %s", consumerToConsumerTag.ID)
	return nil
}

func (f consumerToConsumerTagServiceImpl) Update(consumerToConsumerTag *consumer_to_consumer_tag_domain.ConsumerToConsumerTag) error {
	consumerToConsumerTagMgr := model.ConsumerToConsumerTagMgr(mysql_infrastructure.Get())
	err := consumerToConsumerTagMgr.Save(consumerToConsumerTag).Error
	if err != nil {
		return err
	}
	log.Infof("update consumerToConsumerTag success: %s", consumerToConsumerTag.ID)
	return nil
}

func (f consumerToConsumerTagServiceImpl) List(size int64, page int64) (*consumer_to_consumer_tag_domain.ListResp, error) {
	resp := new(consumer_to_consumer_tag_domain.ListResp)
	consumerToConsumerTagMgr := model.ConsumerToConsumerTagMgr(mysql_infrastructure.Get())
	consumerToConsumerTagPage := model.NewPage(size, page)
	result, err := consumerToConsumerTagMgr.SelectPage(consumerToConsumerTagPage, consumerToConsumerTagMgr.WithFlag(constant.DataNormal), consumerToConsumerTagMgr.WithActive(constant.DataActivated))
	if err != nil {
		return nil, err
	}
	consumerToConsumerTagList := make([]*consumer_to_consumer_tag_domain.ConsumerToConsumerTag, 0)
	for _, consumerToConsumerTagRepo := range result.GetRecords().([]model.ConsumerToConsumerTag) {
		consumerToConsumerTag := new(consumer_to_consumer_tag_domain.ConsumerToConsumerTag)
		_ = copier.Copy(&consumerToConsumerTag, &consumerToConsumerTagRepo)
		consumerToConsumerTagList = append(consumerToConsumerTagList, consumerToConsumerTag)
	}
	resp.Items = consumerToConsumerTagList
	resp.Current = int(result.GetCurrent())
	resp.Size = int(result.GetSize())
	resp.Total = int(result.GetTotal())

	return resp, nil
}

func (f consumerToConsumerTagServiceImpl) Detail(id string) (*consumer_to_consumer_tag_domain.ConsumerToConsumerTag, error) {
	consumerToConsumerTagDaoList := make([]*model.ConsumerToConsumerTag, 0)
	consumerToConsumerTagMgr := model.ConsumerToConsumerTagMgr(mysql_infrastructure.Get())
	err := consumerToConsumerTagMgr.Where("id=?", id).Limit(1).Find(&consumerToConsumerTagDaoList).Error
	if err != nil {
		return nil, err
	}
	consumerToConsumerTag := new(consumer_to_consumer_tag_domain.ConsumerToConsumerTag)
	_ = copier.Copy(&consumerToConsumerTag, consumerToConsumerTagDaoList[0])
	return consumerToConsumerTag, nil
}

func (f consumerToConsumerTagServiceImpl) Create(t *consumer_to_consumer_tag_domain.ConsumerToConsumerTag) (string, error) {
	consumerToConsumerTagDao := new(model.ConsumerToConsumerTag)
	_ = copier.Copy(consumerToConsumerTagDao, t)
	consumerToConsumerTagDao.ID = consumerToConsumerTagDao.ConsumerID + "_" + consumerToConsumerTagDao.ConsumerTagID
	consumerToConsumerTagMgr := model.ConsumerToConsumerTagMgr(mysql_infrastructure.Get())
	err := consumerToConsumerTagMgr.Omit("create_time", "update_time").Create(consumerToConsumerTagDao).Error
	if err != nil {
		return "", err
	}
	log.Infof("create consumerToConsumerTag success: %s", consumerToConsumerTagDao.ID)
	return consumerToConsumerTagDao.ID, nil
}
