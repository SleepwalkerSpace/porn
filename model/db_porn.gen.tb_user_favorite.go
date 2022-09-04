package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _TbUserFavoriteMgr struct {
	*_BaseMgr
}

// TbUserFavoriteMgr open func
func TbUserFavoriteMgr(db *gorm.DB) *_TbUserFavoriteMgr {
	if db == nil {
		panic(fmt.Errorf("TbUserFavoriteMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TbUserFavoriteMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tb_user_favorite"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TbUserFavoriteMgr) GetTableName() string {
	return "tb_user_favorite"
}

// Reset 重置gorm会话
func (obj *_TbUserFavoriteMgr) Reset() *_TbUserFavoriteMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TbUserFavoriteMgr) Get() (result TbUserFavorite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUserFavorite{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("tb_user").Where("id = ?", result.UserID).Find(&result.TbUser).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("tb_movie").Where("id = ?", result.MovieID).Find(&result.TbMovie).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_TbUserFavoriteMgr) Gets() (results []*TbUserFavorite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUserFavorite{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_user").Where("id = ?", results[i].UserID).Find(&results[i].TbUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TbUserFavoriteMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TbUserFavorite{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithUserID user_id获取
func (obj *_TbUserFavoriteMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithMovieID movie_id获取
func (obj *_TbUserFavoriteMgr) WithMovieID(movieID int) Option {
	return optionFunc(func(o *options) { o.query["movie_id"] = movieID })
}

// WithCreateAt create_at获取
func (obj *_TbUserFavoriteMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_TbUserFavoriteMgr) GetByOption(opts ...Option) (result TbUserFavorite, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TbUserFavorite{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("tb_user").Where("id = ?", result.UserID).Find(&result.TbUser).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("tb_movie").Where("id = ?", result.MovieID).Find(&result.TbMovie).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TbUserFavoriteMgr) GetByOptions(opts ...Option) (results []*TbUserFavorite, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TbUserFavorite{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_user").Where("id = ?", results[i].UserID).Find(&results[i].TbUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_TbUserFavoriteMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TbUserFavorite, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TbUserFavorite{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_user").Where("id = ?", results[i].UserID).Find(&results[i].TbUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
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

// GetFromUserID 通过user_id获取内容
func (obj *_TbUserFavoriteMgr) GetFromUserID(userID int) (results []*TbUserFavorite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUserFavorite{}).Where("`user_id` = ?", userID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_user").Where("id = ?", results[i].UserID).Find(&results[i].TbUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUserID 批量查找
func (obj *_TbUserFavoriteMgr) GetBatchFromUserID(userIDs []int) (results []*TbUserFavorite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUserFavorite{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_user").Where("id = ?", results[i].UserID).Find(&results[i].TbUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromMovieID 通过movie_id获取内容
func (obj *_TbUserFavoriteMgr) GetFromMovieID(movieID int) (results []*TbUserFavorite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUserFavorite{}).Where("`movie_id` = ?", movieID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_user").Where("id = ?", results[i].UserID).Find(&results[i].TbUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromMovieID 批量查找
func (obj *_TbUserFavoriteMgr) GetBatchFromMovieID(movieIDs []int) (results []*TbUserFavorite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUserFavorite{}).Where("`movie_id` IN (?)", movieIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_user").Where("id = ?", results[i].UserID).Find(&results[i].TbUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_TbUserFavoriteMgr) GetFromCreateAt(createAt time.Time) (results []*TbUserFavorite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUserFavorite{}).Where("`create_at` = ?", createAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_user").Where("id = ?", results[i].UserID).Find(&results[i].TbUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateAt 批量查找
func (obj *_TbUserFavoriteMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*TbUserFavorite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUserFavorite{}).Where("`create_at` IN (?)", createAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_user").Where("id = ?", results[i].UserID).Find(&results[i].TbUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchIndexByUserID  获取多个内容
func (obj *_TbUserFavoriteMgr) FetchIndexByUserID(userID int) (results []*TbUserFavorite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUserFavorite{}).Where("`user_id` = ?", userID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_user").Where("id = ?", results[i].UserID).Find(&results[i].TbUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByMovieID  获取多个内容
func (obj *_TbUserFavoriteMgr) FetchIndexByMovieID(movieID int) (results []*TbUserFavorite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUserFavorite{}).Where("`movie_id` = ?", movieID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_user").Where("id = ?", results[i].UserID).Find(&results[i].TbUser).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
