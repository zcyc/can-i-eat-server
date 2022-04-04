package model

import (
	category_repo "can-i-eat/internal/repo/category"
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CategoryMgr struct {
	*_BaseMgr
}

// CategoryMgr open func
func CategoryMgr(db *gorm.DB) *_CategoryMgr {
	if db == nil {
		panic(fmt.Errorf("CategoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CategoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("category"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CategoryMgr) GetTableName() string {
	return "category"
}

// Reset 重置gorm会话
func (obj *_CategoryMgr) Reset() *_CategoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CategoryMgr) Get() (result category_repo.CategoryDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(category_repo.CategoryDao{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CategoryMgr) Gets() (results []*category_repo.CategoryDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(category_repo.CategoryDao{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CategoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(category_repo.CategoryDao{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithActive active获取 表示数据是否处于可用状态， active = 1 可用，active=0不可用，操作可逆转
func (obj *_CategoryMgr) WithActive(active int8) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithFlag flag获取 是否删除 1 删除 0 未删除
func (obj *_CategoryMgr) WithFlag(flag int8) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithCreateTime create_time获取 记录写入时间
func (obj *_CategoryMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 记录更新时间
func (obj *_CategoryMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithID id获取 主键
func (obj *_CategoryMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *_CategoryMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// GetByOption 功能选项模式获取
func (obj *_CategoryMgr) GetByOption(opts ...Option) (result category_repo.CategoryDao, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(category_repo.CategoryDao{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CategoryMgr) GetByOptions(opts ...Option) (results []*category_repo.CategoryDao, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(category_repo.CategoryDao{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CategoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]category_repo.CategoryDao, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(category_repo.CategoryDao{}).Where(options.query)
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
func (obj *_CategoryMgr) GetFromActive(active int8) (results []*category_repo.CategoryDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(category_repo.CategoryDao{}).Where("`active` = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量查找 表示数据是否处于可用状态， active = 1 可用，active=0不可用，操作可逆转
func (obj *_CategoryMgr) GetBatchFromActive(actives []int8) (results []*category_repo.CategoryDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(category_repo.CategoryDao{}).Where("`active` IN (?)", actives).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容 是否删除 1 删除 0 未删除
func (obj *_CategoryMgr) GetFromFlag(flag int8) (results []*category_repo.CategoryDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(category_repo.CategoryDao{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找 是否删除 1 删除 0 未删除
func (obj *_CategoryMgr) GetBatchFromFlag(flags []int8) (results []*category_repo.CategoryDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(category_repo.CategoryDao{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 记录写入时间
func (obj *_CategoryMgr) GetFromCreateTime(createTime time.Time) (results []*category_repo.CategoryDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(category_repo.CategoryDao{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 记录写入时间
func (obj *_CategoryMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*category_repo.CategoryDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(category_repo.CategoryDao{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 记录更新时间
func (obj *_CategoryMgr) GetFromUpdateTime(updateTime time.Time) (results []*category_repo.CategoryDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(category_repo.CategoryDao{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 记录更新时间
func (obj *_CategoryMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*category_repo.CategoryDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(category_repo.CategoryDao{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromID 通过id获取内容 主键
func (obj *_CategoryMgr) GetFromID(id uint64) (result category_repo.CategoryDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(category_repo.CategoryDao{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_CategoryMgr) GetBatchFromID(ids []uint64) (results []*category_repo.CategoryDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(category_repo.CategoryDao{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_CategoryMgr) GetFromName(name string) (results []*category_repo.CategoryDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(category_repo.CategoryDao{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 名称
func (obj *_CategoryMgr) GetBatchFromName(names []string) (results []*category_repo.CategoryDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(category_repo.CategoryDao{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_CategoryMgr) FetchByPrimaryKey(id uint64) (result category_repo.CategoryDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(category_repo.CategoryDao{}).Where("`id` = ?", id).First(&result).Error

	return
}
