package model

// TbMovieActor [...]
type TbMovieActor struct {
	MovieID int     `gorm:"column:movie_id" json:"movie_id"`
	TbMovie TbMovie `gorm:"joinForeignKey:movie_id;foreignKey:id;references:MovieID" json:"tb_movie_list"`
	ActorID int     `gorm:"column:actor_id" json:"actor_id"`
	TbActor TbActor `gorm:"joinForeignKey:actor_id;foreignKey:id;references:ActorID" json:"tb_actor_list"`
}

// TableName get sql table name.获取数据库表名
func (m *TbMovieActor) TableName() string {
	return "tb_movie_actor"
}

// TbMovieActorColumns get sql column name.获取数据库列名
var TbMovieActorColumns = struct {
	MovieID string
	ActorID string
}{
	MovieID: "movie_id",
	ActorID: "actor_id",
}
