package model

import (
	food_tag_repo "can-i-eat/internal/repo/food_tag"
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FoodTagMgr struct {
	*_BaseMgr
}

// FoodTagMgr open func
func FoodTagMgr(db *gorm.DB) *_FoodTagMgr {
	if db == nil {
		panic(fmt.Errorf("FoodTagMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FoodTagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("food_tag"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FoodTagMgr) GetTableName() string {
	return "food_tag"
}

// Reset 重置gorm会话
func (obj *_FoodTagMgr) Reset() *_FoodTagMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FoodTagMgr) Get() (result food_tag_repo.FoodTagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FoodTagMgr) Gets() (results []*food_tag_repo.FoodTagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FoodTagMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithActive active获取 表示数据是否处于可用状态， active = 1 可用，active=0不可用，操作可逆转
func (obj *_FoodTagMgr) WithActive(active int8) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithFlag flag获取 是否删除 1 删除 0 未删除
func (obj *_FoodTagMgr) WithFlag(flag int8) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithCreateTime create_time获取 记录写入时间
func (obj *_FoodTagMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 记录更新时间
func (obj *_FoodTagMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithID id获取 主键
func (obj *_FoodTagMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithFoodID food_id获取 名称
func (obj *_FoodTagMgr) WithFoodID(foodID int64) Option {
	return optionFunc(func(o *options) { o.query["food_id"] = foodID })
}

// WithTagID tag_id获取
func (obj *_FoodTagMgr) WithTagID(tagID int64) Option {
	return optionFunc(func(o *options) { o.query["tag_id"] = tagID })
}

// WithType type获取
func (obj *_FoodTagMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// GetByOption 功能选项模式获取
func (obj *_FoodTagMgr) GetByOption(opts ...Option) (result food_tag_repo.FoodTagDao, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FoodTagMgr) GetByOptions(opts ...Option) (results []*food_tag_repo.FoodTagDao, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FoodTagMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]food_tag_repo.FoodTagDao, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Where(options.query)
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
func (obj *_FoodTagMgr) GetFromActive(active int8) (results []*food_tag_repo.FoodTagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Where("`active` = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量查找 表示数据是否处于可用状态， active = 1 可用，active=0不可用，操作可逆转
func (obj *_FoodTagMgr) GetBatchFromActive(actives []int8) (results []*food_tag_repo.FoodTagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Where("`active` IN (?)", actives).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容 是否删除 1 删除 0 未删除
func (obj *_FoodTagMgr) GetFromFlag(flag int8) (results []*food_tag_repo.FoodTagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找 是否删除 1 删除 0 未删除
func (obj *_FoodTagMgr) GetBatchFromFlag(flags []int8) (results []*food_tag_repo.FoodTagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 记录写入时间
func (obj *_FoodTagMgr) GetFromCreateTime(createTime time.Time) (results []*food_tag_repo.FoodTagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 记录写入时间
func (obj *_FoodTagMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*food_tag_repo.FoodTagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 记录更新时间
func (obj *_FoodTagMgr) GetFromUpdateTime(updateTime time.Time) (results []*food_tag_repo.FoodTagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 记录更新时间
func (obj *_FoodTagMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*food_tag_repo.FoodTagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromID 通过id获取内容 主键
func (obj *_FoodTagMgr) GetFromID(id uint64) (result food_tag_repo.FoodTagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FoodTagMgr) GetBatchFromID(ids []uint64) (results []*food_tag_repo.FoodTagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromFoodID 通过food_id获取内容 名称
func (obj *_FoodTagMgr) GetFromFoodID(foodID int64) (results []*food_tag_repo.FoodTagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Where("`food_id` = ?", foodID).Find(&results).Error

	return
}

// GetBatchFromFoodID 批量查找 名称
func (obj *_FoodTagMgr) GetBatchFromFoodID(foodIDs []int64) (results []*food_tag_repo.FoodTagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Where("`food_id` IN (?)", foodIDs).Find(&results).Error

	return
}

// GetFromTagID 通过tag_id获取内容
func (obj *_FoodTagMgr) GetFromTagID(tagID int64) (results []*food_tag_repo.FoodTagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Where("`tag_id` = ?", tagID).Find(&results).Error

	return
}

// GetBatchFromTagID 批量查找
func (obj *_FoodTagMgr) GetBatchFromTagID(tagIDs []int64) (results []*food_tag_repo.FoodTagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Where("`tag_id` IN (?)", tagIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_FoodTagMgr) GetFromType(_type string) (results []*food_tag_repo.FoodTagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_FoodTagMgr) GetBatchFromType(_types []string) (results []*food_tag_repo.FoodTagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FoodTagMgr) FetchByPrimaryKey(id uint64) (result food_tag_repo.FoodTagDao, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(food_tag_repo.FoodTagDao{}).Where("`id` = ?", id).First(&result).Error

	return
}
