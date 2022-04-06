package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _TagMgr struct {
	*_BaseMgr
}

// TagMgr open func
func TagMgr(db *gorm.DB) *_TagMgr {
	if db == nil {
		panic(fmt.Errorf("TagMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tag"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TagMgr) GetTableName() string {
	return "tag"
}

// Reset 重置gorm会话
func (obj *_TagMgr) Reset() *_TagMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TagMgr) Get() (result Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tag{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TagMgr) Gets() (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tag{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TagMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Tag{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithActive active获取
func (obj *_TagMgr) WithActive(active int8) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithFlag flag获取
func (obj *_TagMgr) WithFlag(flag int8) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithCreateTime create_time获取
func (obj *_TagMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_TagMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithID id获取 主键
func (obj *_TagMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_TagMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithParentID parent_id获取
func (obj *_TagMgr) WithParentID(parentID string) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// GetByOption 功能选项模式获取
func (obj *_TagMgr) GetByOption(opts ...Option) (result Tag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Tag{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TagMgr) GetByOptions(opts ...Option) (results []*Tag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Tag{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TagMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Tag, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Tag{}).Where(options.query)
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
func (obj *_TagMgr) GetFromActive(active int8) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tag{}).Where("`active` = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量查找
func (obj *_TagMgr) GetBatchFromActive(actives []int8) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tag{}).Where("`active` IN (?)", actives).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容
func (obj *_TagMgr) GetFromFlag(flag int8) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tag{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找
func (obj *_TagMgr) GetBatchFromFlag(flags []int8) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tag{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_TagMgr) GetFromCreateTime(createTime time.Time) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tag{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_TagMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tag{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_TagMgr) GetFromUpdateTime(updateTime time.Time) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tag{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_TagMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tag{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromID 通过id获取内容 主键
func (obj *_TagMgr) GetFromID(id string) (result Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tag{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_TagMgr) GetBatchFromID(ids []string) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tag{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_TagMgr) GetFromName(name string) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tag{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_TagMgr) GetBatchFromName(names []string) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tag{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容
func (obj *_TagMgr) GetFromParentID(parentID string) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tag{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找
func (obj *_TagMgr) GetBatchFromParentID(parentIDs []string) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tag{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TagMgr) FetchByPrimaryKey(id string) (result Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tag{}).Where("`id` = ?", id).First(&result).Error

	return
}
