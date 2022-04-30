package eat_mode_service

import (
	"can-i-eat/common/constant"
	eat_mode_domain "can-i-eat/internal/domain/eat_mode"
	"can-i-eat/internal/infrastructure/model"
	mysql_infrastructure "can-i-eat/internal/infrastructure/mysql"
	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
	"github.com/mozillazg/go-pinyin"
	"gorm.io/gorm/clause"
	"strings"
)

var Impl EatModeService = &eatModeServiceImpl{}

type eatModeServiceImpl struct {
}

func (f eatModeServiceImpl) BatchCreate(t []*eat_mode_domain.EatMode) error {
	// domain 对象转 repo 对象
	eatModeDaoList := make([]*model.EatMode, 0)
	for _, eatMode := range t {
		eatModeDao := new(model.EatMode)
		_ = copier.Copy(eatModeDao, eatMode)
		eatModeDaoList = append(eatModeDaoList, eatModeDao)
	}

	// 执行批量
	eatModeMgr := model.EatModeMgr(mysql_infrastructure.Get())
	err := eatModeMgr.Omit("create_time", "update_time").Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&eatModeDaoList, 100).Error
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (f eatModeServiceImpl) ListByIDs(ids []string) ([]*eat_mode_domain.EatMode, error) {
	eatModeDaoList := make([]*model.EatMode, 0)
	eatModeMgr := model.EatModeMgr(mysql_infrastructure.Get())
	err := eatModeMgr.Where("id in ?", ids).Find(&eatModeDaoList).Error
	if err != nil {
		return nil, err
	}
	eatModeList := make([]*eat_mode_domain.EatMode, 0)
	for _, eatModeRepo := range eatModeDaoList {
		eatMode := new(eat_mode_domain.EatMode)
		_ = copier.Copy(&eatMode, &eatModeRepo)
		eatModeList = append(eatModeList, eatMode)
	}

	return eatModeList, nil
}

func (f eatModeServiceImpl) Delete(eat_mode *eat_mode_domain.EatMode) error {
	eatModeMgr := model.EatModeMgr(mysql_infrastructure.Get())
	err := eatModeMgr.Update("flag", constant.DataDeleted).Where("id=?", eat_mode.ID).Error
	if err != nil {
		return err
	}
	log.Infof("delete eat_mode success: %s", eat_mode.ID)
	return nil
}

func (f eatModeServiceImpl) Update(eat_mode *eat_mode_domain.EatMode) error {
	eatModeMgr := model.EatModeMgr(mysql_infrastructure.Get())
	err := eatModeMgr.Save(eat_mode).Error
	if err != nil {
		return err
	}
	log.Infof("update eat_mode success: %s", eat_mode.ID)
	return nil
}

func (f eatModeServiceImpl) List(size int64, page int64) (*eat_mode_domain.ListResp, error) {
	resp := new(eat_mode_domain.ListResp)
	eatModeMgr := model.EatModeMgr(mysql_infrastructure.Get())
	eatModePage := model.NewPage(size, page)
	result, err := eatModeMgr.SelectPage(eatModePage, eatModeMgr.WithFlag(constant.DataNormal), eatModeMgr.WithActive(constant.DataActivated))
	if err != nil {
		return nil, err
	}
	eatModeList := make([]*eat_mode_domain.EatMode, 0)
	for _, eatModeRepo := range result.GetRecords().([]model.EatMode) {
		eatMode := new(eat_mode_domain.EatMode)
		_ = copier.Copy(&eatMode, &eatModeRepo)
		eatModeList = append(eatModeList, eatMode)
	}
	resp.Items = eatModeList
	resp.Current = int(result.GetCurrent())
	resp.Size = int(result.GetSize())
	resp.Total = int(result.GetTotal())

	return resp, nil
}

func (f eatModeServiceImpl) Detail(id string) (*eat_mode_domain.EatMode, error) {
	eatModeRepoList := make([]*model.EatMode, 0)
	eatModeMgr := model.EatModeMgr(mysql_infrastructure.Get())
	err := eatModeMgr.Where("id=?", id).Limit(1).Find(&eatModeRepoList).Error
	if err != nil {
		return nil, err
	}
	eatMode := new(eat_mode_domain.EatMode)
	_ = copier.Copy(&eatMode, eatModeRepoList[0])
	return eatMode, nil
}

func (f eatModeServiceImpl) Create(t *eat_mode_domain.EatMode) (string, error) {
	eatModeDao := new(model.EatMode)
	_ = copier.Copy(eatModeDao, t)
	eatModeDao.ID = strings.Join(pinyin.LazyConvert(eatModeDao.Name, nil), "_")
	eatModeMgr := model.EatModeMgr(mysql_infrastructure.Get())
	err := eatModeMgr.Omit("create_time", "update_time").Create(eatModeDao).Error
	if err != nil {
		return "", err
	}
	log.Infof("create eat_mode success: %s", eatModeDao.ID)
	return eatModeDao.ID, nil
}
