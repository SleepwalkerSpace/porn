package model

import (
	"time"
)

// TbUser [...]
type TbUser struct {
	ID            int       `gorm:"primaryKey;column:id" json:"id"`
	IP            string    `gorm:"column:ip" json:"ip"`
	Nickname      string    `gorm:"column:nickname" json:"nickname"`
	DevieName     string    `gorm:"column:devie_name" json:"devie_name"`
	DevieCode     string    `gorm:"column:devie_code" json:"devie_code"`
	Email         string    `gorm:"column:email" json:"email"`
	EmailVerify   bool      `gorm:"column:email_verify" json:"email_verify"`
	EmailVerifyAt time.Time `gorm:"column:email_verify_at" json:"email_verify_at"`
	Enable        bool      `gorm:"column:enable" json:"enable"`
	CreateAt      time.Time `gorm:"column:create_at" json:"create_at"`
}

// TableName get sql table name.获取数据库表名
func (m *TbUser) TableName() string {
	return "tb_user"
}

// TbUserColumns get sql column name.获取数据库列名
var TbUserColumns = struct {
	ID            string
	IP            string
	Nickname      string
	DevieName     string
	DevieCode     string
	Email         string
	EmailVerify   string
	EmailVerifyAt string
	Enable        string
	CreateAt      string
}{
	ID:            "id",
	IP:            "ip",
	Nickname:      "nickname",
	DevieName:     "devie_name",
	DevieCode:     "devie_code",
	Email:         "email",
	EmailVerify:   "email_verify",
	EmailVerifyAt: "email_verify_at",
	Enable:        "enable",
	CreateAt:      "create_at",
}
