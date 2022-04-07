package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FoodToFoodTagMgr struct {
	*_BaseMgr
}

// FoodToFoodTagMgr open func
func FoodToFoodTagMgr(db *gorm.DB) *_FoodToFoodTagMgr {
	if db == nil {
		panic(fmt.Errorf("FoodToFoodTagMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FoodToFoodTagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("food_to_food_tag"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FoodToFoodTagMgr) GetTableName() string {
	return "food_to_food_tag"
}

// Reset 重置gorm会话
func (obj *_FoodToFoodTagMgr) Reset() *_FoodToFoodTagMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FoodToFoodTagMgr) Get() (result FoodToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodToFoodTag{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FoodToFoodTagMgr) Gets() (results []*FoodToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodToFoodTag{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FoodToFoodTagMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FoodToFoodTag{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithActive active获取
func (obj *_FoodToFoodTagMgr) WithActive(active int8) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithFlag flag获取
func (obj *_FoodToFoodTagMgr) WithFlag(flag int8) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithCreateTime create_time获取
func (obj *_FoodToFoodTagMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_FoodToFoodTagMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithID id获取 主键
func (obj *_FoodToFoodTagMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithFoodID food_id获取
func (obj *_FoodToFoodTagMgr) WithFoodID(foodID string) Option {
	return optionFunc(func(o *options) { o.query["food_id"] = foodID })
}

// WithFoodTagID food_tag_id获取
func (obj *_FoodToFoodTagMgr) WithFoodTagID(foodTagID string) Option {
	return optionFunc(func(o *options) { o.query["food_tag_id"] = foodTagID })
}

// GetByOption 功能选项模式获取
func (obj *_FoodToFoodTagMgr) GetByOption(opts ...Option) (result FoodToFoodTag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FoodToFoodTag{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FoodToFoodTagMgr) GetByOptions(opts ...Option) (results []*FoodToFoodTag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FoodToFoodTag{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FoodToFoodTagMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FoodToFoodTag, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FoodToFoodTag{}).Where(options.query)
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
func (obj *_FoodToFoodTagMgr) GetFromActive(active int8) (results []*FoodToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodToFoodTag{}).Where("`active` = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量查找
func (obj *_FoodToFoodTagMgr) GetBatchFromActive(actives []int8) (results []*FoodToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodToFoodTag{}).Where("`active` IN (?)", actives).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容
func (obj *_FoodToFoodTagMgr) GetFromFlag(flag int8) (results []*FoodToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodToFoodTag{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找
func (obj *_FoodToFoodTagMgr) GetBatchFromFlag(flags []int8) (results []*FoodToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodToFoodTag{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_FoodToFoodTagMgr) GetFromCreateTime(createTime time.Time) (results []*FoodToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodToFoodTag{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_FoodToFoodTagMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FoodToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodToFoodTag{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_FoodToFoodTagMgr) GetFromUpdateTime(updateTime time.Time) (results []*FoodToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodToFoodTag{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_FoodToFoodTagMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FoodToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodToFoodTag{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromID 通过id获取内容 主键
func (obj *_FoodToFoodTagMgr) GetFromID(id string) (result FoodToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodToFoodTag{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FoodToFoodTagMgr) GetBatchFromID(ids []string) (results []*FoodToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodToFoodTag{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromFoodID 通过food_id获取内容
func (obj *_FoodToFoodTagMgr) GetFromFoodID(foodID string) (results []*FoodToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodToFoodTag{}).Where("`food_id` = ?", foodID).Find(&results).Error

	return
}

// GetBatchFromFoodID 批量查找
func (obj *_FoodToFoodTagMgr) GetBatchFromFoodID(foodIDs []string) (results []*FoodToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodToFoodTag{}).Where("`food_id` IN (?)", foodIDs).Find(&results).Error

	return
}

// GetFromFoodTagID 通过food_tag_id获取内容
func (obj *_FoodToFoodTagMgr) GetFromFoodTagID(foodTagID string) (results []*FoodToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodToFoodTag{}).Where("`food_tag_id` = ?", foodTagID).Find(&results).Error

	return
}

// GetBatchFromFoodTagID 批量查找
func (obj *_FoodToFoodTagMgr) GetBatchFromFoodTagID(foodTagIDs []string) (results []*FoodToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodToFoodTag{}).Where("`food_tag_id` IN (?)", foodTagIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FoodToFoodTagMgr) FetchByPrimaryKey(id string) (result FoodToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FoodToFoodTag{}).Where("`id` = ?", id).First(&result).Error

	return
}
