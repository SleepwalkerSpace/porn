package model

// TbActor [...]
type TbActor struct {
	ID          int    `gorm:"primaryKey;column:id" json:"id"`
	StageNameJp string `gorm:"column:stage_name_jp" json:"stage_name_jp"`
	StageNameCn string `gorm:"column:stage_name_cn" json:"stage_name_cn"`
	Portrait    string `gorm:"column:portrait" json:"portrait"`
}

// TableName get sql table name.获取数据库表名
func (m *TbActor) TableName() string {
	return "tb_actor"
}

// TbActorColumns get sql column name.获取数据库列名
var TbActorColumns = struct {
	ID          string
	StageNameJp string
	StageNameCn string
	Portrait    string
}{
	ID:          "id",
	StageNameJp: "stage_name_jp",
	StageNameCn: "stage_name_cn",
	Portrait:    "portrait",
}
