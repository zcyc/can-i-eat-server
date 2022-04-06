package tag_service

import (
	"can-i-eat/common/constant"
	id_util "can-i-eat/common/util/id"
	tag_domain "can-i-eat/internal/domain/tag"
	"can-i-eat/internal/infrastructure/model"
	mysql_infrastructure "can-i-eat/internal/infrastructure/mysql"
	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
)

var Impl TagService = &tagServiceImpl{}

type tagServiceImpl struct {
}

func (f tagServiceImpl) Delete(tag *tag_domain.Tag) error {
	tagMgr := model.TagMgr(mysql_infrastructure.Get())
	err := tagMgr.Update("flag", constant.Deleted).Where("id=?", tag.ID).Error
	if err != nil {
		return err
	}
	log.Infof("delete tag success: %d", tag.ID)
	return nil
}

func (f tagServiceImpl) Update(tag *tag_domain.Tag) error {
	tagMgr := model.TagMgr(mysql_infrastructure.Get())
	err := tagMgr.Save(tag).Error
	if err != nil {
		return err
	}
	log.Infof("update tag success: %d", tag.ID)
	return nil
}

func (f tagServiceImpl) List(size int64, page int64) (*tag_domain.ListResp, error) {
	resp := new(tag_domain.ListResp)
	tagMgr := model.TagMgr(mysql_infrastructure.Get())
	tagPage := model.NewPage(size, page)
	result, err := tagMgr.SelectPage(tagPage, tagMgr.WithFlag(constant.Normal), tagMgr.WithActive(constant.Activated))
	if err != nil {
		return nil, err
	}
	tagList := make([]*tag_domain.Tag, 0)
	for _, tagRepo := range result.GetRecords().([]model.Tag) {
		tag := new(tag_domain.Tag)
		_ = copier.Copy(&tag, &tagRepo)
		tagList = append(tagList, tag)
	}
	resp.Items = tagList
	resp.Current = int(result.GetCurrent())
	resp.Size = int(result.GetSize())
	resp.Total = int(result.GetTotal())

	return resp, nil
}

func (f tagServiceImpl) Detail(id int64) (*tag_domain.Tag, error) {
	tagDaoList := make([]*model.Tag, 0)
	tagMgr := model.TagMgr(mysql_infrastructure.Get())
	err := tagMgr.Where("id=?", id).Limit(1).Find(&tagDaoList).Error
	if err != nil {
		return nil, err
	}
	tag := new(tag_domain.Tag)
	_ = copier.Copy(&tag, tagDaoList[0])
	return tag, nil
}

func (f tagServiceImpl) Create(tag *tag_domain.Tag) (uint64, error) {
	tagDao := new(model.Tag)
	_ = copier.Copy(tagDao, tag)
	tagDao.ID, _ = id_util.NextID()
	tagMgr := model.TagMgr(mysql_infrastructure.Get())
	err := tagMgr.Omit("create_time", "update_time").Create(tagDao).Error
	if err != nil {
		return 0, err
	}
	log.Infof("create tag success: %d", tagDao.ID)
	return tagDao.ID, nil
}
