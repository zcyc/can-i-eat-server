package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _EatModeMgr struct {
	*_BaseMgr
}

// EatModeMgr open func
func EatModeMgr(db *gorm.DB) *_EatModeMgr {
	if db == nil {
		panic(fmt.Errorf("EatModeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EatModeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eat_mode"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EatModeMgr) GetTableName() string {
	return "eat_mode"
}

// Reset 重置gorm会话
func (obj *_EatModeMgr) Reset() *_EatModeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EatModeMgr) Get() (result EatMode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EatMode{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EatModeMgr) Gets() (results []*EatMode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EatMode{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EatModeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(EatMode{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithActive active获取
func (obj *_EatModeMgr) WithActive(active int8) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithFlag flag获取
func (obj *_EatModeMgr) WithFlag(flag int8) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithCreateTime create_time获取
func (obj *_EatModeMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_EatModeMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithID id获取 主键
func (obj *_EatModeMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_EatModeMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// GetByOption 功能选项模式获取
func (obj *_EatModeMgr) GetByOption(opts ...Option) (result EatMode, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EatMode{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EatModeMgr) GetByOptions(opts ...Option) (results []*EatMode, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(EatMode{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EatModeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]EatMode, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(EatMode{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromActive 通过active获取内容
func (obj *_EatModeMgr) GetFromActive(active int8) (results []*EatMode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EatMode{}).Where("`active` = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量查找
func (obj *_EatModeMgr) GetBatchFromActive(actives []int8) (results []*EatMode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EatMode{}).Where("`active` IN (?)", actives).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容
func (obj *_EatModeMgr) GetFromFlag(flag int8) (results []*EatMode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EatMode{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找
func (obj *_EatModeMgr) GetBatchFromFlag(flags []int8) (results []*EatMode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EatMode{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_EatModeMgr) GetFromCreateTime(createTime time.Time) (results []*EatMode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EatMode{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_EatModeMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*EatMode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EatMode{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_EatModeMgr) GetFromUpdateTime(updateTime time.Time) (results []*EatMode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EatMode{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_EatModeMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*EatMode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EatMode{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromID 通过id获取内容 主键
func (obj *_EatModeMgr) GetFromID(id string) (result EatMode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EatMode{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_EatModeMgr) GetBatchFromID(ids []string) (results []*EatMode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EatMode{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_EatModeMgr) GetFromName(name string) (results []*EatMode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EatMode{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_EatModeMgr) GetBatchFromName(names []string) (results []*EatMode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EatMode{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EatModeMgr) FetchByPrimaryKey(id string) (result EatMode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(EatMode{}).Where("`id` = ?", id).First(&result).Error

	return
}
