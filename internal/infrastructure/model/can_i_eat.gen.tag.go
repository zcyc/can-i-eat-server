package model

import (
	tag_repo "can-i-eat/internal/repo/tag"
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
func (obj *_TagMgr) Get() (result tag_repo.TagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TagMgr) Gets() (results []*tag_repo.TagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TagMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithActive active获取 表示数据是否处于可用状态， active = 1 可用，active=0不可用，操作可逆转
func (obj *_TagMgr) WithActive(active int8) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithFlag flag获取 是否删除 1 删除 0 未删除
func (obj *_TagMgr) WithFlag(flag int8) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithCreateTime create_time获取 记录写入时间
func (obj *_TagMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 记录更新时间
func (obj *_TagMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithID id获取 主键
func (obj *_TagMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *_TagMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAlias alias获取 别名
func (obj *_TagMgr) WithAlias(alias string) Option {
	return optionFunc(func(o *options) { o.query["alias"] = alias })
}

// WithCategory category获取 分类
func (obj *_TagMgr) WithCategory(category string) Option {
	return optionFunc(func(o *options) { o.query["category"] = category })
}

// GetByOption 功能选项模式获取
func (obj *_TagMgr) GetByOption(opts ...Option) (result tag_repo.TagDao, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TagMgr) GetByOptions(opts ...Option) (results []*tag_repo.TagDao, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Where(options.query).Find(&results).Error

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
	results := make([]tag_repo.TagDao, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Where(options.query)
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
func (obj *_TagMgr) GetFromActive(active int8) (results []*tag_repo.TagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Where("`active` = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量查找 表示数据是否处于可用状态， active = 1 可用，active=0不可用，操作可逆转
func (obj *_TagMgr) GetBatchFromActive(actives []int8) (results []*tag_repo.TagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Where("`active` IN (?)", actives).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容 是否删除 1 删除 0 未删除
func (obj *_TagMgr) GetFromFlag(flag int8) (results []*tag_repo.TagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找 是否删除 1 删除 0 未删除
func (obj *_TagMgr) GetBatchFromFlag(flags []int8) (results []*tag_repo.TagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 记录写入时间
func (obj *_TagMgr) GetFromCreateTime(createTime time.Time) (results []*tag_repo.TagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 记录写入时间
func (obj *_TagMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*tag_repo.TagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 记录更新时间
func (obj *_TagMgr) GetFromUpdateTime(updateTime time.Time) (results []*tag_repo.TagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 记录更新时间
func (obj *_TagMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*tag_repo.TagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromID 通过id获取内容 主键
func (obj *_TagMgr) GetFromID(id uint64) (result tag_repo.TagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_TagMgr) GetBatchFromID(ids []uint64) (results []*tag_repo.TagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_TagMgr) GetFromName(name string) (results []*tag_repo.TagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 名称
func (obj *_TagMgr) GetBatchFromName(names []string) (results []*tag_repo.TagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromAlias 通过alias获取内容 别名
func (obj *_TagMgr) GetFromAlias(alias string) (results []*tag_repo.TagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Where("`alias` = ?", alias).Find(&results).Error

	return
}

// GetBatchFromAlias 批量查找 别名
func (obj *_TagMgr) GetBatchFromAlias(aliass []string) (results []*tag_repo.TagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Where("`alias` IN (?)", aliass).Find(&results).Error

	return
}

// GetFromCategory 通过category获取内容 分类
func (obj *_TagMgr) GetFromCategory(category string) (results []*tag_repo.TagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Where("`category` = ?", category).Find(&results).Error

	return
}

// GetBatchFromCategory 批量查找 分类
func (obj *_TagMgr) GetBatchFromCategory(categorys []string) (results []*tag_repo.TagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Where("`category` IN (?)", categorys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TagMgr) FetchByPrimaryKey(id uint64) (result tag_repo.TagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(tag_repo.TagDao{}).Where("`id` = ?", id).First(&result).Error

	return
}
