package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FoodTagMgr struct {
	*_BaseMgr
}

// FoodTagMgr open func
func FoodTagMgr(db *gorm.DB) *_FoodTagMgr {
	if db == nil {
		panic(fmt.Errorf("FoodTagMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FoodTagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("food_tag"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FoodTagMgr) GetTableName() string {
	return "food_tag"
}

// Reset 重置gorm会话
func (obj *_FoodTagMgr) Reset() *_FoodTagMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FoodTagMgr) Get() (result FoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodTag{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FoodTagMgr) Gets() (results []*FoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodTag{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FoodTagMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FoodTag{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithActive active获取
func (obj *_FoodTagMgr) WithActive(active int8) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithFlag flag获取
func (obj *_FoodTagMgr) WithFlag(flag int8) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithCreateTime create_time获取
func (obj *_FoodTagMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_FoodTagMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithID id获取 主键
func (obj *_FoodTagMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_FoodTagMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithParentID parent_id获取
func (obj *_FoodTagMgr) WithParentID(parentID string) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// GetByOption 功能选项模式获取
func (obj *_FoodTagMgr) GetByOption(opts ...Option) (result FoodTag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FoodTag{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FoodTagMgr) GetByOptions(opts ...Option) (results []*FoodTag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FoodTag{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FoodTagMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FoodTag, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FoodTag{}).Where(options.query)
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
func (obj *_FoodTagMgr) GetFromActive(active int8) (results []*FoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodTag{}).Where("`active` = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量查找
func (obj *_FoodTagMgr) GetBatchFromActive(actives []int8) (results []*FoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodTag{}).Where("`active` IN (?)", actives).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容
func (obj *_FoodTagMgr) GetFromFlag(flag int8) (results []*FoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodTag{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找
func (obj *_FoodTagMgr) GetBatchFromFlag(flags []int8) (results []*FoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodTag{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_FoodTagMgr) GetFromCreateTime(createTime time.Time) (results []*FoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodTag{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_FoodTagMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodTag{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_FoodTagMgr) GetFromUpdateTime(updateTime time.Time) (results []*FoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodTag{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_FoodTagMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodTag{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromID 通过id获取内容 主键
func (obj *_FoodTagMgr) GetFromID(id string) (result FoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodTag{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FoodTagMgr) GetBatchFromID(ids []string) (results []*FoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodTag{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_FoodTagMgr) GetFromName(name string) (results []*FoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodTag{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_FoodTagMgr) GetBatchFromName(names []string) (results []*FoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodTag{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容
func (obj *_FoodTagMgr) GetFromParentID(parentID string) (results []*FoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodTag{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找
func (obj *_FoodTagMgr) GetBatchFromParentID(parentIDs []string) (results []*FoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodTag{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FoodTagMgr) FetchByPrimaryKey(id string) (result FoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodTag{}).Where("`id` = ?", id).First(&result).Error

	return
}
