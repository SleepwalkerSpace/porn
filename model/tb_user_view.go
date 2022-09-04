package model

import (
	"time"
)

// TbUserView [...]
type TbUserView struct {
	UserID   int       `gorm:"column:user_id" json:"user_id"`
	TbUser   TbUser    `gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"tb_user_list"`
	MovieID  int       `gorm:"column:movie_id" json:"movie_id"`
	TbMovie  TbMovie   `gorm:"joinForeignKey:movie_id;foreignKey:id;references:MovieID" json:"tb_movie_list"`
	CreateAt time.Time `gorm:"column:create_at" json:"create_at"`
}

// TableName get sql table name.获取数据库表名
func (m *TbUserView) TableName() string {
	return "tb_user_view"
}

// TbUserViewColumns get sql column name.获取数据库列名
var TbUserViewColumns = struct {
	UserID   string
	MovieID  string
	CreateAt string
}{
	UserID:   "user_id",
	MovieID:  "movie_id",
	CreateAt: "create_at",
}
