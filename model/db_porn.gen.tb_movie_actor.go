package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TbMovieActorMgr struct {
	*_BaseMgr
}

// TbMovieActorMgr open func
func TbMovieActorMgr(db *gorm.DB) *_TbMovieActorMgr {
	if db == nil {
		panic(fmt.Errorf("TbMovieActorMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TbMovieActorMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tb_movie_actor"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TbMovieActorMgr) GetTableName() string {
	return "tb_movie_actor"
}

// Reset 重置gorm会话
func (obj *_TbMovieActorMgr) Reset() *_TbMovieActorMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TbMovieActorMgr) Get() (result TbMovieActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovieActor{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("tb_movie").Where("id = ?", result.MovieID).Find(&result.TbMovie).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("tb_actor").Where("id = ?", result.ActorID).Find(&result.TbActor).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_TbMovieActorMgr) Gets() (results []*TbMovieActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovieActor{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_actor").Where("id = ?", results[i].ActorID).Find(&results[i].TbActor).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TbMovieActorMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TbMovieActor{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithMovieID movie_id获取
func (obj *_TbMovieActorMgr) WithMovieID(movieID int) Option {
	return optionFunc(func(o *options) { o.query["movie_id"] = movieID })
}

// WithActorID actor_id获取
func (obj *_TbMovieActorMgr) WithActorID(actorID int) Option {
	return optionFunc(func(o *options) { o.query["actor_id"] = actorID })
}

// GetByOption 功能选项模式获取
func (obj *_TbMovieActorMgr) GetByOption(opts ...Option) (result TbMovieActor, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TbMovieActor{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("tb_movie").Where("id = ?", result.MovieID).Find(&result.TbMovie).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("tb_actor").Where("id = ?", result.ActorID).Find(&result.TbActor).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TbMovieActorMgr) GetByOptions(opts ...Option) (results []*TbMovieActor, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TbMovieActor{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_actor").Where("id = ?", results[i].ActorID).Find(&results[i].TbActor).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_TbMovieActorMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TbMovieActor, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TbMovieActor{}).Where(options.query)
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
			if err = obj.NewDB().Table("tb_actor").Where("id = ?", results[i].ActorID).Find(&results[i].TbActor).Error; err != nil { //
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
func (obj *_TbMovieActorMgr) GetFromMovieID(movieID int) (results []*TbMovieActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovieActor{}).Where("`movie_id` = ?", movieID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_actor").Where("id = ?", results[i].ActorID).Find(&results[i].TbActor).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromMovieID 批量查找
func (obj *_TbMovieActorMgr) GetBatchFromMovieID(movieIDs []int) (results []*TbMovieActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovieActor{}).Where("`movie_id` IN (?)", movieIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_actor").Where("id = ?", results[i].ActorID).Find(&results[i].TbActor).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromActorID 通过actor_id获取内容
func (obj *_TbMovieActorMgr) GetFromActorID(actorID int) (results []*TbMovieActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovieActor{}).Where("`actor_id` = ?", actorID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_actor").Where("id = ?", results[i].ActorID).Find(&results[i].TbActor).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromActorID 批量查找
func (obj *_TbMovieActorMgr) GetBatchFromActorID(actorIDs []int) (results []*TbMovieActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovieActor{}).Where("`actor_id` IN (?)", actorIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_actor").Where("id = ?", results[i].ActorID).Find(&results[i].TbActor).Error; err != nil { //
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
func (obj *_TbMovieActorMgr) FetchIndexByMovieID(movieID int) (results []*TbMovieActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovieActor{}).Where("`movie_id` = ?", movieID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_actor").Where("id = ?", results[i].ActorID).Find(&results[i].TbActor).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByActorID  获取多个内容
func (obj *_TbMovieActorMgr) FetchIndexByActorID(actorID int) (results []*TbMovieActor, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbMovieActor{}).Where("`actor_id` = ?", actorID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tb_movie").Where("id = ?", results[i].MovieID).Find(&results[i].TbMovie).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("tb_actor").Where("id = ?", results[i].ActorID).Find(&results[i].TbActor).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
