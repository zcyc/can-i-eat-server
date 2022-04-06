package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ConsumerGroupMgr struct {
	*_BaseMgr
}

// ConsumerGroupMgr open func
func ConsumerGroupMgr(db *gorm.DB) *_ConsumerGroupMgr {
	if db == nil {
		panic(fmt.Errorf("ConsumerGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ConsumerGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("consumer_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ConsumerGroupMgr) GetTableName() string {
	return "consumer_group"
}

// Reset 重置gorm会话
func (obj *_ConsumerGroupMgr) Reset() *_ConsumerGroupMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ConsumerGroupMgr) Get() (result ConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerGroup{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ConsumerGroupMgr) Gets() (results []*ConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerGroup{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ConsumerGroupMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ConsumerGroup{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithActive active获取
func (obj *_ConsumerGroupMgr) WithActive(active int8) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithFlag flag获取
func (obj *_ConsumerGroupMgr) WithFlag(flag int8) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithCreateTime create_time获取
func (obj *_ConsumerGroupMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_ConsumerGroupMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithID id获取 主键
func (obj *_ConsumerGroupMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithConsumerID consumer_id获取
func (obj *_ConsumerGroupMgr) WithConsumerID(consumerID int64) Option {
	return optionFunc(func(o *options) { o.query["consumer_id"] = consumerID })
}

// WithGroupID group_id获取
func (obj *_ConsumerGroupMgr) WithGroupID(groupID int64) Option {
	return optionFunc(func(o *options) { o.query["group_id"] = groupID })
}

// GetByOption 功能选项模式获取
func (obj *_ConsumerGroupMgr) GetByOption(opts ...Option) (result ConsumerGroup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ConsumerGroup{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ConsumerGroupMgr) GetByOptions(opts ...Option) (results []*ConsumerGroup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ConsumerGroup{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ConsumerGroupMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ConsumerGroup, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ConsumerGroup{}).Where(options.query)
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
func (obj *_ConsumerGroupMgr) GetFromActive(active int8) (results []*ConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerGroup{}).Where("`active` = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量查找
func (obj *_ConsumerGroupMgr) GetBatchFromActive(actives []int8) (results []*ConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerGroup{}).Where("`active` IN (?)", actives).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容
func (obj *_ConsumerGroupMgr) GetFromFlag(flag int8) (results []*ConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerGroup{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找
func (obj *_ConsumerGroupMgr) GetBatchFromFlag(flags []int8) (results []*ConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerGroup{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_ConsumerGroupMgr) GetFromCreateTime(createTime time.Time) (results []*ConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerGroup{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_ConsumerGroupMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*ConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerGroup{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_ConsumerGroupMgr) GetFromUpdateTime(updateTime time.Time) (results []*ConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerGroup{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_ConsumerGroupMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*ConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerGroup{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromID 通过id获取内容 主键
func (obj *_ConsumerGroupMgr) GetFromID(id uint64) (result ConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerGroup{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_ConsumerGroupMgr) GetBatchFromID(ids []uint64) (results []*ConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerGroup{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromConsumerID 通过consumer_id获取内容
func (obj *_ConsumerGroupMgr) GetFromConsumerID(consumerID int64) (results []*ConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerGroup{}).Where("`consumer_id` = ?", consumerID).Find(&results).Error

	return
}

// GetBatchFromConsumerID 批量查找
func (obj *_ConsumerGroupMgr) GetBatchFromConsumerID(consumerIDs []int64) (results []*ConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerGroup{}).Where("`consumer_id` IN (?)", consumerIDs).Find(&results).Error

	return
}

// GetFromGroupID 通过group_id获取内容
func (obj *_ConsumerGroupMgr) GetFromGroupID(groupID int64) (results []*ConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerGroup{}).Where("`group_id` = ?", groupID).Find(&results).Error

	return
}

// GetBatchFromGroupID 批量查找
func (obj *_ConsumerGroupMgr) GetBatchFromGroupID(groupIDs []int64) (results []*ConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerGroup{}).Where("`group_id` IN (?)", groupIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ConsumerGroupMgr) FetchByPrimaryKey(id uint64) (result ConsumerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerGroup{}).Where("`id` = ?", id).First(&result).Error

	return
}
