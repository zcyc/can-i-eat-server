package consumer_service

import (
	"can-i-eat/common/constant"
	id_util "can-i-eat/common/util/id"
	consumer_domain "can-i-eat/internal/domain/consumer"
	"can-i-eat/internal/infrastructure/model"
	mysql_infrastructure "can-i-eat/internal/infrastructure/mysql"
	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
)

var Impl ConsumerService = &consumerServiceImpl{}

type consumerServiceImpl struct {
}

func (f consumerServiceImpl) Delete(consumer *consumer_domain.Consumer) error {
	consumerMgr := model.ConsumerMgr(mysql_infrastructure.Get())
	err := consumerMgr.Update("flag", constant.Deleted).Error
	if err != nil {
		return err
	}
	log.Infof("delete food success: %d", consumer.ID)
	return nil
}

func (f consumerServiceImpl) Update(consumer *consumer_domain.Consumer) error {
	consumerMgr := model.ConsumerMgr(mysql_infrastructure.Get())
	err := consumerMgr.Save(consumer).Error
	if err != nil {
		return err
	}
	log.Infof("update food success: %d", consumer.ID)
	return nil
}

func (f consumerServiceImpl) ListForPage(size int64, page int64) (*consumer_domain.ListResp, error) {
	resp := new(consumer_domain.ListResp)
	consumerMgr := model.ConsumerMgr(mysql_infrastructure.Get())
	foodPage := model.NewPage(size, page)
	result, err := consumerMgr.SelectPage(foodPage, consumerMgr.WithFlag(constant.Normal), consumerMgr.WithActive(constant.Activated))
	if err != nil {
		return nil, err
	}
	consumerList := make([]*consumer_domain.Consumer, 0)
	for _, consumerRepo := range result.GetRecords().([]model.Consumer) {
		consumer := new(consumer_domain.Consumer)
		_ = copier.Copy(&consumer, &consumerRepo)
		consumerList = append(consumerList, consumer)
	}
	resp.Items = consumerList
	resp.Current = int(result.GetCurrent())
	resp.Size = int(result.GetSize())
	resp.Total = int(result.GetTotal())

	return resp, nil
}

func (f consumerServiceImpl) FoodDetail(id int64) (*consumer_domain.Consumer, error) {
	foodRepoList := make([]*model.Consumer, 0)
	consumerMgr := model.ConsumerMgr(mysql_infrastructure.Get())
	err := consumerMgr.Where("id=?", id).Limit(1).Find(&foodRepoList).Error
	if err != nil {
		return nil, err
	}
	food := new(consumer_domain.Consumer)
	_ = copier.Copy(&food, foodRepoList[0])
	return food, nil
}

func (f consumerServiceImpl) Create(consumer *consumer_domain.Consumer) (uint64, error) {
	consumerDao := new(model.Consumer)
	_ = copier.Copy(consumerDao, consumer)
	consumerDao.ID, _ = id_util.NextID()
	consumerMgr := model.ConsumerMgr(mysql_infrastructure.Get())
	err := consumerMgr.Omit("create_time", "update_time").Create(consumerDao).Error
	if err != nil {
		return 0, err
	}
	log.Infof("create food success: %d", consumerDao.ID)
	return consumerDao.ID, nil
}
