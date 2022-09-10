package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TbTagMgr struct {
	*_BaseMgr
}

// TbTagMgr open func
func TbTagMgr(db *gorm.DB) *_TbTagMgr {
	if db == nil {
		panic(fmt.Errorf("TbTagMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TbTagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tb_tag"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TbTagMgr) GetTableName() string {
	return "tb_tag"
}

// Reset 重置gorm会话
func (obj *_TbTagMgr) Reset() *_TbTagMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TbTagMgr) Get() (result TbTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbTag{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TbTagMgr) Gets() (results []*TbTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbTag{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TbTagMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TbTag{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_TbTagMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTitel titel获取
func (obj *_TbTagMgr) WithTitel(titel string) Option {
	return optionFunc(func(o *options) { o.query["titel"] = titel })
}

// WithCover cover获取
func (obj *_TbTagMgr) WithCover(cover string) Option {
	return optionFunc(func(o *options) { o.query["cover"] = cover })
}

// GetByOption 功能选项模式获取
func (obj *_TbTagMgr) GetByOption(opts ...Option) (result TbTag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TbTag{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TbTagMgr) GetByOptions(opts ...Option) (results []*TbTag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TbTag{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TbTagMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TbTag, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TbTag{}).Where(options.query)
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

// GetFromID 通过id获取内容
func (obj *_TbTagMgr) GetFromID(id int) (result TbTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbTag{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_TbTagMgr) GetBatchFromID(ids []int) (results []*TbTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbTag{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTitel 通过titel获取内容
func (obj *_TbTagMgr) GetFromTitel(titel string) (results []*TbTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbTag{}).Where("`titel` = ?", titel).Find(&results).Error

	return
}

// GetBatchFromTitel 批量查找
func (obj *_TbTagMgr) GetBatchFromTitel(titels []string) (results []*TbTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbTag{}).Where("`titel` IN (?)", titels).Find(&results).Error

	return
}

// GetFromCover 通过cover获取内容
func (obj *_TbTagMgr) GetFromCover(cover string) (results []*TbTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbTag{}).Where("`cover` = ?", cover).Find(&results).Error

	return
}

// GetBatchFromCover 批量查找
func (obj *_TbTagMgr) GetBatchFromCover(covers []string) (results []*TbTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbTag{}).Where("`cover` IN (?)", covers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TbTagMgr) FetchByPrimaryKey(id int) (result TbTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbTag{}).Where("`id` = ?", id).First(&result).Error

	return
}
