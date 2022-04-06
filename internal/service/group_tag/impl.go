package group_tag_service

import (
	"can-i-eat/common/constant"
	"can-i-eat/internal/domain/group_tag"
	"can-i-eat/internal/infrastructure/model"
	mysql_infrastructure "can-i-eat/internal/infrastructure/mysql"
	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
)

var Impl GroupTagService = &groupTagServiceImpl{}

type groupTagServiceImpl struct {
}

func (f groupTagServiceImpl) ListByGroupIDs(ids []string) ([]*group_tag_domain.GroupTag, error) {
	groupTagDaoList := make([]*model.GroupTag, 0)
	groupTagMgr := model.GroupTagMgr(mysql_infrastructure.Get())
	err := groupTagMgr.Where("group_id in ?", ids).Find(&groupTagDaoList).Error
	if err != nil {
		return nil, err
	}
	groupTagList := make([]*group_tag_domain.GroupTag, 0)
	for _, groupTagRepo := range groupTagDaoList {
		groupTag := new(group_tag_domain.GroupTag)
		_ = copier.Copy(&groupTag, &groupTagRepo)
		groupTagList = append(groupTagList, groupTag)
	}

	return groupTagList, nil
}

func (f groupTagServiceImpl) ListByGroup(id string) ([]*group_tag_domain.GroupTag, error) {
	groupTagDaoList := make([]*model.GroupTag, 0)
	groupTagMgr := model.GroupTagMgr(mysql_infrastructure.Get())
	err := groupTagMgr.Where("group_id = ?", id).Find(&groupTagDaoList).Error
	if err != nil {
		return nil, err
	}
	groupTagList := make([]*group_tag_domain.GroupTag, 0)
	for _, groupTagRepo := range groupTagDaoList {
		groupTag := new(group_tag_domain.GroupTag)
		_ = copier.Copy(&groupTag, &groupTagRepo)
		groupTagList = append(groupTagList, groupTag)
	}

	return groupTagList, nil
}

func (f groupTagServiceImpl) Delete(groupTag *group_tag_domain.GroupTag) error {
	groupTagMgr := model.GroupTagMgr(mysql_infrastructure.Get())
	err := groupTagMgr.Update("flag", constant.Deleted).Where("id=?", groupTag.ID).Error
	if err != nil {
		return err
	}
	log.Infof("delete groupTag success: %s", groupTag.ID)
	return nil
}

func (f groupTagServiceImpl) Update(groupTag *group_tag_domain.GroupTag) error {
	groupTagMgr := model.GroupTagMgr(mysql_infrastructure.Get())
	err := groupTagMgr.Save(groupTag).Error
	if err != nil {
		return err
	}
	log.Infof("update groupTag success: %s", groupTag.ID)
	return nil
}

func (f groupTagServiceImpl) List(size int64, page int64) (*group_tag_domain.ListResp, error) {
	resp := new(group_tag_domain.ListResp)
	groupTagMgr := model.GroupTagMgr(mysql_infrastructure.Get())
	groupPage := model.NewPage(size, page)
	result, err := groupTagMgr.SelectPage(groupPage, groupTagMgr.WithFlag(constant.Normal), groupTagMgr.WithActive(constant.Activated))
	if err != nil {
		return nil, err
	}
	groupTagList := make([]*group_tag_domain.GroupTag, 0)
	for _, groupTagRepo := range result.GetRecords().([]model.GroupTag) {
		groupTag := new(group_tag_domain.GroupTag)
		_ = copier.Copy(&groupTag, &groupTagRepo)
		groupTagList = append(groupTagList, groupTag)
	}
	resp.Items = groupTagList
	resp.Current = int(result.GetCurrent())
	resp.Size = int(result.GetSize())
	resp.Total = int(result.GetTotal())

	return resp, nil
}

func (f groupTagServiceImpl) Detail(id string) (*group_tag_domain.GroupTag, error) {
	groupTagDaoList := make([]*model.GroupTag, 0)
	groupTagMgr := model.GroupTagMgr(mysql_infrastructure.Get())
	err := groupTagMgr.Where("id=?", id).Limit(1).Find(&groupTagDaoList).Error
	if err != nil {
		return nil, err
	}
	groupTag := new(group_tag_domain.GroupTag)
	_ = copier.Copy(&groupTag, groupTagDaoList[0])
	return groupTag, nil
}

func (f groupTagServiceImpl) Create(t *group_tag_domain.GroupTag) (string, error) {
	groupTagDao := new(model.GroupTag)
	_ = copier.Copy(groupTagDao, t)
	groupTagDao.ID = groupTagDao.GroupID + "_" + groupTagDao.TagID
	groupTagMgr := model.GroupTagMgr(mysql_infrastructure.Get())
	err := groupTagMgr.Omit("create_time", "update_time").Create(groupTagDao).Error
	if err != nil {
		return "", err
	}
	log.Infof("create groupTag success: %s", groupTagDao.ID)
	return groupTagDao.ID, nil
}
