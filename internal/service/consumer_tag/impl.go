package consumer_tag_service

import (
	"can-i-eat/common/constant"
	consumer_tag_domain "can-i-eat/internal/domain/consumer_tag"
	"can-i-eat/internal/infrastructure/model"
	mysql_infrastructure "can-i-eat/internal/infrastructure/mysql"
	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
	"github.com/mozillazg/go-pinyin"
	"gorm.io/gorm/clause"
	"strings"
)

var Impl ConsumerTagService = &consumerTagServiceImpl{}

type consumerTagServiceImpl struct {
}

func (f consumerTagServiceImpl) BatchCreate(t []*consumer_tag_domain.ConsumerTag) error {
	// domain 对象转 repo 对象
	consumerTagDaoList := make([]*model.ConsumerTag, 0)
	for _, consumerTag := range t {
		consumerTagDao := new(model.ConsumerTag)
		_ = copier.Copy(consumerTagDao, consumerTag)
		consumerTagDaoList = append(consumerTagDaoList, consumerTagDao)
	}

	// 执行批量
	consumerTagMgr := model.ConsumerTagMgr(mysql_infrastructure.Get())
	err := consumerTagMgr.Omit("create_time", "update_time").Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&consumerTagDaoList, 100).Error
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (f consumerTagServiceImpl) ListByIDs(id []int64) ([]*consumer_tag_domain.ConsumerTag, error) {
	consumerTagDaoList := make([]*model.ConsumerTag, 0)
	consumerTagMgr := model.ConsumerTagMgr(mysql_infrastructure.Get())
	err := consumerTagMgr.Where("id in ?", id).Find(&consumerTagDaoList).Error
	if err != nil {
		return nil, err
	}
	consumerTagList := make([]*consumer_tag_domain.ConsumerTag, 0)
	for i := range consumerTagDaoList {
		consumerTag := new(consumer_tag_domain.ConsumerTag)
		_ = copier.Copy(&consumerTag, &consumerTagDaoList[i])
		consumerTagList = append(consumerTagList, consumerTag)
	}
	return consumerTagList, nil
}

func (f consumerTagServiceImpl) Delete(consumerTag *consumer_tag_domain.ConsumerTag) error {
	consumerTagMgr := model.ConsumerTagMgr(mysql_infrastructure.Get())
	err := consumerTagMgr.Update("flag", constant.Deleted).Where("id=?", consumerTag.ID).Error
	if err != nil {
		return err
	}
	log.Infof("delete consumerTag success: %s", consumerTag.ID)
	return nil
}

func (f consumerTagServiceImpl) Update(consumerTag *consumer_tag_domain.ConsumerTag) error {
	consumerTagMgr := model.ConsumerTagMgr(mysql_infrastructure.Get())
	err := consumerTagMgr.Save(consumerTag).Error
	if err != nil {
		return err
	}
	log.Infof("update consumerTag success: %s", consumerTag.ID)
	return nil
}

func (f consumerTagServiceImpl) List(size int64, page int64) (*consumer_tag_domain.ListResp, error) {
	resp := new(consumer_tag_domain.ListResp)
	consumerTagMgr := model.ConsumerTagMgr(mysql_infrastructure.Get())
	consumerTagPage := model.NewPage(size, page)
	result, err := consumerTagMgr.SelectPage(consumerTagPage, consumerTagMgr.WithFlag(constant.Normal), consumerTagMgr.WithActive(constant.Activated))
	if err != nil {
		return nil, err
	}
	consumerTagList := make([]*consumer_tag_domain.ConsumerTag, 0)
	for _, consumerTagRepo := range result.GetRecords().([]model.ConsumerTag) {
		consumerTag := new(consumer_tag_domain.ConsumerTag)
		_ = copier.Copy(&consumerTag, &consumerTagRepo)
		consumerTagList = append(consumerTagList, consumerTag)
	}
	resp.Items = consumerTagList
	resp.Current = int(result.GetCurrent())
	resp.Size = int(result.GetSize())
	resp.Total = int(result.GetTotal())

	return resp, nil
}

func (f consumerTagServiceImpl) Detail(id int64) (*consumer_tag_domain.ConsumerTag, error) {
	consumerTagDaoList := make([]*model.ConsumerTag, 0)
	consumerTagMgr := model.ConsumerTagMgr(mysql_infrastructure.Get())
	err := consumerTagMgr.Where("id=?", id).Limit(1).Find(&consumerTagDaoList).Error
	if err != nil {
		return nil, err
	}
	consumerTag := new(consumer_tag_domain.ConsumerTag)
	_ = copier.Copy(&consumerTag, consumerTagDaoList[0])
	return consumerTag, nil
}

func (f consumerTagServiceImpl) Create(t *consumer_tag_domain.ConsumerTag) (string, error) {
	consumerTagDao := new(model.ConsumerTag)
	_ = copier.Copy(consumerTagDao, t)
	consumerTagDao.ID = strings.Join(pinyin.LazyConvert(consumerTagDao.Name, nil), "_")
	consumerTagMgr := model.ConsumerTagMgr(mysql_infrastructure.Get())
	err := consumerTagMgr.Omit("create_time", "update_time").Create(consumerTagDao).Error
	if err != nil {
		return "", err
	}
	log.Infof("create consumerTag success: %s", consumerTagDao.ID)
	return consumerTagDao.ID, nil
}
