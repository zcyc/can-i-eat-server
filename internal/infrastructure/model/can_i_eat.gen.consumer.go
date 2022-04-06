package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ConsumerMgr struct {
	*_BaseMgr
}

// ConsumerMgr open func
func ConsumerMgr(db *gorm.DB) *_ConsumerMgr {
	if db == nil {
		panic(fmt.Errorf("ConsumerMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ConsumerMgr{_BaseMgr: &_BaseMgr{DB: db.Table("consumer"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ConsumerMgr) GetTableName() string {
	return "consumer"
}

// Reset 重置gorm会话
func (obj *_ConsumerMgr) Reset() *_ConsumerMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ConsumerMgr) Get() (result Consumer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Consumer{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ConsumerMgr) Gets() (results []*Consumer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Consumer{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ConsumerMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Consumer{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithActive active获取
func (obj *_ConsumerMgr) WithActive(active int8) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithFlag flag获取
func (obj *_ConsumerMgr) WithFlag(flag int8) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithCreateTime create_time获取
func (obj *_ConsumerMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_ConsumerMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithID id获取 主键
func (obj *_ConsumerMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_ConsumerMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAccount account获取
func (obj *_ConsumerMgr) WithAccount(account string) Option {
	return optionFunc(func(o *options) { o.query["account"] = account })
}

// WithPassword password获取
func (obj *_ConsumerMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// GetByOption 功能选项模式获取
func (obj *_ConsumerMgr) GetByOption(opts ...Option) (result Consumer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Consumer{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ConsumerMgr) GetByOptions(opts ...Option) (results []*Consumer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Consumer{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ConsumerMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Consumer, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Consumer{}).Where(options.query)
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
func (obj *_ConsumerMgr) GetFromActive(active int8) (results []*Consumer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Consumer{}).Where("`active` = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量查找
func (obj *_ConsumerMgr) GetBatchFromActive(actives []int8) (results []*Consumer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Consumer{}).Where("`active` IN (?)", actives).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容
func (obj *_ConsumerMgr) GetFromFlag(flag int8) (results []*Consumer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Consumer{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找
func (obj *_ConsumerMgr) GetBatchFromFlag(flags []int8) (results []*Consumer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Consumer{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_ConsumerMgr) GetFromCreateTime(createTime time.Time) (results []*Consumer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Consumer{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_ConsumerMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Consumer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Consumer{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_ConsumerMgr) GetFromUpdateTime(updateTime time.Time) (results []*Consumer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Consumer{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_ConsumerMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*Consumer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Consumer{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromID 通过id获取内容 主键
func (obj *_ConsumerMgr) GetFromID(id uint64) (result Consumer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Consumer{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_ConsumerMgr) GetBatchFromID(ids []uint64) (results []*Consumer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Consumer{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_ConsumerMgr) GetFromName(name string) (results []*Consumer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Consumer{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_ConsumerMgr) GetBatchFromName(names []string) (results []*Consumer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Consumer{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromAccount 通过account获取内容
func (obj *_ConsumerMgr) GetFromAccount(account string) (results []*Consumer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Consumer{}).Where("`account` = ?", account).Find(&results).Error

	return
}

// GetBatchFromAccount 批量查找
func (obj *_ConsumerMgr) GetBatchFromAccount(accounts []string) (results []*Consumer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Consumer{}).Where("`account` IN (?)", accounts).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容
func (obj *_ConsumerMgr) GetFromPassword(password string) (results []*Consumer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Consumer{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找
func (obj *_ConsumerMgr) GetBatchFromPassword(passwords []string) (results []*Consumer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Consumer{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ConsumerMgr) FetchByPrimaryKey(id uint64) (result Consumer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Consumer{}).Where("`id` = ?", id).First(&result).Error

	return
}
