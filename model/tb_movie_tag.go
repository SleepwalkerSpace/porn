package model

// TbMovieTag [...]
type TbMovieTag struct {
	MovieID int     `gorm:"column:movie_id" json:"movie_id"`
	TbMovie TbMovie `gorm:"joinForeignKey:movie_id;foreignKey:id;references:MovieID" json:"tb_movie_list"`
	TagID   int     `gorm:"column:tag_id" json:"tag_id"`
	TbTag   TbTag   `gorm:"joinForeignKey:tag_id;foreignKey:id;references:TagID" json:"tb_tag_list"`
}

// TableName get sql table name.获取数据库表名
func (m *TbMovieTag) TableName() string {
	return "tb_movie_tag"
}

// TbMovieTagColumns get sql column name.获取数据库列名
var TbMovieTagColumns = struct {
	MovieID string
	TagID   string
}{
	MovieID: "movie_id",
	TagID:   "tag_id",
}
