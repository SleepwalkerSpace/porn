package model

// TbTag [...]
type TbTag struct {
	ID    int    `gorm:"primaryKey;column:id" json:"id"`
	Named string `gorm:"column:named" json:"named"`
	Cover string `gorm:"column:cover" json:"cover"`
}

// TableName get sql table name.获取数据库表名
func (m *TbTag) TableName() string {
	return "tb_tag"
}

// TbTagColumns get sql column name.获取数据库列名
var TbTagColumns = struct {
	ID    string
	Named string
	Cover string
}{
	ID:    "id",
	Named: "named",
	Cover: "cover",
}
