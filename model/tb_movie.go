package model

import (
	"time"
)

// TbMovie [...]
type TbMovie struct {
	ID          int       `gorm:"primaryKey;column:id" json:"id"`
	Code        string    `gorm:"column:code" json:"code"`
	Title       string    `gorm:"column:title" json:"title"`
	CoverCross  string    `gorm:"column:cover_cross" json:"cover_cross"`
	CoverShaft  string    `gorm:"column:cover_shaft" json:"cover_shaft"`
	ReleaseDate time.Time `gorm:"column:release_date" json:"release_date"`
	Enable      bool      `gorm:"column:enable" json:"enable"`
	CreateAt    time.Time `gorm:"column:create_at" json:"create_at"`
}

// TableName get sql table name.获取数据库表名
func (m *TbMovie) TableName() string {
	return "tb_movie"
}

// TbMovieColumns get sql column name.获取数据库列名
var TbMovieColumns = struct {
	ID          string
	Code        string
	Title       string
	CoverCross  string
	CoverShaft  string
	ReleaseDate string
	Enable      string
	CreateAt    string
}{
	ID:          "id",
	Code:        "code",
	Title:       "title",
	CoverCross:  "cover_cross",
	CoverShaft:  "cover_shaft",
	ReleaseDate: "release_date",
	Enable:      "enable",
	CreateAt:    "create_at",
}
