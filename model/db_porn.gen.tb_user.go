package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _TbUserMgr struct {
	*_BaseMgr
}

// TbUserMgr open func
func TbUserMgr(db *gorm.DB) *_TbUserMgr {
	if db == nil {
		panic(fmt.Errorf("TbUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TbUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tb_user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TbUserMgr) GetTableName() string {
	return "tb_user"
}

// Reset 重置gorm会话
func (obj *_TbUserMgr) Reset() *_TbUserMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TbUserMgr) Get() (result TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TbUserMgr) Gets() (results []*TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TbUserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TbUser{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_TbUserMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithIP ip获取
func (obj *_TbUserMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["ip"] = ip })
}

// WithNickname nickname获取
func (obj *_TbUserMgr) WithNickname(nickname string) Option {
	return optionFunc(func(o *options) { o.query["nickname"] = nickname })
}

// WithDevieName devie_name获取
func (obj *_TbUserMgr) WithDevieName(devieName string) Option {
	return optionFunc(func(o *options) { o.query["devie_name"] = devieName })
}

// WithDevieCode devie_code获取
func (obj *_TbUserMgr) WithDevieCode(devieCode string) Option {
	return optionFunc(func(o *options) { o.query["devie_code"] = devieCode })
}

// WithEmail email获取
func (obj *_TbUserMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithEmailVerify email_verify获取
func (obj *_TbUserMgr) WithEmailVerify(emailVerify bool) Option {
	return optionFunc(func(o *options) { o.query["email_verify"] = emailVerify })
}

// WithEmailVerifyAt email_verify_at获取
func (obj *_TbUserMgr) WithEmailVerifyAt(emailVerifyAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["email_verify_at"] = emailVerifyAt })
}

// WithEnable enable获取
func (obj *_TbUserMgr) WithEnable(enable bool) Option {
	return optionFunc(func(o *options) { o.query["enable"] = enable })
}

// WithCreateAt create_at获取
func (obj *_TbUserMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_TbUserMgr) GetByOption(opts ...Option) (result TbUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TbUserMgr) GetByOptions(opts ...Option) (results []*TbUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TbUserMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TbUser, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where(options.query)
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
func (obj *_TbUserMgr) GetFromID(id int) (result TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_TbUserMgr) GetBatchFromID(ids []int) (results []*TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromIP 通过ip获取内容
func (obj *_TbUserMgr) GetFromIP(ip string) (results []*TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where("`ip` = ?", ip).Find(&results).Error

	return
}

// GetBatchFromIP 批量查找
func (obj *_TbUserMgr) GetBatchFromIP(ips []string) (results []*TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where("`ip` IN (?)", ips).Find(&results).Error

	return
}

// GetFromNickname 通过nickname获取内容
func (obj *_TbUserMgr) GetFromNickname(nickname string) (results []*TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where("`nickname` = ?", nickname).Find(&results).Error

	return
}

// GetBatchFromNickname 批量查找
func (obj *_TbUserMgr) GetBatchFromNickname(nicknames []string) (results []*TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where("`nickname` IN (?)", nicknames).Find(&results).Error

	return
}

// GetFromDevieName 通过devie_name获取内容
func (obj *_TbUserMgr) GetFromDevieName(devieName string) (results []*TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where("`devie_name` = ?", devieName).Find(&results).Error

	return
}

// GetBatchFromDevieName 批量查找
func (obj *_TbUserMgr) GetBatchFromDevieName(devieNames []string) (results []*TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where("`devie_name` IN (?)", devieNames).Find(&results).Error

	return
}

// GetFromDevieCode 通过devie_code获取内容
func (obj *_TbUserMgr) GetFromDevieCode(devieCode string) (results []*TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where("`devie_code` = ?", devieCode).Find(&results).Error

	return
}

// GetBatchFromDevieCode 批量查找
func (obj *_TbUserMgr) GetBatchFromDevieCode(devieCodes []string) (results []*TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where("`devie_code` IN (?)", devieCodes).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容
func (obj *_TbUserMgr) GetFromEmail(email string) (results []*TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where("`email` = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量查找
func (obj *_TbUserMgr) GetBatchFromEmail(emails []string) (results []*TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where("`email` IN (?)", emails).Find(&results).Error

	return
}

// GetFromEmailVerify 通过email_verify获取内容
func (obj *_TbUserMgr) GetFromEmailVerify(emailVerify bool) (results []*TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where("`email_verify` = ?", emailVerify).Find(&results).Error

	return
}

// GetBatchFromEmailVerify 批量查找
func (obj *_TbUserMgr) GetBatchFromEmailVerify(emailVerifys []bool) (results []*TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where("`email_verify` IN (?)", emailVerifys).Find(&results).Error

	return
}

// GetFromEmailVerifyAt 通过email_verify_at获取内容
func (obj *_TbUserMgr) GetFromEmailVerifyAt(emailVerifyAt time.Time) (results []*TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where("`email_verify_at` = ?", emailVerifyAt).Find(&results).Error

	return
}

// GetBatchFromEmailVerifyAt 批量查找
func (obj *_TbUserMgr) GetBatchFromEmailVerifyAt(emailVerifyAts []time.Time) (results []*TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where("`email_verify_at` IN (?)", emailVerifyAts).Find(&results).Error

	return
}

// GetFromEnable 通过enable获取内容
func (obj *_TbUserMgr) GetFromEnable(enable bool) (results []*TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where("`enable` = ?", enable).Find(&results).Error

	return
}

// GetBatchFromEnable 批量查找
func (obj *_TbUserMgr) GetBatchFromEnable(enables []bool) (results []*TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where("`enable` IN (?)", enables).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_TbUserMgr) GetFromCreateAt(createAt time.Time) (results []*TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where("`create_at` = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量查找
func (obj *_TbUserMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where("`create_at` IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TbUserMgr) FetchByPrimaryKey(id int) (result TbUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TbUser{}).Where("`id` = ?", id).First(&result).Error

	return
}
