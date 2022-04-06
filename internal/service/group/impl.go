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

func (f groupServiceImpl) ListByIDs(id []int64) ([]*group_domain.Group, error) {
	groupDaoList := make([]*model.Group, 0)
	groupMgr := model.GroupMgr(mysql_infrastructure.Get())
	err := groupMgr.Where("id in ?", id).Find(&groupDaoList).Error
	if err != nil {
		return nil, err
	}
	groupList := make([]*group_domain.Group, 0)
	for i := range groupDaoList {
		group := new(group_domain.Group)
		_ = copier.Copy(&group, &groupDaoList[i])
		groupList = append(groupList, group)
	}
	return groupList, nil
}

func (f groupServiceImpl) Delete(group *group_domain.Group) error {
	consumerGroupMgr := model.ConsumerGroupMgr(mysql_infrastructure.Get())
	err := consumerGroupMgr.Update("flag", constant.Deleted).Where("id=?", group.ID).Error
	if err != nil {
		return err
	}
	log.Infof("delete group success: %d", group.ID)
	return nil
}

func (f groupServiceImpl) Update(group *group_domain.Group) error {
	consumerGroupMgr := model.ConsumerGroupMgr(mysql_infrastructure.Get())
	err := consumerGroupMgr.Save(group).Error
	if err != nil {
		return err
	}
	log.Infof("update group success: %d", group.ID)
	return nil
}

func (f groupServiceImpl) List(size int64, page int64) (*group_domain.ListResp, error) {
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

func (f groupServiceImpl) Detail(id int64) (*group_domain.Group, error) {
	groupDaoList := make([]*model.Group, 0)
	consumerGroupMgr := model.ConsumerGroupMgr(mysql_infrastructure.Get())
	err := consumerGroupMgr.Where("id=?", id).Limit(1).Find(&groupDaoList).Error
	if err != nil {
		return nil, err
	}
	group := new(group_domain.Group)
	_ = copier.Copy(&group, groupDaoList[0])
	return group, nil
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
	log.Infof("create group success: %d", groupDao.ID)
	return groupDao.ID, nil
}
