package handle

import (
	"net/http"
	"porn/common/httpenc"
	"porn/common/persistence"
	"porn/model"

	"github.com/gin-gonic/gin"
)

type TagListParam struct {
	Pagination httpenc.Pagination `json:"pagination"`
}
type TagInfo struct {
	ID    int    `json:"id"`
	Named string `json:"named"`
	Cover string `json:"cover"`
	Count int    `json:"count"`
}
type TagListData struct {
	Pagination httpenc.Pagination `json:"pagination"`
	TagInfos   []TagInfo          `json:"infos"`
	Total      int64              `json:"total"`
}

func TagList(c *gin.Context) {
	var (
		param TagListParam
		data  TagListData
	)
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusOK, httpenc.MakeFailHTTPRespPackage(httpenc.HTTPRespCode_ParamUnmarshal))
		return
	}

	sql :=
		`
		SELECT tb_tag.id, tb_tag.named, tb_tag.cover, COUNT(*) AS 'count' 
		FROM tb_tag, tb_movie_tag 
		WHERE tb_tag.id = tb_movie_tag.tag_id 
		GROUP BY tb_tag.id 
		ORDER BY 'count' DESC 
		LIMIT ? 
		OFFSET ?;
		`
	limit, offset := param.Pagination.LimitOrOffset()
	if err := persistence.DB.Raw(sql, limit, offset).Scan(&data.TagInfos).Error; err != nil {
		c.JSON(200, httpenc.MakeFailHTTPRespPackage(httpenc.HTTPRespCode_ParamIllegal))
		return
	}

	if err := model.TbTagMgr(persistence.DB).Count(&data.Total).Error; err != nil {
		c.JSON(500, nil)
		return
	}
	param.Pagination.PrevOrNext(data.Total)
	data.Pagination = param.Pagination
	c.JSON(http.StatusOK, httpenc.MakeSuccHTTPRespPackage(data))
}
func TagCreate(c *gin.Context) {}
func TagUpdate(c *gin.Context) {}
func TagDelete(c *gin.Context) {}
