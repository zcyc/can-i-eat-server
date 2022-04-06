package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _GroupTagMgr struct {
	*_BaseMgr
}

// GroupTagMgr open func
func GroupTagMgr(db *gorm.DB) *_GroupTagMgr {
	if db == nil {
		panic(fmt.Errorf("GroupTagMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GroupTagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("group_tag"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GroupTagMgr) GetTableName() string {
	return "group_tag"
}

// Reset 重置gorm会话
func (obj *_GroupTagMgr) Reset() *_GroupTagMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GroupTagMgr) Get() (result GroupTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTag{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GroupTagMgr) Gets() (results []*GroupTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GroupTagMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithActive active获取
func (obj *_GroupTagMgr) WithActive(active int8) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithFlag flag获取
func (obj *_GroupTagMgr) WithFlag(flag int8) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithCreateTime create_time获取
func (obj *_GroupTagMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_GroupTagMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithID id获取 主键
func (obj *_GroupTagMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGroupID group_id获取
func (obj *_GroupTagMgr) WithGroupID(groupID string) Option {
	return optionFunc(func(o *options) { o.query["group_id"] = groupID })
}

// WithTagID tag_id获取
func (obj *_GroupTagMgr) WithTagID(tagID string) Option {
	return optionFunc(func(o *options) { o.query["tag_id"] = tagID })
}

// WithType type获取
func (obj *_GroupTagMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// GetByOption 功能选项模式获取
func (obj *_GroupTagMgr) GetByOption(opts ...Option) (result GroupTag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GroupTagMgr) GetByOptions(opts ...Option) (results []*GroupTag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_GroupTagMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]GroupTag, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Where(options.query)
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
func (obj *_GroupTagMgr) GetFromActive(active int8) (results []*GroupTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Where("`active` = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量查找
func (obj *_GroupTagMgr) GetBatchFromActive(actives []int8) (results []*GroupTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Where("`active` IN (?)", actives).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容
func (obj *_GroupTagMgr) GetFromFlag(flag int8) (results []*GroupTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找
func (obj *_GroupTagMgr) GetBatchFromFlag(flags []int8) (results []*GroupTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_GroupTagMgr) GetFromCreateTime(createTime time.Time) (results []*GroupTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_GroupTagMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*GroupTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_GroupTagMgr) GetFromUpdateTime(updateTime time.Time) (results []*GroupTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_GroupTagMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*GroupTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromID 通过id获取内容 主键
func (obj *_GroupTagMgr) GetFromID(id string) (result GroupTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_GroupTagMgr) GetBatchFromID(ids []string) (results []*GroupTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromGroupID 通过group_id获取内容
func (obj *_GroupTagMgr) GetFromGroupID(groupID string) (results []*GroupTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Where("`group_id` = ?", groupID).Find(&results).Error

	return
}

// GetBatchFromGroupID 批量查找
func (obj *_GroupTagMgr) GetBatchFromGroupID(groupIDs []string) (results []*GroupTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Where("`group_id` IN (?)", groupIDs).Find(&results).Error

	return
}

// GetFromTagID 通过tag_id获取内容
func (obj *_GroupTagMgr) GetFromTagID(tagID string) (results []*GroupTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Where("`tag_id` = ?", tagID).Find(&results).Error

	return
}

// GetBatchFromTagID 批量查找
func (obj *_GroupTagMgr) GetBatchFromTagID(tagIDs []string) (results []*GroupTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Where("`tag_id` IN (?)", tagIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_GroupTagMgr) GetFromType(_type string) (results []*GroupTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_GroupTagMgr) GetBatchFromType(_types []string) (results []*GroupTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GroupTagMgr) FetchByPrimaryKey(id string) (result GroupTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTag{}).Where("`id` = ?", id).First(&result).Error

	return
}
