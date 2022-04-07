package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ConsumerToConsumerTagMgr struct {
	*_BaseMgr
}

// ConsumerToConsumerTagMgr open func
func ConsumerToConsumerTagMgr(db *gorm.DB) *_ConsumerToConsumerTagMgr {
	if db == nil {
		panic(fmt.Errorf("ConsumerToConsumerTagMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ConsumerToConsumerTagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("consumer_to_consumer_tag"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ConsumerToConsumerTagMgr) GetTableName() string {
	return "consumer_to_consumer_tag"
}

// Reset 重置gorm会话
func (obj *_ConsumerToConsumerTagMgr) Reset() *_ConsumerToConsumerTagMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ConsumerToConsumerTagMgr) Get() (result ConsumerToConsumerTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerToConsumerTag{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ConsumerToConsumerTagMgr) Gets() (results []*ConsumerToConsumerTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerToConsumerTag{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ConsumerToConsumerTagMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ConsumerToConsumerTag{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithActive active获取
func (obj *_ConsumerToConsumerTagMgr) WithActive(active int8) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithFlag flag获取
func (obj *_ConsumerToConsumerTagMgr) WithFlag(flag int8) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithCreateTime create_time获取
func (obj *_ConsumerToConsumerTagMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_ConsumerToConsumerTagMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithID id获取 主键
func (obj *_ConsumerToConsumerTagMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithConsumerID consumer_id获取
func (obj *_ConsumerToConsumerTagMgr) WithConsumerID(consumerID string) Option {
	return optionFunc(func(o *options) { o.query["consumer_id"] = consumerID })
}

// WithConsumerTagID consumer_tag_id获取
func (obj *_ConsumerToConsumerTagMgr) WithConsumerTagID(consumerTagID string) Option {
	return optionFunc(func(o *options) { o.query["consumer_tag_id"] = consumerTagID })
}

// GetByOption 功能选项模式获取
func (obj *_ConsumerToConsumerTagMgr) GetByOption(opts ...Option) (result ConsumerToConsumerTag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ConsumerToConsumerTag{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ConsumerToConsumerTagMgr) GetByOptions(opts ...Option) (results []*ConsumerToConsumerTag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ConsumerToConsumerTag{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ConsumerToConsumerTagMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ConsumerToConsumerTag, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ConsumerToConsumerTag{}).Where(options.query)
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
func (obj *_ConsumerToConsumerTagMgr) GetFromActive(active int8) (results []*ConsumerToConsumerTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerToConsumerTag{}).Where("`active` = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量查找
func (obj *_ConsumerToConsumerTagMgr) GetBatchFromActive(actives []int8) (results []*ConsumerToConsumerTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerToConsumerTag{}).Where("`active` IN (?)", actives).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容
func (obj *_ConsumerToConsumerTagMgr) GetFromFlag(flag int8) (results []*ConsumerToConsumerTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerToConsumerTag{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找
func (obj *_ConsumerToConsumerTagMgr) GetBatchFromFlag(flags []int8) (results []*ConsumerToConsumerTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerToConsumerTag{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_ConsumerToConsumerTagMgr) GetFromCreateTime(createTime time.Time) (results []*ConsumerToConsumerTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerToConsumerTag{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_ConsumerToConsumerTagMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*ConsumerToConsumerTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerToConsumerTag{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_ConsumerToConsumerTagMgr) GetFromUpdateTime(updateTime time.Time) (results []*ConsumerToConsumerTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerToConsumerTag{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_ConsumerToConsumerTagMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*ConsumerToConsumerTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerToConsumerTag{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromID 通过id获取内容 主键
func (obj *_ConsumerToConsumerTagMgr) GetFromID(id string) (result ConsumerToConsumerTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerToConsumerTag{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_ConsumerToConsumerTagMgr) GetBatchFromID(ids []string) (results []*ConsumerToConsumerTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerToConsumerTag{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromConsumerID 通过consumer_id获取内容
func (obj *_ConsumerToConsumerTagMgr) GetFromConsumerID(consumerID string) (results []*ConsumerToConsumerTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerToConsumerTag{}).Where("`consumer_id` = ?", consumerID).Find(&results).Error

	return
}

// GetBatchFromConsumerID 批量查找
func (obj *_ConsumerToConsumerTagMgr) GetBatchFromConsumerID(consumerIDs []string) (results []*ConsumerToConsumerTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerToConsumerTag{}).Where("`consumer_id` IN (?)", consumerIDs).Find(&results).Error

	return
}

// GetFromConsumerTagID 通过consumer_tag_id获取内容
func (obj *_ConsumerToConsumerTagMgr) GetFromConsumerTagID(consumerTagID string) (results []*ConsumerToConsumerTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerToConsumerTag{}).Where("`consumer_tag_id` = ?", consumerTagID).Find(&results).Error

	return
}

// GetBatchFromConsumerTagID 批量查找
func (obj *_ConsumerToConsumerTagMgr) GetBatchFromConsumerTagID(consumerTagIDs []string) (results []*ConsumerToConsumerTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerToConsumerTag{}).Where("`consumer_tag_id` IN (?)", consumerTagIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ConsumerToConsumerTagMgr) FetchByPrimaryKey(id string) (result ConsumerToConsumerTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ConsumerToConsumerTag{}).Where("`id` = ?", id).First(&result).Error

	return
}
