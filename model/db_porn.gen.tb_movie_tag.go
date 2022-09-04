package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TbMovieTagMgr struct {
	*_BaseMgr
}

// TbMovieTagMgr open func
func TbMovieTagMgr(db *gorm.DB) *_TbMovieTagMgr {
	if db == nil {
		panic(fmt.Errorf("TbMovieTagMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TbMovieTagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tb_movie_tag"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TbMovieTagMgr) GetTableName() string {
	return "tb_movie_tag"
}

// Reset 重置gorm会话
func (obj *_TbMovieTagMgr) Reset() *_TbMovieTagMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TbMovieTagMgr) Get() (result TbMovieTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovieTag{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("tb_movie").Where("id = ?", result.MovieID).Find(&result.TbMovie).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("tb_tag").Where("id = ?", result.TagID).Find(&result.TbTag).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_TbMovieTagMgr) Gets() (results []*TbMovieTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovieTag{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_tag").Where("id = ?", results[i].TagID).Find(&results[i].TbTag).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TbMovieTagMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TbMovieTag{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithMovieID movie_id获取
func (obj *_TbMovieTagMgr) WithMovieID(movieID int) Option {
	return optionFunc(func(o *options) { o.query["movie_id"] = movieID })
}

// WithTagID tag_id获取
func (obj *_TbMovieTagMgr) WithTagID(tagID int) Option {
	return optionFunc(func(o *options) { o.query["tag_id"] = tagID })
}

// GetByOption 功能选项模式获取
func (obj *_TbMovieTagMgr) GetByOption(opts ...Option) (result TbMovieTag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TbMovieTag{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("tb_movie").Where("id = ?", result.MovieID).Find(&result.TbMovie).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("tb_tag").Where("id = ?", result.TagID).Find(&result.TbTag).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TbMovieTagMgr) GetByOptions(opts ...Option) (results []*TbMovieTag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TbMovieTag{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_tag").Where("id = ?", results[i].TagID).Find(&results[i].TbTag).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_TbMovieTagMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TbMovieTag, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TbMovieTag{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_tag").Where("id = ?", results[i].TagID).Find(&results[i].TbTag).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromMovieID 通过movie_id获取内容
func (obj *_TbMovieTagMgr) GetFromMovieID(movieID int) (results []*TbMovieTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovieTag{}).Where("`movie_id` = ?", movieID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_tag").Where("id = ?", results[i].TagID).Find(&results[i].TbTag).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromMovieID 批量查找
func (obj *_TbMovieTagMgr) GetBatchFromMovieID(movieIDs []int) (results []*TbMovieTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovieTag{}).Where("`movie_id` IN (?)", movieIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_tag").Where("id = ?", results[i].TagID).Find(&results[i].TbTag).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTagID 通过tag_id获取内容
func (obj *_TbMovieTagMgr) GetFromTagID(tagID int) (results []*TbMovieTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovieTag{}).Where("`tag_id` = ?", tagID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_tag").Where("id = ?", results[i].TagID).Find(&results[i].TbTag).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromTagID 批量查找
func (obj *_TbMovieTagMgr) GetBatchFromTagID(tagIDs []int) (results []*TbMovieTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovieTag{}).Where("`tag_id` IN (?)", tagIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_tag").Where("id = ?", results[i].TagID).Find(&results[i].TbTag).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchIndexByMovieID  获取多个内容
func (obj *_TbMovieTagMgr) FetchIndexByMovieID(movieID int) (results []*TbMovieTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovieTag{}).Where("`movie_id` = ?", movieID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_tag").Where("id = ?", results[i].TagID).Find(&results[i].TbTag).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByTagID  获取多个内容
func (obj *_TbMovieTagMgr) FetchIndexByTagID(tagID int) (results []*TbMovieTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovieTag{}).Where("`tag_id` = ?", tagID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_tag").Where("id = ?", results[i].TagID).Find(&results[i].TbTag).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
