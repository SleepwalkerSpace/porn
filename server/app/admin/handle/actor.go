package handle

import (
	"net/http"
	"porn/common/httpenc"
	"porn/common/persistence"
	"porn/common/sqlenc"
	"porn/model"

	"github.com/gin-gonic/gin"
)

type ActorListParam struct {
	Pagination httpenc.Pagination `json:"pagination"`
	SortType   int                `json:"sort_type"`
}

type ActorInfo struct {
	ID          int    `json:"id"`
	StageNameJp string `json:"stage_name_jp"`
	StageNameCn string `json:"stage_name_cn"`
	Portrait    string `json:"portrait"`
	Count       int    `json:"count"`
}

type ActorListData struct {
	Infos []ActorInfo `json:"infos"`
	Total int64       `json:"total"`
}

func ActorList(c *gin.Context) {
	var (
		param ActorListParam
		data  ActorListData
	)
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusOK, httpenc.MakeFailHTTPRespPackage(httpenc.HTTPRespCode_ParamUnmarshal))
		return
	}
	limit, offset := param.Pagination.LimitOrOffset()

	var sql string
	switch param.SortType {
	case 0, 1:
		sql = sqlenc.ActorListByMovieCount(limit, offset)
	case 2:
		sql = sqlenc.ActorListByMovieView(limit, offset)
	case 3:
		sql = sqlenc.ActorListByMovieFavorite(limit, offset)
	}
	if err := persistence.DB.Raw(sql, limit, offset).Scan(&data.Infos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if err := model.TbActorMgr(persistence.DB).Count(&param.Pagination.Total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	param.Pagination.PrevOrNext(param.Pagination.Total)
	c.JSON(http.StatusOK, data)
}
func ActorCreate(c *gin.Context) {}
func ActorUpdate(c *gin.Context) {}
func ActorDelete(c *gin.Context) {}
