package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ConsumerTagToFoodTagMgr struct {
	*_BaseMgr
}

// ConsumerTagToFoodTagMgr open func
func ConsumerTagToFoodTagMgr(db *gorm.DB) *_ConsumerTagToFoodTagMgr {
	if db == nil {
		panic(fmt.Errorf("ConsumerTagToFoodTagMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ConsumerTagToFoodTagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("consumer_tag_to_food_tag"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ConsumerTagToFoodTagMgr) GetTableName() string {
	return "consumer_tag_to_food_tag"
}

// Reset 重置gorm会话
func (obj *_ConsumerTagToFoodTagMgr) Reset() *_ConsumerTagToFoodTagMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ConsumerTagToFoodTagMgr) Get() (result ConsumerTagToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ConsumerTagToFoodTagMgr) Gets() (results []*ConsumerTagToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ConsumerTagToFoodTagMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithActive active获取
func (obj *_ConsumerTagToFoodTagMgr) WithActive(active int8) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithFlag flag获取
func (obj *_ConsumerTagToFoodTagMgr) WithFlag(flag int8) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithCreateTime create_time获取
func (obj *_ConsumerTagToFoodTagMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_ConsumerTagToFoodTagMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithID id获取 主键
func (obj *_ConsumerTagToFoodTagMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithConsumerTagID consumer_tag_id获取
func (obj *_ConsumerTagToFoodTagMgr) WithConsumerTagID(consumerTagID string) Option {
	return optionFunc(func(o *options) { o.query["consumer_tag_id"] = consumerTagID })
}

// WithFoodTagID food_tag_id获取
func (obj *_ConsumerTagToFoodTagMgr) WithFoodTagID(foodTagID string) Option {
	return optionFunc(func(o *options) { o.query["food_tag_id"] = foodTagID })
}

// WithEatMode eat_mode获取
func (obj *_ConsumerTagToFoodTagMgr) WithEatMode(eatMode string) Option {
	return optionFunc(func(o *options) { o.query["eat_mode"] = eatMode })
}

// GetByOption 功能选项模式获取
func (obj *_ConsumerTagToFoodTagMgr) GetByOption(opts ...Option) (result ConsumerTagToFoodTag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ConsumerTagToFoodTagMgr) GetByOptions(opts ...Option) (results []*ConsumerTagToFoodTag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ConsumerTagToFoodTagMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ConsumerTagToFoodTag, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Where(options.query)
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
func (obj *_ConsumerTagToFoodTagMgr) GetFromActive(active int8) (results []*ConsumerTagToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Where("`active` = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量查找
func (obj *_ConsumerTagToFoodTagMgr) GetBatchFromActive(actives []int8) (results []*ConsumerTagToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Where("`active` IN (?)", actives).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容
func (obj *_ConsumerTagToFoodTagMgr) GetFromFlag(flag int8) (results []*ConsumerTagToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找
func (obj *_ConsumerTagToFoodTagMgr) GetBatchFromFlag(flags []int8) (results []*ConsumerTagToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_ConsumerTagToFoodTagMgr) GetFromCreateTime(createTime time.Time) (results []*ConsumerTagToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_ConsumerTagToFoodTagMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*ConsumerTagToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_ConsumerTagToFoodTagMgr) GetFromUpdateTime(updateTime time.Time) (results []*ConsumerTagToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_ConsumerTagToFoodTagMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*ConsumerTagToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromID 通过id获取内容 主键
func (obj *_ConsumerTagToFoodTagMgr) GetFromID(id string) (result ConsumerTagToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_ConsumerTagToFoodTagMgr) GetBatchFromID(ids []string) (results []*ConsumerTagToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromConsumerTagID 通过consumer_tag_id获取内容
func (obj *_ConsumerTagToFoodTagMgr) GetFromConsumerTagID(consumerTagID string) (results []*ConsumerTagToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Where("`consumer_tag_id` = ?", consumerTagID).Find(&results).Error

	return
}

// GetBatchFromConsumerTagID 批量查找
func (obj *_ConsumerTagToFoodTagMgr) GetBatchFromConsumerTagID(consumerTagIDs []string) (results []*ConsumerTagToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Where("`consumer_tag_id` IN (?)", consumerTagIDs).Find(&results).Error

	return
}

// GetFromFoodTagID 通过food_tag_id获取内容
func (obj *_ConsumerTagToFoodTagMgr) GetFromFoodTagID(foodTagID string) (results []*ConsumerTagToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Where("`food_tag_id` = ?", foodTagID).Find(&results).Error

	return
}

// GetBatchFromFoodTagID 批量查找
func (obj *_ConsumerTagToFoodTagMgr) GetBatchFromFoodTagID(foodTagIDs []string) (results []*ConsumerTagToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Where("`food_tag_id` IN (?)", foodTagIDs).Find(&results).Error

	return
}

// GetFromEatMode 通过eat_mode获取内容
func (obj *_ConsumerTagToFoodTagMgr) GetFromEatMode(eatMode string) (results []*ConsumerTagToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Where("`eat_mode` = ?", eatMode).Find(&results).Error

	return
}

// GetBatchFromEatMode 批量查找
func (obj *_ConsumerTagToFoodTagMgr) GetBatchFromEatMode(eatModes []string) (results []*ConsumerTagToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Where("`eat_mode` IN (?)", eatModes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ConsumerTagToFoodTagMgr) FetchByPrimaryKey(id string) (result ConsumerTagToFoodTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerTagToFoodTag{}).Where("`id` = ?", id).First(&result).Error

	return
}
