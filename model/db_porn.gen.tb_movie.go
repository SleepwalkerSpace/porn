package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _TbMovieMgr struct {
	*_BaseMgr
}

// TbMovieMgr open func
func TbMovieMgr(db *gorm.DB) *_TbMovieMgr {
	if db == nil {
		panic(fmt.Errorf("TbMovieMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TbMovieMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tb_movie"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TbMovieMgr) GetTableName() string {
	return "tb_movie"
}

// Reset 重置gorm会话
func (obj *_TbMovieMgr) Reset() *_TbMovieMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TbMovieMgr) Get() (result TbMovie, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovie{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TbMovieMgr) Gets() (results []*TbMovie, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TbMovieMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_TbMovieMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取
func (obj *_TbMovieMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithTitle title获取
func (obj *_TbMovieMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithCoverCross cover_cross获取
func (obj *_TbMovieMgr) WithCoverCross(coverCross string) Option {
	return optionFunc(func(o *options) { o.query["cover_cross"] = coverCross })
}

// WithCoverShaft cover_shaft获取
func (obj *_TbMovieMgr) WithCoverShaft(coverShaft string) Option {
	return optionFunc(func(o *options) { o.query["cover_shaft"] = coverShaft })
}

// WithReleaseDate release_date获取
func (obj *_TbMovieMgr) WithReleaseDate(releaseDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["release_date"] = releaseDate })
}

// WithEnable enable获取
func (obj *_TbMovieMgr) WithEnable(enable bool) Option {
	return optionFunc(func(o *options) { o.query["enable"] = enable })
}

// WithCreateAt create_at获取
func (obj *_TbMovieMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_TbMovieMgr) GetByOption(opts ...Option) (result TbMovie, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TbMovieMgr) GetByOptions(opts ...Option) (results []*TbMovie, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TbMovieMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TbMovie, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Where(options.query)
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
func (obj *_TbMovieMgr) GetFromID(id int) (result TbMovie, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_TbMovieMgr) GetBatchFromID(ids []int) (results []*TbMovie, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容
func (obj *_TbMovieMgr) GetFromCode(code string) (results []*TbMovie, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找
func (obj *_TbMovieMgr) GetBatchFromCode(codes []string) (results []*TbMovie, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容
func (obj *_TbMovieMgr) GetFromTitle(title string) (results []*TbMovie, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找
func (obj *_TbMovieMgr) GetBatchFromTitle(titles []string) (results []*TbMovie, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromCoverCross 通过cover_cross获取内容
func (obj *_TbMovieMgr) GetFromCoverCross(coverCross string) (results []*TbMovie, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Where("`cover_cross` = ?", coverCross).Find(&results).Error

	return
}

// GetBatchFromCoverCross 批量查找
func (obj *_TbMovieMgr) GetBatchFromCoverCross(coverCrosss []string) (results []*TbMovie, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Where("`cover_cross` IN (?)", coverCrosss).Find(&results).Error

	return
}

// GetFromCoverShaft 通过cover_shaft获取内容
func (obj *_TbMovieMgr) GetFromCoverShaft(coverShaft string) (results []*TbMovie, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Where("`cover_shaft` = ?", coverShaft).Find(&results).Error

	return
}

// GetBatchFromCoverShaft 批量查找
func (obj *_TbMovieMgr) GetBatchFromCoverShaft(coverShafts []string) (results []*TbMovie, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Where("`cover_shaft` IN (?)", coverShafts).Find(&results).Error

	return
}

// GetFromReleaseDate 通过release_date获取内容
func (obj *_TbMovieMgr) GetFromReleaseDate(releaseDate time.Time) (results []*TbMovie, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Where("`release_date` = ?", releaseDate).Find(&results).Error

	return
}

// GetBatchFromReleaseDate 批量查找
func (obj *_TbMovieMgr) GetBatchFromReleaseDate(releaseDates []time.Time) (results []*TbMovie, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Where("`release_date` IN (?)", releaseDates).Find(&results).Error

	return
}

// GetFromEnable 通过enable获取内容
func (obj *_TbMovieMgr) GetFromEnable(enable bool) (results []*TbMovie, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Where("`enable` = ?", enable).Find(&results).Error

	return
}

// GetBatchFromEnable 批量查找
func (obj *_TbMovieMgr) GetBatchFromEnable(enables []bool) (results []*TbMovie, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Where("`enable` IN (?)", enables).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_TbMovieMgr) GetFromCreateAt(createAt time.Time) (results []*TbMovie, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Where("`create_at` = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量查找
func (obj *_TbMovieMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*TbMovie, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Where("`create_at` IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TbMovieMgr) FetchByPrimaryKey(id int) (result TbMovie, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovie{}).Where("`id` = ?", id).First(&result).Error

	return
}
