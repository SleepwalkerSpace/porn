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

type TagInfo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Cover string `json:"cover"`
	Count int    `json:"count"`
}
type TagListData struct {
	Infos []TagInfo `json:"infos"`
	Total int64     `json:"total"`
}

func TagList(c *gin.Context) {

	var data TagListData
	if err := persistence.DB.Raw(sqlenc.TagListSql()).Scan(&data.Infos).Error; err != nil {
		c.JSON(http.StatusOK, httpenc.MakeFailHTTPRespPackage(httpenc.HTTPRespCode_ParamIllegal))
		return
	}
	if err := model.TbTagMgr(persistence.DB).Count(&data.Total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, httpenc.MakeSuccHTTPRespPackage(data))
}

type TagCreateParam struct {
	Title string `json:"title"`
	Cover string `json:"cover"`
}

func TagCreate(c *gin.Context) {
	var param TagCreateParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusOK, httpenc.MakeFailHTTPRespPackage(httpenc.HTTPRespCode_ParamUnmarshal))
		return
	}
	mgr := model.TbTagMgr(persistence.DB)
	if _, err := mgr.GetByOption(mgr.WithTitel(param.Title)); err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, httpenc.MakeFailHTTPRespPackage(httpenc.HTTPRespCode_Repeated))
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	val := model.TbTag{Titel: param.Title, Cover: param.Cover}
	if err := mgr.Create(&val).Error; err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, httpenc.MakeSuccHTTPRespPackage(nil))
}

type TagUpdateParam struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Cover string `json:"cover"`
}

func TagUpdate(c *gin.Context) {
	var param TagUpdateParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusOK, httpenc.MakeFailHTTPRespPackage(httpenc.HTTPRespCode_ParamUnmarshal))
		return
	}
	mgr := model.TbTagMgr(persistence.DB)
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

	val.Titel = param.Title
	val.Cover = param.Cover
	if err := mgr.Save(&val).Error; err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, httpenc.MakeSuccHTTPRespPackage(nil))
}

type TagDeleteParam struct {
	ID int `json:"id"`
}

func TagDelete(c *gin.Context) {
	var param TagDeleteParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusOK, httpenc.MakeFailHTTPRespPackage(httpenc.HTTPRespCode_ParamUnmarshal))
		return
	}
	mgr := model.TbTagMgr(persistence.DB)
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

	if err := model.TbMovieTagMgr(persistence.DB).Delete(&model.TbMovieTag{}, fmt.Sprintf("%s = ?", model.TbMovieTagColumns.TagID), val.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if err := mgr.Delete(&val).Error; err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, httpenc.MakeSuccHTTPRespPackage(nil))
}
