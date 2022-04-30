package consumer_service

import (
	"can-i-eat/common/constant"
	consumer_domain "can-i-eat/internal/domain/consumer"
	"can-i-eat/internal/infrastructure/model"
	mysql_infrastructure "can-i-eat/internal/infrastructure/mysql"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
)

var Impl ConsumerService = &consumerServiceImpl{}

type consumerServiceImpl struct {
}

func (f consumerServiceImpl) Delete(consumer *consumer_domain.Consumer) error {
	consumerMgr := model.ConsumerMgr(mysql_infrastructure.Get())
	err := consumerMgr.Update("flag", constant.DataDeleted).Error
	if err != nil {
		return err
	}
	log.Infof("delete consumer success: %s", consumer.ID)
	return nil
}

func (f consumerServiceImpl) Update(consumer *consumer_domain.Consumer) error {
	consumerMgr := model.ConsumerMgr(mysql_infrastructure.Get())
	err := consumerMgr.Save(consumer).Error
	if err != nil {
		return err
	}
	log.Infof("update consumer success: %s", consumer.ID)
	return nil
}

func (f consumerServiceImpl) List(size int64, page int64) (*consumer_domain.ListResp, error) {
	resp := new(consumer_domain.ListResp)
	consumerMgr := model.ConsumerMgr(mysql_infrastructure.Get())
	consumerPage := model.NewPage(size, page)
	result, err := consumerMgr.SelectPage(consumerPage, consumerMgr.WithFlag(constant.DataNormal), consumerMgr.WithActive(constant.DataActivated))
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

func (f consumerServiceImpl) Detail(id int64) (*consumer_domain.Consumer, error) {
	consumerDaoList := make([]*model.Consumer, 0)
	consumerMgr := model.ConsumerMgr(mysql_infrastructure.Get())
	err := consumerMgr.Where("id=?", id).Limit(1).Find(&consumerDaoList).Error
	if err != nil {
		return nil, err
	}
	consumer := new(consumer_domain.Consumer)
	_ = copier.Copy(&consumer, consumerDaoList[0])
	return consumer, nil
}

func (f consumerServiceImpl) Create(t *consumer_domain.Consumer) (string, error) {
	consumerDao := new(model.Consumer)
	_ = copier.Copy(consumerDao, t)
	newUUID, _ := uuid.NewV4()
	consumerDao.ID = newUUID.String()
	consumerMgr := model.ConsumerMgr(mysql_infrastructure.Get())
	err := consumerMgr.Omit("create_time", "update_time").Create(consumerDao).Error
	if err != nil {
		return "", err
	}
	log.Infof("create consumer success: %s", consumerDao.ID)
	return consumerDao.ID, nil
}
