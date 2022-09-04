package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TbActorMgr struct {
	*_BaseMgr
}

// TbActorMgr open func
func TbActorMgr(db *gorm.DB) *_TbActorMgr {
	if db == nil {
		panic(fmt.Errorf("TbActorMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TbActorMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tb_actor"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TbActorMgr) GetTableName() string {
	return "tb_actor"
}

// Reset 重置gorm会话
func (obj *_TbActorMgr) Reset() *_TbActorMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TbActorMgr) Get() (result TbActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbActor{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TbActorMgr) Gets() (results []*TbActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbActor{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TbActorMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TbActor{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_TbActorMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithStageNameJp stage_name_jp获取
func (obj *_TbActorMgr) WithStageNameJp(stageNameJp string) Option {
	return optionFunc(func(o *options) { o.query["stage_name_jp"] = stageNameJp })
}

// WithStageNameCn stage_name_cn获取
func (obj *_TbActorMgr) WithStageNameCn(stageNameCn string) Option {
	return optionFunc(func(o *options) { o.query["stage_name_cn"] = stageNameCn })
}

// WithPortrait portrait获取
func (obj *_TbActorMgr) WithPortrait(portrait string) Option {
	return optionFunc(func(o *options) { o.query["portrait"] = portrait })
}

// GetByOption 功能选项模式获取
func (obj *_TbActorMgr) GetByOption(opts ...Option) (result TbActor, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TbActor{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TbActorMgr) GetByOptions(opts ...Option) (results []*TbActor, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TbActor{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TbActorMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TbActor, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TbActor{}).Where(options.query)
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
func (obj *_TbActorMgr) GetFromID(id int) (result TbActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbActor{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_TbActorMgr) GetBatchFromID(ids []int) (results []*TbActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbActor{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromStageNameJp 通过stage_name_jp获取内容
func (obj *_TbActorMgr) GetFromStageNameJp(stageNameJp string) (results []*TbActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbActor{}).Where("`stage_name_jp` = ?", stageNameJp).Find(&results).Error

	return
}

// GetBatchFromStageNameJp 批量查找
func (obj *_TbActorMgr) GetBatchFromStageNameJp(stageNameJps []string) (results []*TbActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbActor{}).Where("`stage_name_jp` IN (?)", stageNameJps).Find(&results).Error

	return
}

// GetFromStageNameCn 通过stage_name_cn获取内容
func (obj *_TbActorMgr) GetFromStageNameCn(stageNameCn string) (results []*TbActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbActor{}).Where("`stage_name_cn` = ?", stageNameCn).Find(&results).Error

	return
}

// GetBatchFromStageNameCn 批量查找
func (obj *_TbActorMgr) GetBatchFromStageNameCn(stageNameCns []string) (results []*TbActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbActor{}).Where("`stage_name_cn` IN (?)", stageNameCns).Find(&results).Error

	return
}

// GetFromPortrait 通过portrait获取内容
func (obj *_TbActorMgr) GetFromPortrait(portrait string) (results []*TbActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbActor{}).Where("`portrait` = ?", portrait).Find(&results).Error

	return
}

// GetBatchFromPortrait 批量查找
func (obj *_TbActorMgr) GetBatchFromPortrait(portraits []string) (results []*TbActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbActor{}).Where("`portrait` IN (?)", portraits).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TbActorMgr) FetchByPrimaryKey(id int) (result TbActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbActor{}).Where("`id` = ?", id).First(&result).Error

	return
}
