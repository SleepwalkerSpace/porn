package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _TbUserViewMgr struct {
	*_BaseMgr
}

// TbUserViewMgr open func
func TbUserViewMgr(db *gorm.DB) *_TbUserViewMgr {
	if db == nil {
		panic(fmt.Errorf("TbUserViewMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TbUserViewMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tb_user_view"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TbUserViewMgr) GetTableName() string {
	return "tb_user_view"
}

// Reset 重置gorm会话
func (obj *_TbUserViewMgr) Reset() *_TbUserViewMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TbUserViewMgr) Get() (result TbUserView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUserView{}).First(&result).Error
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
func (obj *_TbUserViewMgr) Gets() (results []*TbUserView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUserView{}).Find(&results).Error
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
func (obj *_TbUserViewMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TbUserView{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithUserID user_id获取
func (obj *_TbUserViewMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithMovieID movie_id获取
func (obj *_TbUserViewMgr) WithMovieID(movieID int) Option {
	return optionFunc(func(o *options) { o.query["movie_id"] = movieID })
}

// WithCreateAt create_at获取
func (obj *_TbUserViewMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_TbUserViewMgr) GetByOption(opts ...Option) (result TbUserView, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TbUserView{}).Where(options.query).First(&result).Error
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
func (obj *_TbUserViewMgr) GetByOptions(opts ...Option) (results []*TbUserView, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TbUserView{}).Where(options.query).Find(&results).Error
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
func (obj *_TbUserViewMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TbUserView, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TbUserView{}).Where(options.query)
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
func (obj *_TbUserViewMgr) GetFromUserID(userID int) (results []*TbUserView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUserView{}).Where("`user_id` = ?", userID).Find(&results).Error
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
func (obj *_TbUserViewMgr) GetBatchFromUserID(userIDs []int) (results []*TbUserView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUserView{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error
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
func (obj *_TbUserViewMgr) GetFromMovieID(movieID int) (results []*TbUserView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUserView{}).Where("`movie_id` = ?", movieID).Find(&results).Error
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
func (obj *_TbUserViewMgr) GetBatchFromMovieID(movieIDs []int) (results []*TbUserView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUserView{}).Where("`movie_id` IN (?)", movieIDs).Find(&results).Error
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
func (obj *_TbUserViewMgr) GetFromCreateAt(createAt time.Time) (results []*TbUserView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUserView{}).Where("`create_at` = ?", createAt).Find(&results).Error
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
func (obj *_TbUserViewMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*TbUserView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUserView{}).Where("`create_at` IN (?)", createAts).Find(&results).Error
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
func (obj *_TbUserViewMgr) FetchIndexByUserID(userID int) (results []*TbUserView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUserView{}).Where("`user_id` = ?", userID).Find(&results).Error
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
func (obj *_TbUserViewMgr) FetchIndexByMovieID(movieID int) (results []*TbUserView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUserView{}).Where("`movie_id` = ?", movieID).Find(&results).Error
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
