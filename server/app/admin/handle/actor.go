package handle

import (
	"fmt"
	"net/http"
	"porn/common/httpenc"
	"porn/common/persistence"
	"porn/common/sqlenc"
	"porn/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ActorListParam struct {
	Pagination httpenc.Pagination `json:"pagination"`
	SortType   int                `json:"sort_type"`
}

type ActorInfo struct {
	ID            int    `json:"id"`
	StageNameJp   string `json:"stage_name_jp"`
	StageNameCn   string `json:"stage_name_cn"`
	Portrait      string `json:"portrait"`
	MovieCount    int64  `json:"movie_count"`
	ViewCount     int64  `json:"view_count"`
	FavoriteCount int64  `json:"favorite_count"`
}

type ActorListData struct {
	Pagination httpenc.Pagination `json:"pagination"`
	Infos      []ActorInfo        `json:"infos"`
	Total      int64              `json:"total"`
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
		sql = sqlenc.ActorListByMovieCount()
	case 2:
		sql = sqlenc.ActorListByMovieView()
	case 3:
		sql = sqlenc.ActorListByMovieFavorite()

	default:
		sql = sqlenc.ActorListByMovieCount()
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
	data.Pagination = param.Pagination
	data.Total = param.Pagination.Total
	c.JSON(http.StatusOK, data)
}

type ActorCreateParam struct {
	StageNameJp string `json:"stage_name_jp"`
	StageNameCn string `json:"stage_name_cn"`
	Portrait    string `json:"portrait"`
}

func ActorCreate(c *gin.Context) {
	var param ActorCreateParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusOK, httpenc.MakeFailHTTPRespPackage(httpenc.HTTPRespCode_ParamUnmarshal))
		return
	}
	mgr := model.TbActorMgr(persistence.DB)
	if _, err := mgr.GetByOption(mgr.WithStageNameJp(param.StageNameJp)); err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, httpenc.MakeFailHTTPRespPackage(httpenc.HTTPRespCode_Repeated))
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	val := model.TbActor{
		StageNameJp: param.StageNameJp,
		StageNameCn: param.StageNameCn,
		Portrait:    param.Portrait,
	}
	if err := mgr.Create(&val).Error; err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, httpenc.MakeSuccHTTPRespPackage(nil))
}

type ActorUpdateParam struct {
	ID          int    `json:"id"`
	StageNameJp string `json:"stage_name_jp"`
	StageNameCn string `json:"stage_name_cn"`
	Portrait    string `json:"portrait"`
}

func ActorUpdate(c *gin.Context) {
	var param ActorUpdateParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusOK, httpenc.MakeFailHTTPRespPackage(httpenc.HTTPRespCode_ParamUnmarshal))
		return
	}
	mgr := model.TbActorMgr(persistence.DB)
	val, err := mgr.FetchByPrimaryKey(param.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, httpenc.MakeFailHTTPRespPackage(httpenc.HTTPRespCode_Repeated))
			return
		} else {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
	}
	val.StageNameJp = param.StageNameJp
	val.StageNameCn = param.StageNameCn
	val.Portrait = param.Portrait
	if err := mgr.Save(&val).Error; err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, httpenc.MakeSuccHTTPRespPackage(nil))
}

type ActorDeleteParam struct {
	ID int `json:"id"`
}

func ActorDelete(c *gin.Context) {
	var param ActorDeleteParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusOK, httpenc.MakeFailHTTPRespPackage(httpenc.HTTPRespCode_ParamUnmarshal))
		return
	}
	mgr := model.TbActorMgr(persistence.DB)
	val, err := mgr.FetchByPrimaryKey(param.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, httpenc.MakeFailHTTPRespPackage(httpenc.HTTPRespCode_Repeated))
			return
		} else {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
	}
	if err := model.TbMovieActorMgr(persistence.DB).Delete(&model.TbMovieActor{}, fmt.Sprintf("%s = ?", model.TbMovieActorColumns.ActorID), val.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if err := mgr.Delete(&val).Error; err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, httpenc.MakeSuccHTTPRespPackage(nil))
}
