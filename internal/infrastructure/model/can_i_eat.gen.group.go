package model

import (
	group_repo "can-i-eat/internal/repo/group"
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _GroupMgr struct {
	*_BaseMgr
}

// GroupMgr open func
func GroupMgr(db *gorm.DB) *_GroupMgr {
	if db == nil {
		panic(fmt.Errorf("GroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GroupMgr) GetTableName() string {
	return "group"
}

// Reset 重置gorm会话
func (obj *_GroupMgr) Reset() *_GroupMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GroupMgr) Get() (result group_repo.GroupDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(group_repo.GroupDao{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GroupMgr) Gets() (results []*group_repo.GroupDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(group_repo.GroupDao{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GroupMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(group_repo.GroupDao{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithActive active获取 表示数据是否处于可用状态， active = 1 可用，active=0不可用，操作可逆转
func (obj *_GroupMgr) WithActive(active int8) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithFlag flag获取 是否删除 1 删除 0 未删除
func (obj *_GroupMgr) WithFlag(flag int8) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithCreateTime create_time获取 记录写入时间
func (obj *_GroupMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 记录更新时间
func (obj *_GroupMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithID id获取 主键
func (obj *_GroupMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *_GroupMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// GetByOption 功能选项模式获取
func (obj *_GroupMgr) GetByOption(opts ...Option) (result group_repo.GroupDao, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(group_repo.GroupDao{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GroupMgr) GetByOptions(opts ...Option) (results []*group_repo.GroupDao, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(group_repo.GroupDao{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_GroupMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]group_repo.GroupDao, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(group_repo.GroupDao{}).Where(options.query)
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

// GetFromActive 通过active获取内容 表示数据是否处于可用状态， active = 1 可用，active=0不可用，操作可逆转
func (obj *_GroupMgr) GetFromActive(active int8) (results []*group_repo.GroupDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(group_repo.GroupDao{}).Where("`active` = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量查找 表示数据是否处于可用状态， active = 1 可用，active=0不可用，操作可逆转
func (obj *_GroupMgr) GetBatchFromActive(actives []int8) (results []*group_repo.GroupDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(group_repo.GroupDao{}).Where("`active` IN (?)", actives).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容 是否删除 1 删除 0 未删除
func (obj *_GroupMgr) GetFromFlag(flag int8) (results []*group_repo.GroupDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(group_repo.GroupDao{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找 是否删除 1 删除 0 未删除
func (obj *_GroupMgr) GetBatchFromFlag(flags []int8) (results []*group_repo.GroupDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(group_repo.GroupDao{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 记录写入时间
func (obj *_GroupMgr) GetFromCreateTime(createTime time.Time) (results []*group_repo.GroupDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(group_repo.GroupDao{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 记录写入时间
func (obj *_GroupMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*group_repo.GroupDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(group_repo.GroupDao{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 记录更新时间
func (obj *_GroupMgr) GetFromUpdateTime(updateTime time.Time) (results []*group_repo.GroupDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(group_repo.GroupDao{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 记录更新时间
func (obj *_GroupMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*group_repo.GroupDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(group_repo.GroupDao{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromID 通过id获取内容 主键
func (obj *_GroupMgr) GetFromID(id uint64) (result group_repo.GroupDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(group_repo.GroupDao{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_GroupMgr) GetBatchFromID(ids []uint64) (results []*group_repo.GroupDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(group_repo.GroupDao{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_GroupMgr) GetFromName(name string) (results []*group_repo.GroupDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(group_repo.GroupDao{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 名称
func (obj *_GroupMgr) GetBatchFromName(names []string) (results []*group_repo.GroupDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(group_repo.GroupDao{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GroupMgr) FetchByPrimaryKey(id uint64) (result group_repo.GroupDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(group_repo.GroupDao{}).Where("`id` = ?", id).First(&result).Error

	return
}