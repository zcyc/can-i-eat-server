package group_service

import (
	"can-i-eat/common/constant"
	id_util "can-i-eat/common/util/id"
	group_domain "can-i-eat/internal/domain/group"
	"can-i-eat/internal/infrastructure/model"
	mysql_infrastructure "can-i-eat/internal/infrastructure/mysql"
	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
)

var Impl GroupService = &groupServiceImpl{}

type groupServiceImpl struct {
}

func (f groupServiceImpl) Delete(group *group_domain.Group) error {
	consumerGroupMgr := model.ConsumerGroupMgr(mysql_infrastructure.Get())
	err := consumerGroupMgr.Update("flag", constant.Deleted).Where("id=?", group.ID).Error
	if err != nil {
		return err
	}
	log.Infof("delete food success: %d", group.ID)
	return nil
}

func (f groupServiceImpl) Update(group *group_domain.Group) error {
	consumerGroupMgr := model.ConsumerGroupMgr(mysql_infrastructure.Get())
	err := consumerGroupMgr.Save(group).Error
	if err != nil {
		return err
	}
	log.Infof("update food success: %d", group.ID)
	return nil
}

func (f groupServiceImpl) ListForPage(size int64, page int64) (*group_domain.ListResp, error) {
	resp := new(group_domain.ListResp)
	consumerGroupMgr := model.ConsumerGroupMgr(mysql_infrastructure.Get())
	groupPage := model.NewPage(size, page)
	result, err := consumerGroupMgr.SelectPage(groupPage, consumerGroupMgr.WithFlag(constant.Normal), consumerGroupMgr.WithActive(constant.Activated))
	if err != nil {
		return nil, err
	}
	groupList := make([]*group_domain.Group, 0)
	for _, groupRepo := range result.GetRecords().([]model.Group) {
		group := new(group_domain.Group)
		_ = copier.Copy(&group, &groupRepo)
		groupList = append(groupList, group)
	}
	resp.Items = groupList
	resp.Current = int(result.GetCurrent())
	resp.Size = int(result.GetSize())
	resp.Total = int(result.GetTotal())

	return resp, nil
}

func (f groupServiceImpl) FoodDetail(id int64) (*group_domain.Group, error) {
	foodRepoList := make([]*model.Group, 0)
	consumerGroupMgr := model.ConsumerGroupMgr(mysql_infrastructure.Get())
	err := consumerGroupMgr.Where("id=?", id).Limit(1).Find(&foodRepoList).Error
	if err != nil {
		return nil, err
	}
	food := new(group_domain.Group)
	_ = copier.Copy(&food, foodRepoList[0])
	return food, nil
}

func (f groupServiceImpl) Create(group *group_domain.Group) (uint64, error) {
	groupDao := new(model.Group)
	_ = copier.Copy(groupDao, group)
	groupDao.ID, _ = id_util.NextID()
	consumerGroupMgr := model.ConsumerGroupMgr(mysql_infrastructure.Get())
	err := consumerGroupMgr.Omit("create_time", "update_time").Create(groupDao).Error
	if err != nil {
		return 0, err
	}
	log.Infof("create food success: %d", groupDao.ID)
	return groupDao.ID, nil
}
