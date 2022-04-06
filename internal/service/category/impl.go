package category_service

import (
	"can-i-eat/common/constant"
	id_util "can-i-eat/common/util/id"
	category_domain "can-i-eat/internal/domain/category"
	"can-i-eat/internal/infrastructure/model"
	mysql_infrastructure "can-i-eat/internal/infrastructure/mysql"
	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
)

var Impl CategoryService = &categoryServiceImpl{}

type categoryServiceImpl struct {
}

func (f categoryServiceImpl) Delete(category *category_domain.Category) error {
	categoryMgr := model.CategoryMgr(mysql_infrastructure.Get())
	err := categoryMgr.Update("flag", constant.Deleted).Where("id=?", category.ID).Error
	if err != nil {
		return err
	}
	log.Infof("delete food success: %d", category.ID)
	return nil
}

func (f categoryServiceImpl) Update(category *category_domain.Category) error {
	categoryMgr := model.CategoryMgr(mysql_infrastructure.Get())
	err := categoryMgr.Save(category).Error
	if err != nil {
		return err
	}
	log.Infof("update food success: %d", category.ID)
	return nil
}

func (f categoryServiceImpl) ListForPage(size int64, page int64) (*category_domain.ListResp, error) {
	resp := new(category_domain.ListResp)
	categoryMgr := model.CategoryMgr(mysql_infrastructure.Get())
	categoryPage := model.NewPage(size, page)
	result, err := categoryMgr.SelectPage(categoryPage, categoryMgr.WithFlag(constant.Normal), categoryMgr.WithActive(constant.Activated))
	if err != nil {
		return nil, err
	}
	categoryList := make([]*category_domain.Category, 0)
	for _, categoryRepo := range result.GetRecords().([]model.Category) {
		category := new(category_domain.Category)
		_ = copier.Copy(&category, &categoryRepo)
		categoryList = append(categoryList, category)
	}
	resp.Items = categoryList
	resp.Current = int(result.GetCurrent())
	resp.Size = int(result.GetSize())
	resp.Total = int(result.GetTotal())

	return resp, nil
}

func (f categoryServiceImpl) FoodDetail(id int64) (*category_domain.Category, error) {
	foodRepoList := make([]*model.Category, 0)
	categoryMgr := model.CategoryMgr(mysql_infrastructure.Get())
	err := categoryMgr.Where("id=?", id).Limit(1).Find(&foodRepoList).Error
	if err != nil {
		return nil, err
	}
	food := new(category_domain.Category)
	_ = copier.Copy(&food, foodRepoList[0])
	return food, nil
}

func (f categoryServiceImpl) Create(category *category_domain.Category) (uint64, error) {
	categoryDao := new(model.Category)
	_ = copier.Copy(categoryDao, category)
	categoryDao.ID, _ = id_util.NextID()
	categoryMgr := model.CategoryMgr(mysql_infrastructure.Get())
	err := categoryMgr.Omit("create_time", "update_time").Create(categoryDao).Error
	if err != nil {
		return 0, err
	}
	log.Infof("create food success: %d", categoryDao.ID)
	return categoryDao.ID, nil
}
