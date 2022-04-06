package consumer_group_service

import (
	"can-i-eat/common/constant"
	consumer_group_domain "can-i-eat/internal/domain/consumer_group"
	"can-i-eat/internal/infrastructure/model"
	mysql_infrastructure "can-i-eat/internal/infrastructure/mysql"
	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
)

var Impl ConsumerGroupService = &consumerGroupServiceImpl{}

type consumerGroupServiceImpl struct {
}

func (f consumerGroupServiceImpl) ListByConsumer(id string) ([]*consumer_group_domain.ConsumerGroup, error) {
	consumerGroupDaoList := make([]*model.ConsumerGroup, 0)
	consumerGroupMgr := model.ConsumerGroupMgr(mysql_infrastructure.Get())
	err := consumerGroupMgr.Where("consumer_id = ?", id).Find(&consumerGroupDaoList).Error
	if err != nil {
		return nil, err
	}
	consumerGroupList := make([]*consumer_group_domain.ConsumerGroup, 0)
	for i := range consumerGroupDaoList {
		consumerGroup := new(consumer_group_domain.ConsumerGroup)
		_ = copier.Copy(consumerGroup, consumerGroupDaoList[i])
		consumerGroupList = append(consumerGroupList, consumerGroup)
	}
	log.Infof("ListByConsumer consumerGroup success: %s", len(consumerGroupList))
	return consumerGroupList, nil
}

func (f consumerGroupServiceImpl) Delete(consumerGroup *consumer_group_domain.ConsumerGroup) error {
	consumerGroupMgr := model.ConsumerGroupMgr(mysql_infrastructure.Get())
	err := consumerGroupMgr.Update("flag", constant.Deleted).Where("id=?", consumerGroup.ID).Error
	if err != nil {
		return err
	}
	log.Infof("delete consumerGroup success: %s", consumerGroup.ID)
	return nil
}

func (f consumerGroupServiceImpl) Update(consumerGroup *consumer_group_domain.ConsumerGroup) error {
	consumerGroupMgr := model.ConsumerGroupMgr(mysql_infrastructure.Get())
	err := consumerGroupMgr.Save(consumerGroup).Error
	if err != nil {
		return err
	}
	log.Infof("update consumerGroup success: %s", consumerGroup.ID)
	return nil
}

func (f consumerGroupServiceImpl) List(size int64, page int64) (*consumer_group_domain.ListResp, error) {
	resp := new(consumer_group_domain.ListResp)
	consumerGroupMgr := model.ConsumerGroupMgr(mysql_infrastructure.Get())
	consumerGroupPage := model.NewPage(size, page)
	result, err := consumerGroupMgr.SelectPage(consumerGroupPage, consumerGroupMgr.WithFlag(constant.Normal), consumerGroupMgr.WithActive(constant.Activated))
	if err != nil {
		return nil, err
	}
	consumerGroupList := make([]*consumer_group_domain.ConsumerGroup, 0)
	for _, consumerGroupRepo := range result.GetRecords().([]model.ConsumerGroup) {
		consumerGroup := new(consumer_group_domain.ConsumerGroup)
		_ = copier.Copy(&consumerGroup, &consumerGroupRepo)
		consumerGroupList = append(consumerGroupList, consumerGroup)
	}
	resp.Items = consumerGroupList
	resp.Current = int(result.GetCurrent())
	resp.Size = int(result.GetSize())
	resp.Total = int(result.GetTotal())

	return resp, nil
}

func (f consumerGroupServiceImpl) Detail(id string) (*consumer_group_domain.ConsumerGroup, error) {
	consumerGroupDaoList := make([]*model.ConsumerGroup, 0)
	consumerGroupMgr := model.ConsumerGroupMgr(mysql_infrastructure.Get())
	err := consumerGroupMgr.Where("id=?", id).Limit(1).Find(&consumerGroupDaoList).Error
	if err != nil {
		return nil, err
	}
	consumerGroup := new(consumer_group_domain.ConsumerGroup)
	_ = copier.Copy(&consumerGroup, consumerGroupDaoList[0])
	return consumerGroup, nil
}

func (f consumerGroupServiceImpl) Create(t *consumer_group_domain.ConsumerGroup) (string, error) {
	consumerGroupDao := new(model.ConsumerGroup)
	_ = copier.Copy(consumerGroupDao, t)
	consumerGroupDao.ID = consumerGroupDao.ConsumerID + "_" + consumerGroupDao.GroupID
	consumerGroupMgr := model.ConsumerGroupMgr(mysql_infrastructure.Get())
	err := consumerGroupMgr.Omit("create_time", "update_time").Create(consumerGroupDao).Error
	if err != nil {
		return "", err
	}
	log.Infof("create consumerGroup success: %s", consumerGroupDao.ID)
	return consumerGroupDao.ID, nil
}
