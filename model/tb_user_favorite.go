package model

import (
	"time"
)

// TbUserFavorite [...]
type TbUserFavorite struct {
	UserID   int       `gorm:"column:user_id" json:"user_id"`
	TbUser   TbUser    `gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"tb_user_list"`
	MovieID  int       `gorm:"column:movie_id" json:"movie_id"`
	TbMovie  TbMovie   `gorm:"joinForeignKey:movie_id;foreignKey:id;references:MovieID" json:"tb_movie_list"`
	CreateAt time.Time `gorm:"column:create_at" json:"create_at"`
}

// TableName get sql table name.获取数据库表名
func (m *TbUserFavorite) TableName() string {
	return "tb_user_favorite"
}

// TbUserFavoriteColumns get sql column name.获取数据库列名
var TbUserFavoriteColumns = struct {
	UserID   string
	MovieID  string
	CreateAt string
}{
	UserID:   "user_id",
	MovieID:  "movie_id",
	CreateAt: "create_at",
}
