package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FoodMgr struct {
	*_BaseMgr
}

// FoodMgr open func
func FoodMgr(db *gorm.DB) *_FoodMgr {
	if db == nil {
		panic(fmt.Errorf("FoodMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FoodMgr{_BaseMgr: &_BaseMgr{DB: db.Table("food"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FoodMgr) GetTableName() string {
	return "food"
}

// Reset 重置gorm会话
func (obj *_FoodMgr) Reset() *_FoodMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FoodMgr) Get() (result Food, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Food{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FoodMgr) Gets() (results []*Food, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Food{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FoodMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Food{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithActive active获取
func (obj *_FoodMgr) WithActive(active int8) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithFlag flag获取
func (obj *_FoodMgr) WithFlag(flag int8) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithCreateTime create_time获取
func (obj *_FoodMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_FoodMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithID id获取 主键
func (obj *_FoodMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCategoryID category_id获取
func (obj *_FoodMgr) WithCategoryID(categoryID uint64) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithName name获取
func (obj *_FoodMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAlias alias获取
func (obj *_FoodMgr) WithAlias(alias string) Option {
	return optionFunc(func(o *options) { o.query["alias"] = alias })
}

// GetByOption 功能选项模式获取
func (obj *_FoodMgr) GetByOption(opts ...Option) (result Food, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Food{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FoodMgr) GetByOptions(opts ...Option) (results []*Food, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Food{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FoodMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Food, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Food{}).Where(options.query)
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
func (obj *_FoodMgr) GetFromActive(active int8) (results []*Food, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Food{}).Where("`active` = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量查找
func (obj *_FoodMgr) GetBatchFromActive(actives []int8) (results []*Food, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Food{}).Where("`active` IN (?)", actives).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容
func (obj *_FoodMgr) GetFromFlag(flag int8) (results []*Food, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Food{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找
func (obj *_FoodMgr) GetBatchFromFlag(flags []int8) (results []*Food, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Food{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_FoodMgr) GetFromCreateTime(createTime time.Time) (results []*Food, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Food{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_FoodMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Food, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Food{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_FoodMgr) GetFromUpdateTime(updateTime time.Time) (results []*Food, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Food{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_FoodMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*Food, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Food{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromID 通过id获取内容 主键
func (obj *_FoodMgr) GetFromID(id uint64) (result Food, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Food{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FoodMgr) GetBatchFromID(ids []uint64) (results []*Food, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Food{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCategoryID 通过category_id获取内容
func (obj *_FoodMgr) GetFromCategoryID(categoryID uint64) (results []*Food, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Food{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找
func (obj *_FoodMgr) GetBatchFromCategoryID(categoryIDs []uint64) (results []*Food, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Food{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_FoodMgr) GetFromName(name string) (results []*Food, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Food{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_FoodMgr) GetBatchFromName(names []string) (results []*Food, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Food{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromAlias 通过alias获取内容
func (obj *_FoodMgr) GetFromAlias(alias string) (results []*Food, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Food{}).Where("`alias` = ?", alias).Find(&results).Error

	return
}

// GetBatchFromAlias 批量查找
func (obj *_FoodMgr) GetBatchFromAlias(aliass []string) (results []*Food, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Food{}).Where("`alias` IN (?)", aliass).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FoodMgr) FetchByPrimaryKey(id uint64) (result Food, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Food{}).Where("`id` = ?", id).First(&result).Error

	return
}
